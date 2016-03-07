package db

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

/*
dbclient 的初始化

success

userpwderror

connecterror

databaseerror

dbsqlerror

db, err := sql.Open("mysql", "root:king+5688@tcp(192.168.1.11:3306)/test?charset=utf8")
*/
func TestInit(t *testing.T) {
	var client DBClient
	ip := "127.0.0.1"
	username := ""
	userpwd := "ydp_!SP9bgth"
	nPort := 3306
	szDataBase := "test"

	err := client.Init(username, userpwd, ip, nPort, szDataBase)
	if err != "userpwderror" {
		t.Log("DBClient Init", err)
		t.FailNow()
	}

	username = "root"
	userpwd = ""
	err = client.Init(username, userpwd, ip, nPort, szDataBase)
	if err != "userpwderror" {
		t.Log("DBClient Init", err)
		t.FailNow()
	}

	userpwd = "ydp_!SP9bgth"
	ip = ""
	err = client.Init(username, userpwd, ip, nPort, szDataBase)
	if err != "connecterror" {
		t.Log("DBClient Init", err)
		t.FailNow()
	}

	ip = "127.0.0.1"
	nPort = -1
	err = client.Init(username, userpwd, ip, nPort, szDataBase)
	if err != "connecterror" {
		t.Log("DBClient Init", err)
		t.FailNow()
	}

	nPort = 0
	szDataBase = ""
	err = client.Init(username, userpwd, ip, nPort, szDataBase)
	if err != "databaseerror" {
		t.Log("DBClient Init", err)
		t.FailNow()
	}

	szDataBase = "noexist"
	err = client.Init(username, userpwd, ip, nPort, szDataBase)
	if err != "dbsqlerror" {
		t.Log("DBClient Init", err)
		t.FailNow()
	}

	szDataBase = "test"
	err = client.Init(username, userpwd, ip, nPort, szDataBase)
	if err != "success" {
		t.Log("DBClient Init", err)
		t.FailNow()
	}
}

/*
执行一段 sql 语句

success

noconnect

szsqlwrong

dbsqlerror
*/
func TestDoSql(t *testing.T) {
	var client DBClient
	// check for noconnect
	err, _ := client.DoSql("select * from user")
	if err != "noconnect" {
		t.Log("DBClient DoSql", err)
		t.FailNow()
	}

	// check for szsqlwrong
	ip := "192.168.1.11"
	username := "root"
	userpwd := "king+5688"
	szDataBase := "test"

	err = client.Init(username, userpwd, ip, 0, szDataBase)
	if err != "success" {
		t.Log("DBClient DoSql Init", err)
		t.FailNow()
	}

	err, _ = client.DoSql("")
	if err != "szsqlwrong" {
		t.Log("DBClient DoSql", err)
		t.FailNow()
	}

	err, _ = client.DoSql("select * from user")
	if err != "success" {
		t.Log("DBClient DoSql", err)
		t.FailNow()
	}
}

/*
当一个表名不存在的时候创建它
当列名不为空的时候，表中不存在的时候，创建它

success

sqlerror
*/

/*func TestIfNotExitCreate(t *testing.T) {

}*/

/*
执行一个 Modul，可以自动添加表和列名

success

sqlerror
*/

/*func TestDoModul() {

}*/
