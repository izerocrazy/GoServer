package db

import (
	"testing"
)

/*
增加一列

返回值
success

repeatname

wronggrouptype
*/
func TestAddColumn(t *testing.T) {
	var modul Modul
	err := modul.AddColumn("test", E_COLUMN_GROUP_SINGLE)
	if err != "success" {
		t.Log("modul add column", err)
		t.FailNow()
	}

	err = modul.AddColumn("test", E_COLUMN_GROUP_SINGLE)
	if err != "repeatname" {
		t.Log("modul add column", err)
		t.FailNow()
	}

	err = modul.AddColumn("test1", E_COLUMN_GROUP_SINGLE-1)
	if err != "wronggrouptype" {
		t.Log("modul add column", err)
		t.FailNow()
	}

	err = modul.AddColumn("test2", E_COLUMN_GROUP_MAX)
	if err != "success" {
		t.Log("modul add column", err)
		t.FailNow()
	}

	err = modul.AddColumn("test3", E_COLUMN_GROUP_MIN+1)
	if err != "wronggrouptype" {
		t.Log("modul add column", err)
		t.FailNow()
	}
}

/*
增加一个表名

返回值
success

repeatname
*/
func TestAddTableName(t *testing.T) {
	var modul Modul
	err := modul.AddTableName("test")
	if err != "success" {
		t.Log("modul add table name", err)
		t.FailNow()
	}

	err = modul.AddTableName("test1")
	if err != "success" {
		t.Log("modul add table name", err)
		t.FailNow()
	}

	err = modul.AddTableName("test1")
	if err != "repeatname" {
		t.Log("modul add table name", err)
		t.FailNow()
	}
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
func TestAddWhere(t *testing.T) {
	var modul Modul
	err := modul.AddWhere(E_WHERE_CONDITION_OR, "testname", 1, E_WHERE_CONDITION_NEQ)
	if err != "wrongconditiontype" {
		t.Log("modul add where ", err)
		t.FailNow()
	}

	err = modul.AddWhere(E_WHERE_CONDITION_AND, "testname", 0.1, E_WHERE_CONDITION_EQUAL)
	if err != "wrongvaluetype" {
		t.Log("modul add where", err)
		t.FailNow()
	}

	err = modul.AddWhere(E_WHERE_CONDITION_AND, "testname", 1, E_WHERE_CONDITION_OR)
	if err != "wrongvalueconditiontype" {
		t.Log("modul add where", err)
		t.FailNow()
	}

	err = modul.AddWhere(E_WHERE_CONDITION_AND, "testname", "test", E_WHERE_CONDITION_NEQ)
	if err != "success" {
		t.Log("modul add where", err)
		t.FailNow()
	}
}

/*
增加 group

返回值

success

repeatname
*/
func TestAddGroup(t *testing.T) {
	var modul Modul
	err := modul.AddGroup("test")
	if err != "success" {
		t.Log("modul add group", err)
		t.FailNow()
	}

	err = modul.AddGroup("test")
	if err != "repeatname" {
		t.Log("modul add group", err)
		t.FailNow()
	}
}

/*
增加 order

返回值

success

wrongsorttype

repeatname
*/
func TestAddOrder(t *testing.T) {
	var modul Modul
	err := modul.AddOrder("test", E_ORDER_SORT_ASC)
	if err != "success" {
		t.Log("modul add order", err)
		t.FailNow()
	}

	err = modul.AddOrder("test", E_ORDER_SORT_ASC)
	if err != "repeatname" {
		t.Log("modul add order", err)
		t.FailNow()
	}

	err = modul.AddOrder("test1", E_ORDER_SORT_DESC)
	if err != "wrongsorttype" {
		t.Log("modul add order", err)
		t.FailNow()
	}

	err = modul.AddOrder("test1", E_ORDER_SORT_ASC)
	if err != "success" {
		t.Log("modul add order", err)
		t.FailNow()
	}
}

/*
增加 limit

返回值

success

repeatadd

wrongrownum

wrongoffset
*/
func TestAddLimit(t *testing.T) {
	var modul Modul
	err := modul.AddLimit(-1, 1)
	if err != "wrongoffset" {
		t.Log("modul add limit", err)
		t.FailNow()
	}

	err = modul.AddLimit(0, -2)
	if err != "wrongrownum" {
		t.Log("modul add limit", err)
		t.FailNow()
	}

	err = modul.AddLimit(0, -1)
	if err != "success" {
		t.Log("modul add limit", err)
		t.FailNow()
	}

	err = modul.AddLimit(1, 5)
	if err != "repeatadd" {
		t.Log("modul add limit", err)
		t.FailNow()
	}
}

/*
输出 sql 语句

返回值
success

emptytablename
*/
func TestGenerateSelectSql(t *testing.T) {
	var modul Modul
	err, szSql := modul.GenerateSelectSql()
	if err != "emptytablename" {
		t.Log("modul GenerateSelectSql", err)
		t.FailNow()
	}

	// test for AddTableName
	err = modul.AddTableName("test")
	if err != "success" {
		t.Log("modul GenerateSelectSql AddTableName", err)
		t.FailNow()
	}
	err, szSql = modul.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test" {
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}

	err = modul.AddTableName("test1")
	if err != "success" {
		t.Log("modul GenerateSelectSql AddTableName", err)
		t.FailNow()
	}
	err, szSql = modul.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test, test1" {
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}

	// test for AddColumn
	var modul2 Modul
	modul2.AddTableName("test") // 这里就不做判断了
	err = modul2.AddColumn("column1", E_COLUMN_GROUP_SINGLE)
	if err != "success" {
		t.Log("modul GenerateSelectSql AddColumn", err)
		t.FailNow()
	}
	err, szSql = modul2.GenerateSelectSql()
	if err != "success" || szSql != "SELECT column1 FROM test" {
		t.Log("modul GenerateSelectSql", err)
		t.FailNow()
	}
	err = modul2.AddColumn("column2", E_COLUMN_GROUP_SINGLE)
	if err != "success" {
		t.Log("modul GenerateSelectSql AddColumn", err)
		t.FailNow()
	}
	err, szSql = modul2.GenerateSelectSql()
	if err != "success" || szSql != "SELECT column1, column2 FROM test" {
		t.Log("modul GenerateSelectSql", err)
		t.FailNow()
	}

	// test for AddWhere
	var modul3 Modul
	modul3.AddTableName("test")
	err = modul3.AddWhere(E_WHERE_CONDITION_AND, "column1", "value1", E_WHERE_CONDITION_EQUAL)
	if err != "success" {
		t.Log("modul GenerateSelectSql AddWhere", err)
		t.FailNow()
	}
	err, szSql = modul3.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test WHERE (column1 = \"value1\")" {
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}
	err = modul3.AddWhere(E_WHERE_CONDITION_OR, "column2", 1, E_WHERE_CONDITION_NEQ)
	if err != "success" {
		t.Log("modul GenerateSelectSql AddWhere", err)
		t.FailNow()
	}
	err, szSql = modul3.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test WHERE (column1 = \"value1\") OR (column2 != 1)" {
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}

	// test for AddGroup
	var modul4 Modul
	modul4.AddTableName("test")
	err = modul4.AddGroup("group")
	if err != "success" {
		t.Log("modul GenerateSelectSql AddGroup", err)
		t.FailNow()
	}
	err, szSql = modul4.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test GROUP BY group" {
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}
	err = modul4.AddGroup("group2")
	if err != "success" {
		t.Log("modul GenerateSelectSql AddGroup", err)
		t.FailNow()
	}
	err, szSql = modul4.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test GROUP BY group, group2" {
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}

	// test for AddOrder
	var modul5 Modul
	modul5.AddTableName("test")
	err = modul5.AddOrder("order", E_ORDER_SORT_ASC)
	if err != "success" {
		t.Log("modul GenerateSelectSql AddOrder", err)
		t.FailNow()
	}
	err, szSql = modul5.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test ORDER BY order ASC" {
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}
	err = modul5.AddOrder("order2", E_ORDER_SORT_ASC)
	if err != "success" {
		t.Log("modul GenerateSelectSql AddOrder", err)
		t.FailNow()
	}
	err, szSql = modul5.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test ORDER BY order2, order ASC" { // order 是反序
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}

	// test for AddLimit
	var modul6 Modul
	modul6.AddTableName("test")
	err = modul6.AddLimit(0, 5)
	if err != "success" {
		t.Log("modul GenerateSelectSql AddLimit", err)
		t.FailNow()
	}
	err, szSql = modul6.GenerateSelectSql()
	if err != "success" || szSql != "SELECT * FROM test LIMIT 0, 5" {
		t.Log("modul GenerateSelectSql", err, szSql)
		t.FailNow()
	}
}
