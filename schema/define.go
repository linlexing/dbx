package schema

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type colDef struct {
	*Column
	isPK bool
}

func stringifyColumn(c *Column, isPk bool) (rev string, err error) {

	line := c.Name + "\t"
	if c.Type == TypeString {
		if c.MaxLength > 0 {
			line += fmt.Sprintf("STR(%d)", c.MaxLength)
		} else {
			line += "STR"
		}
	} else {
		ty, err := c.Type.String()
		if err != nil {
			return "", err
		}
		line += ty
	}
	//主键不需要not null
	if !isPk && !c.Null {
		line += " NOT NULL"
	}
	if c.Index == Index {
		line += " INDEX"
	} else if c.Index == UniqueIndex {
		line += " UINDEX"
	}
	if isPk {
		line += " PRIMARY KEY"
	}
	return line, nil
}

func columnDefine(propertyName string, line, trueTypeLine, formerNameTag, extendsTag string) (result *colDef, err error) {
	//先去除注释
	result = &colDef{
		Column: &Column{
			PropertyName: propertyName,
		},
	}
	if len(formerNameTag) > 0 {
		result.FormerName = strings.Fields(formerNameTag)
	}

	if idx := strings.Index(line, comment); idx >= 0 {
		line = line[:idx]
	}
	line = strings.ToUpper(strings.TrimSpace(line))

	//如果是后缀PRIMARY KEY，则说明当前字段是主键之一
	if strings.HasSuffix(line, "PRIMARY KEY") {
		line = line[:len(line)-11]
		result.isPK = true
	}

	lineList := columnReg.FindStringSubmatch(line)
	if len(lineList) == 0 {
		err = errors.New(line)
		return
	}
	//第一个是整行，需要去除
	lineList = lineList[1:]
	if len(lineList) == 0 {
		err = errors.New(line)
		return
	}
	result.Name = lineList[0]
	if len(strings.TrimSpace(lineList[1])) == 0 {
		//如果只有列名，则自动从上一个字段取出数据类型等定义
		err = fmt.Errorf("%s type not define", result.Name)
		return
	}
	dataType := strings.TrimSpace(lineList[1])
	notNull := false
	var index string
	var maxLength int64
	if len(lineList) > 2 {
		switch str := strings.TrimSpace(lineList[2]); str {
		case "NOT NULL":
			notNull = true
		case "NULL":
			notNull = false
		case "":
		default:
			err = fmt.Errorf("%s ,error define %s", line, str)
			return
		}
	}
	if len(lineList) > 3 {
		switch str := strings.TrimSpace(lineList[3]); str {
		case "INDEX":
			index = Index
		case "UINDEX":
			index = UniqueIndex
		case "":
			index = NoIndex
		default:
			err = fmt.Errorf("%s ,error define %s", line, str)
			return
		}
	}
	if strings.HasPrefix(dataType, "STR(") {
		maxLength, err = strconv.ParseInt(dataType[4:len(dataType)-1], 10, 64)
		if err != nil {
			return
		}
		dataType = "STR"
	}
	if result.Type, err = ParseDataType(dataType); err != nil {
		return
	}

	result.MaxLength = int(maxLength)
	//主键一定是not null
	result.Null = !(notNull || result.isPK)
	result.Index = index
	//下面处理truetype
	lineList = columnTrueType.FindStringSubmatch(trueTypeLine)
	if len(lineList) > 0 {
		//第一个是整行，需要去除
		lineList = lineList[1:]
		if len(lineList) == 0 {
			return
		}
		if len(lineList) != 2 {
			err = fmt.Errorf("TrueType (%s) not is [db type]", trueTypeLine)
			return
		}
		dbName, typeName := strings.TrimSpace(lineList[0]), strings.TrimSpace(lineList[1])
		if len(dbName) == 0 {
			err = fmt.Errorf("TrueType define dbname can't is empty")
			return
		}
		if len(typeName) == 0 {
			err = fmt.Errorf("TrueType define typename can't is empty")
			return
		}
		result.FetchDriver = dbName
		result.TrueType = typeName
	}
	if len(extendsTag) > 0 {
		result.Column.Extended = map[string]any{}
		//解开扩展属性
		for _, one := range strings.Fields(extendsTag) {
			list := strings.SplitN(one, "=", 2)
			if len(list) == 1 {
				result.Extended[list[0]] = ""
			} else {
				result.Extended[list[0]] = list[1]
			}
		}
	}
	return
}
func columnsDefine(d []*colDef) (columns []*Column, pks []string) {
	pks = []string{}
	columns = []*Column{}
	for _, cd := range d {
		col := cd.Column
		columns = append(columns, col)
		if cd.isPK {
			pks = append(pks, col.Name)
		}
	}
	return
}
