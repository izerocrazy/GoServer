package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

/*
dbclient 的初始化

success

userpwderror

connecterror

dbsqlerror

db, err := sql.Open("mysql", "root:king+5688@tcp(192.168.1.11:3306)/test?charset=utf8")
*/
func TestInit() {
	var client DBClient
	ip := "192.168.1.11"
	username := "root"
	userpwd := "king+5688"

	client.Init(username, userpwd, ip, 0)
}

/*
执行一段 sql 语句

success

sqlerror
*/
func TestDoSql() {
	var client DBClient
	ip := "192.168.1.11"
	username := "root"
	userpwd := "king+5688"

	client.Init(username, userpwd, ip, 0)

	client.DoSql("select * from user")
}

/*
当一个表名不存在的时候创建它
当列名不为空的时候，表中不存在的时候，创建它

success

sqlerror
*/
func TestIfNotExitCreate() {

}

/*
执行一个 Modul，可以自动添加表和列名

success

sqlerror
*/
func TestDoModul() {

}
