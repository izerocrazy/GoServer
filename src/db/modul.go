// 暂不处理的关键字包括有
// join
// seleclt 时 column 不指定 distinct
// have
// todo : 此中代码有优化，0、它同时处理 select insert update，使用上容易出错；1、columnlist 存为数组，需要时再组；2、抽离一些通用的流程
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
	columnlist         string
	columnvaluelist    string // for insert
	columnnewvaluelist string // for update
	tablenamelist      string
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
增加一个 Column 名，和 Column 的值

返回值

success

repeatname

wrongvaluetype
*/
//first check value and column both right
func (m *Modul) AddColumnValue(szColumnName string, ColumnValue interface{}) string {
	// check for value type
	err := "wrongvaluetype"
	switch ColumnValue.(type) {
	case int:
	case string:
	default:
		return err
	}

	// check for column name: repeatname
	err = "repeatname"
	if strings.Contains(m.columnlist, szColumnName) == true {
		return err
	}

	// add
	if m.columnlist == "" {
		m.columnlist = szColumnName
	} else {
		m.columnlist = fmt.Sprintf("%v, %v", m.columnlist, szColumnName)
	}

	szValue := ""
	switch ColumnValue.(type) {
	case int:
		szValue = fmt.Sprintf("%d", ColumnValue.(int))
	case string:
		szValue = fmt.Sprintf("\"%v\"", ColumnValue.(string))
	}

	if m.columnvaluelist == "" {
		m.columnvaluelist = szValue
	} else {
		m.columnvaluelist = fmt.Sprintf("%v, %v", m.columnvaluelist, szValue)
	}

	return "success"
}

/*
增加一个 Column 名，和 Column 的值

返回值

success

repeatname

wrongvaluetype
*/
func (m *Modul) AddNewColumnValue(szColumnName string, ColumnValue interface{}) string {
	// check for value type
	err := "wrongvaluetype"
	switch ColumnValue.(type) {
	case int:
	case string:
	default:
		return err
	}

	// check for column name: repeatname
	err = "repeatname"
	if strings.Contains(m.columnlist, szColumnName) == true {
		return err
	}

	// add
	if m.columnlist == "" {
		m.columnlist = szColumnName
	} else {
		m.columnlist = fmt.Sprintf("%v, %v", m.columnlist, szColumnName)
	}

	szValue := ""
	switch ColumnValue.(type) {
	case int:
		szValue = fmt.Sprintf("%d", ColumnValue.(int))
	case string:
		szValue = fmt.Sprintf("\"%v\"", ColumnValue.(string))
	}

	if m.columnnewvaluelist == "" {
		m.columnnewvaluelist = fmt.Sprintf("%v = %v", szColumnName, szValue)
	} else {
		m.columnnewvaluelist = fmt.Sprintf("%v, %v = %v", m.columnnewvaluelist, szColumnName, szValue)
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
输出 select sql 语句

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

	szSql = fmt.Sprintf("SELECT %v FROM %v", m.columnlist, m.tablenamelist)

	err, szSql = m.addWhereSql(szSql)
	if err != "success" {
		return err, ""
	}

	return "success", szSql
}

/*
输出 insert sql 语句

返回值
success

emptytablename

multtablename

emptyinsertvalue

wronginsertvalue : [[!!! no achieve !!!]]the count of columnlist is not equit the count of columnvaluelist
*/
func (m *Modul) GenerateInsertSql() (err string, szSql string) {
	szSql = ""
	err = "emptytablename"
	if m.tablenamelist == "" {
		return err, szSql
	}

	err = "multtablename"
	if strings.Contains(m.tablenamelist, ",") == true {
		return err, szSql
	}

	err = "emptyinsertvalue"
	if m.columnlist == "" || m.columnvaluelist == "" {
		return err, szSql
	}

	szSql = fmt.Sprintf("INSERT %v (%v) VALUES (%v)", m.tablenamelist, m.columnlist, m.columnvaluelist)
	return "success", szSql
}

/*
输出 update sql 语句

返回值
success

emptytablename

multtablename

emptyupdatevalue
*/
func (m *Modul) GenerateUpdateSql() (err string, szSql string) {
	szSql = ""

	err = "emptytablename"
	if m.tablenamelist == "" {
		return err, ""
	}

	err = "multtablename"
	if strings.Contains(m.tablenamelist, ",") == true {
		return err, szSql
	}

	err = "emptyupdatevalue"
	if m.columnlist == "" || m.columnnewvaluelist == "" {
		return err, ""
	}

	szSql = fmt.Sprintf("UPDATE %v SET (%v)", m.tablenamelist, m.columnnewvaluelist)

	err, szSql = m.addWhereSql(szSql)
	if err != "success" {
		return err, ""
	}

	return "success", szSql
}

func (m *Modul) addWhereSql(szOldSql string) (err string, szSql string) {
	szSql = szOldSql

	if m.where != "" {
		szSql = fmt.Sprintf("%v WHERE %v", szSql, m.where)
	}

	if m.grouplist != "" {
		szSql = fmt.Sprintf("%v GROUP BY %v", szSql, m.grouplist)
	}

	if m.orderlist != "" {
		szSql = fmt.Sprintf("%v ORDER BY %v", szSql, m.orderlist)
	}

	if m.limit != "" {
		szSql = fmt.Sprintf("%v LIMIT %v", szSql, m.limit)
	}

	return "success", szSql
}
