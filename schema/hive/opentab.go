package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
	"github.com/sirupsen/logrus"
)

func (m *meta) OpenTable(db common.DB, tableName string) (*schema.Table, error) {
	t := schema.NewTable(tableName)
	rows, err := db.Query("describe EXTENDED " + tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//返回的结果如下：
	//col_name data_type comment
	//字段名 数据类型  注释（为空）
	//空行
	//Detailed Table Information
	//Constraints Primary Key for default.test1:[a1,a2,a3], Constraint Name: pk_257871129_1626146071032_0
	fetchColumn := true
	for rows.Next() {
		var colName, dataType, comment sql.NullString
		if err = rows.Scan(&colName, &dataType, &comment); err != nil {
			return nil, err
		}
		if len(colName.String) == 0 {
			fetchColumn = false
			continue
		}
		//如果获取的是列信息
		if fetchColumn {
			ty, len, err := fromDBType(dataType.String)
			if err != nil {
				return nil, err
			}
			t.Columns = append(t.Columns, &schema.Column{
				Name:        colName.String,
				Type:        ty,
				MaxLength:   len,
				Null:        true,
				TrueType:    dataType.String,
				FetchDriver: driverName,
				Index:       schema.NoIndex,
				Extended:    map[string]any{},
			})
		} else {
			pstr := "Primary Key for "
			if colName.String == "Constraints" && strings.HasPrefix(dataType.String, pstr) {
				str := dataType.String[len(pstr):]
				i1 := strings.Index(str, "[")
				i2 := strings.Index(str, "]")
				cols := strings.Split(str[i1+1:i2], ",")
				// 主键不能为空
				for _, one := range cols {
					t.ColumnByName(one).Null = false
				}
				t.PrimaryKeys = cols
			}
		}

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return t, nil
}

func fromDBType(ty string) (schema.DataType, int, error) {
	uty := strings.ToUpper(ty)
	switch uty {
	case "TINYINT", "SMALLINT", "INT", "BIGINT", "DECIMAL":
		return schema.TypeInt, 0, nil
	case "FLOAT", "DOUBLE", "DOUBLE PRECISION":
		return schema.TypeFloat, 0, nil
	case "TIMESTAMP", "DATE":
		return schema.TypeDatetime, 0, nil
	case "BINARY":
		return schema.TypeBytea, 0, nil
	case "STRING":
		return schema.TypeString, -1, nil
	default:
		if strings.HasPrefix(uty, "VARCHAR(") {
			i, err := strconv.ParseInt(uty[8:len(uty)-1], 10, 32)
			return schema.TypeString, int(i), err
		}
		if strings.HasPrefix(uty, "CHAR(") {
			i, err := strconv.ParseInt(uty[5:len(uty)-1], 10, 32)
			return schema.TypeString, int(i), err
		}
		if strings.HasPrefix(uty, "DECIMAL(") {
			return schema.TypeFloat, 0, nil
		}
		logrus.WithFields(logrus.Fields{
			"type": ty,
		}).Panic("invalid type")
		return schema.TypeString, 0, fmt.Errorf("invalid type [%s]", ty)
	}
}
