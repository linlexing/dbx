package dbx

import (
	"bytes"
	"dbweb/lib/safe"
	"encoding/gob"
	"fmt"

	"github.com/linlexing/mapfun"

	"github.com/jmoiron/sqlx"
)

//这里的bill就是一个简单的主表-明细表的集合，提供了读写的方法
type Bill struct {
	Main  *DBTable
	Child map[string]*DBTable
}
type BillRows struct {
	bill *Bill
	rows *sqlx.Rows
}
type BillRecord struct {
	Main  map[string]interface{}
	Child map[string][]map[string]interface{}
}

func NewBill(db DB, main string, child ...string) *Bill {
	if len(main) == 0 {
		panic("not main table")
	}
	r := &Bill{
		Main:  NewTable(db, main),
		Child: map[string]*DBTable{},
	}
	for _, v := range child {
		ct := NewTable(db, v)
		//表名称会转换
		r.Child[ct.Name] = ct
	}
	return r
}
func (b *Bill) SetDB(db DB) {
	b.Main.Db = db
	for _, v := range b.Child {
		v.Db = db
	}
}
func (b *Bill) KeyValues(record *BillRecord) []interface{} {
	return mapfun.Values(mapfun.Pick(record.Main, b.Main.PrimaryKeys()...))
}
func (b *Bill) Exists(keyValues ...interface{}) (bool, error) {
	return b.Main.Exists(mapfun.Object(b.Main.PrimaryKeys(), keyValues))
}

//修改一个记录的主键值
func (b *Bill) ChangeKeyValues(record *BillRecord, pks ...interface{}) {
	if record.IsEmpty() {
		return
	}
	//修改主表记录
	for i, name := range b.Main.PrimaryKeys() {
		record.Main[name] = b.Main.Field(name).ConvertToTrueType(pks[i])
	}
	//修改明细表记录，明细表主键的前几位必须是主表主键
	for _, child := range b.Child {
		childPk := child.PrimaryKeys()
		for i, name := range b.Main.PrimaryKeys() {
			for _, row := range record.Child[child.Name] {
				row[childPk[i]] = record.Main[name]
			}
		}
	}
	return
}
func (b *Bill) Clone() *Bill {
	result := &Bill{
		Main: b.Main.Clone(),
	}
	if len(b.Child) == 0 {
		return result
	}
	result.Child = map[string]*DBTable{}
	for k, v := range b.Child {
		result.Child[k] = v.Clone()
	}
	return result
}
func (b *Bill) Create() (err error) {
	if err = b.Main.Create(); err != nil {
		return
	}
	if len(b.Child) == 0 {
		return
	}
	for _, v := range b.Child {
		if err = v.Create(); err != nil {
			return
		}
	}
	return
}
func (b *Bill) Count(where string, param map[string]interface{}) (icount int64, err error) {
	var strSql string

	if len(where) == 0 {
		strSql = fmt.Sprintf("select count(*) from %s", b.Main.Name)
	} else {
		strSql = fmt.Sprintf("select count(*) from %s where %s", b.Main.Name, where)
	}

	strSql = RenderSql(strSql, param)

	vCount, err := GetSqlFun(b.Main.Db, strSql, nil)
	if err != nil {
		return -1, SqlError{strSql, nil, err}
	}
	return safe.Int(vCount), nil
}
func (b *Bill) NameQuery(where string, renderParam map[string]interface{}) (*BillRows, error) {
	var strSql string

	if len(where) == 0 {
		strSql = fmt.Sprintf("select * from %s", b.Main.Name)
	} else {
		strSql = fmt.Sprintf("select * from %s where %s", b.Main.Name, where)
	}

	strSql = RenderSql(strSql, renderParam)

	strSql, sqlParam := BindSql(b.Main.Db, strSql, nil)
	rows, err := b.Main.Db.Queryx(strSql, sqlParam...)
	if err != nil {
		return nil, SqlError{strSql, sqlParam, err}
	}
	return &BillRows{b, rows}, nil
}
func (b *Bill) Remove(oldRecord *BillRecord) error {
	if err := b.Main.Remove(oldRecord.Main); err != nil {
		return err
	}
	if len(oldRecord.Child) == 0 {
		return nil
	}
	mainKeyValues := b.KeyValues(oldRecord)
	for k, v := range oldRecord.Child {
		if err := b.Child[k].Delete(v); err != nil {
			return err
		}
		//删除完成后，还需要检查是否有剩余，有的话说明明细记录已经被人改动过
		childPK := b.Child[k].PrimaryKeys()
		findKey := map[string]interface{}{}

		for si, sv := range mainKeyValues {
			findKey[childPK[si]] = sv
		}
		has, err := b.Child[k].Exists(findKey)
		if err != nil {
			return err
		}
		if has {
			return fmt.Errorf("删除完成后，%s 中还含有数据", k)
		}
	}
	return nil
}
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

