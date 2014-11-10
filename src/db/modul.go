// 暂不处理的关键字包括有
// join
// seleclt 时 column 不指定 distinct
// have
package db

import (
	"fmt"
	"strings"
)

const (
	E_WHERE_CONDITION_AND = iota
	E_WHERE_CONDITION_OR
	E_WHERE_CONDITION_EQUAL
	E_WHERE_CONDITION_NEQ
)

const (
	E_COLUMN_GROUP_SINGLE = iota
	E_COLUMN_GROUP_COUNT
	E_COLUMN_GROUP_SUM
	E_COLUMN_GROUP_AVG
	E_COLUMN_GROUP_MAX
	E_COLUMN_GROUP_MIN
)

const (
	E_ORDER_SORT_ASC = iota // 升
	E_ORDER_SORT_DESC
)

type Modul struct {
	columnlist    string
	tablenamelist string
	// join          string
	where     string
	grouplist string
	// having    string
	orderlist string
	limit     string
}

/*
增加一列

返回值
success

repeatname

wronggrouptype
*/
func (m *Modul) AddColumn(szNew string, nGroupType int) string {
	err := "repeatname"
	if strings.Contains(m.columnlist, szNew) == true {
		return err
	}

	var szNewColumn string
	err = "wronggrouptype"
	switch nGroupType {
	case E_COLUMN_GROUP_SINGLE:
		szNewColumn = szNew
	case E_COLUMN_GROUP_COUNT:
		szNewColumn = fmt.Sprintf("COUNT(%v)", szNew)
	case E_COLUMN_GROUP_SUM:
		szNewColumn = fmt.Sprintf("SUM(%v)", szNew)
	case E_COLUMN_GROUP_AVG:
		szNewColumn = fmt.Sprintf("AVG(%v)", szNew)
	case E_COLUMN_GROUP_MAX:
		szNewColumn = fmt.Sprintf("MAX(%v)", szNew)
	case E_COLUMN_GROUP_MIN:
		szNewColumn = fmt.Sprintf("MIN(%v)", szNew)
	default:
		return err
	}

	if m.columnlist == "" {
		m.columnlist = szNew
	} else {
		m.columnlist = fmt.Sprintf("%v, %v", m.columnlist, szNewColumn)
	}
	return "success"
}

/*
增加一个表名

返回值
success

repeatname
*/
func (m *Modul) AddTableName(szNew string) string {
	err := "repeatname"
	if strings.Contains(m.tablenamelist, szNew) == true {
		return err
	}

	if m.tablenamelist == "" {
		m.tablenamelist = szNew
	} else {
		m.tablenamelist = fmt.Sprintf("%v, %v", m.tablenamelist, szNew)
	}
	return "success"
}

/*
增加一个 where 条件

返回值
success

wrongvaluetype

wrongvalueconditiontype

wrongconditiontype

todo:可以对其进行一些处理，避免用法上失误
*/
func (m *Modul) AddWhere(nConditionType int, szColumnName string, ColumnValue interface{}, nValueCondition int) string {
	szWhere := ""
	err := "wrongvaluetype"
	switch ColumnValue.(type) {
	case int:
		szWhere = fmt.Sprintf("%d", ColumnValue.(int))
	case string:
		szWhere = fmt.Sprintf("\"%v\"", ColumnValue.(string))
	default:
		return err
	}

	err = "wrongvalueconditiontype"
	switch nValueCondition {
	case E_WHERE_CONDITION_EQUAL:
		szWhere = fmt.Sprintf("(%v = %v)", szColumnName, szWhere)
	case E_WHERE_CONDITION_NEQ:
		szWhere = fmt.Sprintf("(%v != %v)", szColumnName, szWhere)
	default:
		return err
	}

	err = "wrongconditiontype"
	if m.where == "" {
		if nConditionType == E_WHERE_CONDITION_OR {
			return err
		}

		m.where = szWhere
	} else {
		if nConditionType == E_WHERE_CONDITION_OR {
			m.where = fmt.Sprintf("%v OR %v", m.where, szWhere)
		} else if nConditionType == E_WHERE_CONDITION_AND {
			m.where = fmt.Sprintf("%v AND %v", m.where, szWhere)
		}
	}

	return "success"
}

/*
增加 group

返回值

success

repeatname
*/
func (m *Modul) AddGroup(szNew string) string {
	err := "repeatname"
	if strings.Contains(m.grouplist, szNew) == true {
		return err
	}

	if m.grouplist == "" {
		m.grouplist = szNew
	} else {
		m.grouplist = fmt.Sprintf("%v, %v", m.grouplist, szNew)
	}
	return "success"
}

/*
增加 order

返回值

success

wrongsorttype

repeatname
*/
func (m *Modul) AddOrder(szNew string, nSortType int) string {
	err := "repeatname"
	if strings.Contains(m.orderlist, szNew) == true {
		return err
	}

	err = "wrongsorttype"
	if nSortType != E_ORDER_SORT_DESC && nSortType != E_ORDER_SORT_ASC {
		return err
	}

	bFirst := false
	if m.orderlist == "" {
		var szOrder string
		switch nSortType {
		case E_ORDER_SORT_DESC:
			szOrder = "DESC"
		case E_ORDER_SORT_ASC:
			szOrder = "ASC"
		}
		m.orderlist = fmt.Sprintf("%v", szOrder)
		bFirst = true
	} else {
		switch nSortType {
		case E_ORDER_SORT_ASC:
			if strings.Contains(m.orderlist, "ASC") == false {
				return err
			}
		case E_ORDER_SORT_DESC:
			if strings.Contains(m.orderlist, "DESC") == false {
				return err
			}
		}
	}

	if bFirst == true {
		m.orderlist = fmt.Sprintf("%v %v", szNew, m.orderlist)
	} else {
		m.orderlist = fmt.Sprintf("%v, %v", szNew, m.orderlist)
	}

	return "success"
}

/*
增加 limit

返回值

success

repeatadd

wrongrownum

wrongoffset
*/
func (m *Modul) AddLimit(nOffset int, nRowNum int) string {
	err := "wrongoffset"
	if nOffset < 0 {
		return err
	}

	err = "wrongrownum"
	if nRowNum < -1 {
		return err
	}

	err = "repeatadd"
	if m.limit != "" {
		return err
	}

	m.limit = fmt.Sprintf("%d, %d", nOffset, nRowNum)

	return "success"
}

/*
输出 sql 语句

返回值
success

emptytablename
*/
func (m *Modul) GenerateSelectSql() (err string, szSql string) {
	szSql = ""

	if m.columnlist == "" {
		m.columnlist = "*"
	}

	err = "emptytablename"
	if m.tablenamelist == "" {
		return err, ""
	}

	a := fmt.Sprintf("SELECT %v FROM %v", m.columnlist, m.tablenamelist)

	// if m.join != "" {
	// 	a = fmt.Sprintf("%v %v", a, m.join)
	// }

	if m.where != "" {
		a = fmt.Sprintf("%v WHERE %v", a, m.where)
	}

	if m.grouplist != "" {
		a = fmt.Sprintf("%v GROUP BY %v", a, m.grouplist)
	}

	if m.orderlist != "" {
		a = fmt.Sprintf("%v ORDER BY %v", a, m.orderlist)
	}

	if m.limit != "" {
		a = fmt.Sprintf("%v LIMIT %v", a, m.limit)
	}

	return "success", a
}
