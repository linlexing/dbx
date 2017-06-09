package schema

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type colDef struct {
	*Column
	copyPrev bool
	isPK     bool
}

func columnDefine(line string) (result *colDef, err error) {
	//先去除注释
	result = &colDef{
		Column: &Column{},
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
		result.copyPrev = true
		return
	}
	dataType := strings.TrimSpace(lineList[1])
	notNull := false
	index := false
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
			index = true
		case "":
		default:
			err = fmt.Errorf("%s ,error define %s", line, str)
		}
	}
	if strings.HasPrefix(dataType, "STR(") {
		maxLength, err = strconv.ParseInt(dataType[4:len(dataType)-1], 10, 64)
		if err != nil {
			return
		}
		dataType = "STR"
	}
	result.Type = ParseDataType(dataType)
	result.MaxLength = int(maxLength)
	result.Null = !notNull
	result.Index = index
	return
}
func columnsDefine(d []*colDef) (columns []*Column, pks []string) {
	pks = []string{}
	columns = []*Column{}
	var prevColumn *Column
	for _, cd := range d {
		col := cd.Column
		if cd.copyPrev {
			newcol := prevColumn.Clone()
			newcol.Name = cd.Name
			col = newcol
		}
		prevColumn = col
		columns = append(columns, col)
		if cd.isPK {
			pks = append(pks, col.Name)
		}
	}
	return
}