//保存一个记录，如果对应的记录存在则被覆盖
func (b *Bill) Save(record *BillRecord) error {
	//主表save
	if err := b.Main.Save(record.Main); err != nil {
		return err
	}
	//明细表取出主键进行判断
	//删除库中多余的记录
	//然后逐个save
	mainKeyValues := b.Main.KeyValues(record.Main)
	for tabName, tab := range b.Child {
		cpk := tab.PrimaryKeys()
		//必须先获取字段，否则运行时，postgres不允许多个query在一个事务中并行取
		//https://github.com/lib/pq/issues/81
		tab.FetchColumns()
		query := map[string]interface{}{}
		for j, v := range mainKeyValues {
			query[cpk[j]] = v
		}
		rows, err := tab.Rows(query, cpk...)
		if err != nil {
			return err
		}
		for _, rv := range mapfun.Difference(rows, record.Child[tabName], cpk) {
			if err = tab.Remove(rv); err != nil {
				return err
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

//更新一个记录，旧记录的值必须要相等
func (b *Bill) Update(oldRecord, newRecord *BillRecord) error {
	if err := b.Main.Update(oldRecord.Main, newRecord.Main); err != nil {
		return err
	}
	if len(b.Child) == 0 {
		return nil
	}
	for _, v := range b.Child {
		if err := v.Replace(oldRecord.Child[v.Name], newRecord.Child[v.Name]); err != nil {
			return err
		}
	}
	return nil
}
func (b *Bill) Record(keyValues ...interface{}) (result *BillRecord, err error) {
	mainRow := b.Main.Row(keyValues...)
	if mainRow == nil {
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
	mainKeyValues := mapfun.Values(mapfun.Pick(mainRow, b.Main.PrimaryKeys()...))

	for _, tab := range b.Child {
		cpk := tab.PrimaryKeys()
		query := map[string]interface{}{}
		for j, v := range mainKeyValues {
			query[cpk[j]] = v
		}
		if rows, err := tab.Rows(query); err != nil {
			return nil, err
		} else {
			result.Child[tab.Name] = rows
		}
	}
	return
}

func (b *BillRows) Next() bool {
	return b.rows.Next()

}
func (b *BillRows) Err() error {
	return b.rows.Err()
}
func (b *BillRows) Close() error {
	return b.rows.Close()
}
func (b *BillRows) Record() (result *BillRecord, err error) {
	mainRow := map[string]interface{}{}
	err = b.rows.MapScan(mainRow)
	if err != nil {
		return nil, err
	}
	//字段名转换成大写，数据类型正确转换
	mainRow = b.bill.Main.ConvertToTrueType(mainRow)
	result = &BillRecord{
		Main: mainRow,
	}
	//load child
	if len(b.bill.Child) == 0 {
		return
	}
	result.Child = map[string][]map[string]interface{}{}
	mainKeyValues := mapfun.Values(mapfun.Pick(mainRow, b.bill.Main.PrimaryKeys()...))
	for _, tab := range b.bill.Child {
		cpk := tab.PrimaryKeys()
		query := map[string]interface{}{}
		for j, v := range mainKeyValues {
			query[cpk[j]] = v
		}
		if rows, err := tab.Rows(query); err != nil {
			return nil, err
		} else {
			result.Child[tab.Name] = rows
		}
	}
	return
}
func init() {
	gob.Register(new(BillRecord))
}

//返回是否为空记录
func (b *BillRecord) IsEmpty() bool {
	return b.Main == nil
}

//将Record压缩为byte数组
func (b *BillRecord) Encode() ([]byte, error) {
	out := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(out).Encode(b); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

//从一个字节数组解压
func DecodeBillRecord(in []byte) (*BillRecord, error) {
	bys := bytes.NewBuffer(in)
	out := new(BillRecord)
	if err := gob.NewDecoder(bys).Decode(out); err != nil {
		return nil, err
	}
	return out, nil
}
