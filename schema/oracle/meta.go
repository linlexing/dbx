package oracle

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

const driverName = "oci8"

type meta struct {
}

//CreateTableAs 执行create table as select语句
func (m *meta) CreateTableAsSQL(db common.DB, tableName, strSQL string, pks []string) ([]string, error) {
	return []string{
		fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL),
		fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ",")),
	}, nil
}

func (m *meta) TableNames(db common.DB) (names []string, err error) {
	strSQL := "SELECT table_name FROM user_tables"
	names = []string{}
	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			log.Println(err)
			return nil, err
		}
		names = append(names, name)
	}

	sort.Strings(names)
	return
}
func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {
	return tableExists(db, tabName)
}

func init() {
	schema.Register(driverName, new(meta))
}

func (m *meta) CreateTableSQL(db common.DB, tab *schema.Table) (rev []string, err error) {
	cols := []string{}
	for _, v := range tab.Columns {
		cols = append(cols, dbDefine(v))
	}
	if len(tab.PrimaryKeys) > 0 {
		rev = append(rev, fmt.Sprintf(
			"CREATE TABLE %s(\n%s,\nCONSTRAINT %s_pkey PRIMARY KEY(%s)\n)",
			tab.FullName(), strings.Join(cols, ",\n"), tab.Name, strings.Join(tab.PrimaryKeys, ",")))
	} else {
		rev = append(rev, fmt.Sprintf(
			"CREATE TABLE %s(\n%s\n)",
			tab.FullName(), strings.Join(cols, ",\n")))
	}
	//最后处理索引
	for _, col := range tab.Columns {
		if col.Index {
			rev = append(rev, createColumnIndexSQL(tab.FullName(), col.Name)...)
		}
	}
	return
}

//DropIndexIfExistsSQL 删除一个存在的索引，不存在返回nil
func (m *meta) DropIndexIfExistsSQL(db common.DB, indexName, tableName string) ([]string, error) {
	return []string{fmt.Sprintf(`
		DECLARE
		  COUNT_INDEXES INTEGER;
		BEGIN
		  SELECT COUNT(*) INTO COUNT_INDEXES
		    FROM USER_INDEXES
		    WHERE INDEX_NAME = '%s';

		  IF COUNT_INDEXES = 1 THEN
		    EXECUTE IMMEDIATE '%s';
		  END IF;
		END;`, strings.ToUpper(indexName),
		fmt.Sprintf("drop index %s", indexName))}, nil

}

func (m *meta) CreateIndexIfNotExistsSQL(db common.DB, indexName, tableName, express string) ([]string, error) {
	return []string{fmt.Sprintf(`
		DECLARE
		  COUNT_INDEXES INTEGER;
		BEGIN
		  SELECT COUNT(*) INTO COUNT_INDEXES
		    FROM USER_INDEXES
		    WHERE INDEX_NAME = '%s';

		  IF COUNT_INDEXES = 0 THEN
		    EXECUTE IMMEDIATE '%s';
		  END IF;
		END;`, strings.ToUpper(indexName),
		fmt.Sprintf("create index %s on %s(%s)", indexName, tableName, express))}, nil
}
