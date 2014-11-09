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
		t.Log("modul add columu", err)
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

}
