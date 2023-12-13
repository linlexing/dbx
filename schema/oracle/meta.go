package oracle

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

var (
	driverName = []string{"oci8", "oracle", "godror"}
)

type meta struct {
}

// CreateTableAs 执行create table as select语句
func (m *meta) CreateTableAsSQL(db common.DB, tableName, strSQL string, param []interface{},
	pks []string) ([]string, error) {
	return []string{
		fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL),
		fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ",")),
	}, nil
}
func (m *meta) TableEmpty(db common.DB, tableName string) (bool, error) {
	var a int
	if err := db.QueryRow(fmt.Sprintf(`SELECT 1 FROM DUAL WHERE EXISTS (SELECT 'X' FROM %s)`,
		tableName)).Scan(&a); err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}
	return false, nil
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
	err = rows.Err()
	sort.Strings(names)
	return
}
func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {
	return tableExists(db, tabName)
}

func init() {
	for _, one := range driverName {
		schema.Register(one, new(meta))
	}
}

func (m *meta) CreateTableSQL(db common.DB, tab *schema.Table) (rev []string, err error) {
	cols := []string{}
	for _, v := range tab.Columns {
		cols = append(cols, dbDefine(v))
	}
	if len(tab.PrimaryKeys) > 0 {
		rev = append(rev, fmt.Sprintf(
			"CREATE TABLE %s(\n\t%s,\n\tCONSTRAINT %s_pkey PRIMARY KEY(%s)\n)",
			tab.FullName(), strings.Join(cols, ",\n\t"), tab.Name, strings.Join(tab.PrimaryKeys, ",")))
	} else {
		rev = append(rev, fmt.Sprintf(
			"CREATE TABLE %s(\n%s\n)",
			tab.FullName(), strings.Join(cols, ",\n")))
	}
	//最后处理索引
	for _, col := range tab.Columns {
		if col.Index == schema.Index {
			rev = append(rev, createColumnIndexSQL(tab.FullName(), false, col.Name)...)
		} else if col.Index == schema.UniqueIndex {
			rev = append(rev, createColumnIndexSQL(tab.FullName(), true, col.Name)...)
		}
	}
	return
}

// DropIndexIfExistsSQL 删除一个存在的索引，不存在返回nil
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

func (m *meta) CreateIndexIfNotExistsSQL(db common.DB, unique bool, indexName, tableName, express string) ([]string, error) {
	idx := "index"
	if unique {
		idx = "unique index"
	}
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
		fmt.Sprintf("create %s %s on %s(%s)", idx, indexName, tableName, express))}, nil
}

/*
创建用户
创建表空间
修改用户默认表空间
为用户赋权
*/
func (m *meta) CreateSchemaSQL(db common.DB, dbInfo schema.DataBaseInfo) ([]string, error) {
	// createUser := fmt.Sprintf("CREATE USER %s IDENTIFIED BY %s", dbInfo.UserName, dbInfo.PassWord)
	// createTableSpace := fmt.Sprintf(`
	// 	CREATE TABLESPACE %s DATAFILE '%s' SIZE %dM
	// `, dbInfo.DBName, dbInfo.DataFile, dbInfo.DBSize)
	// alterUser := fmt.Sprintf("ALTER USER %s DEFAULT TABLESPACE %s", dbInfo.UserName, dbInfo.DBName)
	// grantUser := fmt.Sprintf("GRANT DBA TO %s", dbInfo.UserName)
	// return []string{createUser, createTableSpace, alterUser, grantUser}, nil
	return []string{}, nil
}

/*
删除用户 表空间
*/
func (m *meta) DropSchemaSQL(db common.DB, dbInfo schema.DataBaseInfo) ([]string, error) {
	// deleteUser := fmt.Sprintf("DROP USER %s CASCADE", dbInfo.UserName)
	// deleteTableSpace := fmt.Sprintf("DROP TABLESPACE %s INCLUDING CONTENTS AND DATAFILES", dbInfo.DBName)
	// return []string{deleteUser, deleteTableSpace}, nil
	return []string{}, nil
}
