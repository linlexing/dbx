package oracle

import (
	"fmt"
	"log"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func (m *meta) ChangeTable(db common.DB, change *schema.TableSchemaChange) error {
	tabName := change.NewName
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	if change.OldName != change.NewName {
		if err := tableRename(db, change.OldName, change.NewName); err != nil {
			return err
		}
	}
	//如果主键变更，则需要先除去主键
	if change.PKChange {
		if err := dropTablePrimaryKey(db, tabName); err != nil {
			return err
		}
	}
	//逐个处理字段
	for _, cf := range change.ChangeFields {

		if err := processColumn(db, tabName, cf.OldField, cf.NewField); err != nil {
			return err
		}
	}
	//最后删除字段
	if len(change.RemoveFields) > 0 {
		if err := removeColumns(db, tabName, change.RemoveFields); err != nil {
			return err
		}
	}
	//如果主键变过，则新增主键
	if change.PKChange {
		if err := addTablePrimaryKey(db, tabName, change.PK); err != nil {
			return err
		}
	}
	return nil
}
func processColumn(db common.DB, tabName string, oldCol, newCol *schema.Column) error {
	var strSQL string
	//如果是新增字段
	if oldCol == nil {
		strSQL = fmt.Sprintf("alter table %s add %s", tabName, dbDefine(newCol))
		if _, err := db.Exec(strSQL); err != nil {
			err = common.NewSQLError(err, strSQL)
			log.Println(err)
			return err
		}
		log.Println(strSQL)
		//处理索引
		if newCol.Index {
			if err := createColumnIndex(db, tabName, newCol.Name); err != nil {
				return err
			}
			log.Printf("table %s add column index %s\n", tabName, newCol.Name)
		}
		return nil
	}
	//如果是更名，需要先处理
	if oldCol.Name != newCol.Name {
		strSQL = fmt.Sprintf("alter table %s rename column %s to %s", tabName, oldCol.Name, newCol.Name)
		if _, err := db.Exec(strSQL); err != nil {
			err = common.NewSQLError(err, strSQL)
			log.Println(err)
			return err
		}
		log.Println(strSQL)
	}
	if !oldCol.Eque(newCol) {
		if oldCol.Null != newCol.Null {
			strSQL = fmt.Sprintf("alter table %s MODIFY %s", tabName, dbDefineNull(newCol))

		} else {
			strSQL = fmt.Sprintf("alter table %s MODIFY %s %s", tabName, newCol.Name, colDBType(newCol))
		}
		if _, err := db.Exec(strSQL); err != nil {
			err = common.NewSQLError(err, strSQL)
			log.Println(err)
			return err
		}
		log.Println(strSQL)
	}
	//处理索引,字段更名的操作，oracle、postgres、mysql都是安全的，所以不需处理
	//ref:http://stackoverflow.com/questions/6732896/does-rename-column-take-care-of-indexes
	if oldCol.Index && !newCol.Index {
		//删除索引
		if err := dropColumnIndex(db, tabName, oldCol.IndexName); err != nil {
			return err
		}
		log.Printf("drop table %s column %s index %s\n", tabName, newCol.Name, oldCol.IndexName)
	} else if !oldCol.Index && newCol.Index {
		//新增索引
		if err := createColumnIndex(db, tabName, oldCol.Name); err != nil {
			return err
		}
		log.Printf("create table %s column %s index\n", tabName, newCol.Name)
	}
	return nil
}
