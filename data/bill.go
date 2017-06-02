package data

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/linlexing/dbx/common"

	"github.com/linlexing/mapfun"

	"log"

	"strings"
)

//Bill 一个简单的主表-明细表的集合，提供了读写的方法
type Bill struct {
	Main  *Table
	Child map[string]*Table
}

//BillRecord 代表一个单据记录，一个主表和N个明细表组成
type BillRecord struct {
	Main  map[string]interface{}
	Child map[string][]map[string]interface{}
}

//NewBill 返回一个新单据，传入驱动名称、数据库操作类，这里不用sqlx的原因是减少非必要的依赖
func NewBill(driver string, db common.DB, main string, child ...string) *Bill {
	if len(main) == 0 {
		log.Panic("not main table")
	}
	r := &Bill{
		Main:  NewTable(driver, db, main),
		Child: map[string]*Table{},
	}
	for _, v := range child {
		ct := NewTable(driver, db, v)
		//用表全名称
		r.Child[ct.FullName()] = ct
	}
	return r
}

//SetDB 重新设置DB，方便事务切换
func (b *Bill) SetDB(db common.DB) {
	b.Main.DB = db
	for _, v := range b.Child {
		v.DB = db
	}
}
func (b *Bill) db() common.DB {
	return b.Main.DB
}
func (b *Bill) keyValues(record *BillRecord) []interface{} {
	return b.Main.KeyValues(record.Main)
}
func (b *Bill) exists(keyValues ...interface{}) (bool, error) {
	return b.Main.KeyExists(keyValues...)
}

//ChangeKeyValues 修改一个记录的主键值
func (b *Bill) ChangeKeyValues(record *BillRecord, pks ...interface{}) {
	if record.IsEmpty() {
		return
	}
	//修改主表记录
	for i, name := range b.Main.PrimaryKeys {
		record.Main[name] = b.Main.ColumnByName(name).Type.ParseScan(pks[i])
	}
	//修改明细表记录，明细表主键的前几位必须是主表主键
	for _, child := range b.Child {
		childPk := child.PrimaryKeys
		for i, name := range b.Main.PrimaryKeys {
			for _, row := range record.Child[child.FullName()] {
				row[childPk[i]] = record.Main[name]
			}
		}
	}
	return
}

//Remove 移除一个记录，如果没有找到记录，返回一个error
func (b *Bill) Remove(oldRecord *BillRecord) (err error) {
	var iCount int64
	if iCount, err = b.Main.Remove(oldRecord.Main); err != nil {
		return err
	}
	if iCount == 0 {
		err = errors.New("can't foud then record")
		log.Println(err)
		return
	}
	if len(oldRecord.Child) == 0 {
		return nil
	}
	mainKeyValues := b.Main.KeyValues(oldRecord.Main)
	for k, v := range oldRecord.Child {
		if iCount, err = b.Child[k].Delete(v); err != nil {
			return err
		}
		//如果找不到要删除的字段，说明有其他用户操作过，返回一个错误
		if iCount == 0 {
			err = errors.New("can't foud then record")
			log.Println(err)
			return
		}
		//删除完成后，还需要检查是否有剩余，有的话说明明细记录已经被人改动过

		where := []string{}
		findKeys := []interface{}{}

		for si, sv := range mainKeyValues {
			where = append(where, fmt.Sprintf("%s=?", b.Child[k].PrimaryKeys[si]))
			findKeys = append(findKeys, sv)
		}
		has, err := b.Child[k].Exists(strings.Join(where, " and\n"), findKeys...)
		if err != nil {
			return err
		}
		if has {
			return fmt.Errorf("删除完成后，%s 中还含有数据", k)
		}
	}
	return nil
}

//Insert 插入一个单据记录
func (b *Bill) Insert(record *BillRecord) error {
	if err := b.Main.Insert([]map[string]interface{}{record.Main}); err != nil {
		return err
	}
	if len(record.Child) == 0 {
		return nil
	}
	for k, v := range record.Child {
		if err := b.Child[k].Insert(v); err != nil {
			return err
		}
	}
	return nil
}

//Save 保存一个记录，如果对应的记录存在则被覆盖
func (b *Bill) Save(record *BillRecord) error {
	//主表save
	if err := b.Main.Save(record.Main); err != nil {
		return err
	}
	if len(b.Child) == 0 {
		return nil
	}
	//明细表取出主键进行判断
	//删除库中多余的记录
	//然后逐个save
	mainKeyValues := b.Main.KeyValues(record.Main)
	for tabName, tab := range b.Child {
		cpk := tab.PrimaryKeys
		//组合where条件去查询
		where := []string{}
		whereVals := []interface{}{}
		for j, v := range mainKeyValues {
			where = append(where, fmt.Sprintf("%s=?", cpk[j]))
			whereVals = append(whereVals, v)
		}
		rows, err := tab.QueryRows(strings.Join(where, " and\n"), whereVals...)
		if err != nil {
			return err
		}
		for _, rv := range mapfun.Difference(rows, record.Child[tabName], cpk) {
			var iCount int64
			if iCount, err = tab.Remove(rv); err != nil {
				return err
			}
			if iCount == 0 {
				return errors.New("can't found remove record")
			}
		}
		for _, ins := range record.Child[tabName] {
			if err = tab.Save(ins); err != nil {
				return err
			}
		}

	}
	return nil
}

//Update 更新一个记录，旧记录的值必须要全相等
func (b *Bill) Update(oldRecord, newRecord *BillRecord) (err error) {
	var iCount int64
	if iCount, err = b.Main.Update(oldRecord.Main, newRecord.Main); err != nil {
		return
	}
	if iCount == 0 {
		err = errors.New("can't found the update record at table " + b.Main.FullName())
	}
	if len(b.Child) == 0 {
		return
	}
	for _, v := range b.Child {
		if _, _, _, err = v.Replace(oldRecord.Child[v.FullName()],
			newRecord.Child[v.FullName()]); err != nil {
			return
		}
	}
	return nil
}

//Record 返回一个记录，根据主键值
func (b *Bill) Record(keyValues ...interface{}) (result *BillRecord, err error) {
	mainRow, err := b.Main.Row(keyValues...)
	if err != nil {
		return
	}
	result = &BillRecord{
		Main: mainRow,
	}
	//load child
	if len(b.Child) == 0 {
		return
	}
	result.Child = map[string][]map[string]interface{}{}

	for _, tab := range b.Child {
		where := []string{}
		for i := range keyValues {
			where = append(where, fmt.Sprintf("%s=?", tab.PrimaryKeys[i]))
		}
		rows, err := tab.QueryRows(strings.Join(where, " and\n"),
			keyValues...)
		if err != nil {

		}
		result.Child[tab.FullName()] = rows

	}
	return
}

func init() {
	gob.Register(new(BillRecord))
}

//IsEmpty 返回是否为空记录
func (b *BillRecord) IsEmpty() bool {
	return b.Main == nil
}

//Encode 将Record压缩为byte数组
func (b *BillRecord) Encode() ([]byte, error) {
	out := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(out).Encode(b); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

//DecodeBillRecord 从一个字节数组解压
func DecodeBillRecord(in []byte) (*BillRecord, error) {
	bys := bytes.NewBuffer(in)
	out := new(BillRecord)
	if err := gob.NewDecoder(bys).Decode(out); err != nil {
		return nil, err
	}
	return out, nil
}
