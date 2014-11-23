package db

import (
	"base"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

/*
   现在做成每链接一次是生成一个 connect，查询完了再关闭
*/
type DBClient struct {
	szConnectStr string
}

/*
dbclient 的初始化：并不完全，没有设置一个链接的时长

success

userpwderror

connecterror

databaseerror

dbsqlerror

db, err := sql.Open("mysql", "root:king+5688@tcp(192.168.1.11:3306)/test?charset=utf8")
*/
func (client *DBClient) Init(szUsername string, szPassword string, szIP string, nPort int, szDataBase string) string {
	err := "userpwderror"
	// check for username and password
	if szUsername == "" || szPassword == "" {
		return err
	}

	// check for IP
	err = "connecterror"
	if szIP == "" || len(szIP) < 4 || strings.Contains(szIP, ".") == false {
		return err
	}

	// check for Port
	if nPort < 0 {
		return err
	}
	if nPort == 0 {
		nPort = 3306
	}

	// check for databse
	err = "databaseerror"
	if szDataBase == "" {
		return err
	}

	// compose the string and check open
	client.szConnectStr = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8", szUsername, szPassword, szIP, nPort, szDataBase)
	db, err2 := sql.Open("mysql", client.szConnectStr)
	defer db.Close()
	err = "dbsqlerror"
	if Base.CheckErr(err2) == false {
		return err
	}

	return "success"
}

/*
执行一段 sql 语句

success

noconnect

szsqlwrong

dbsqlerror
*/
func (client *DBClient) DoSql(szSql string) (err string, rows *sql.Rows) {
	err = "noconnect"
	// check for dbclient
	if client.szConnectStr == "" {
		return err, rows
	}

	// check for szSql
	err = "szsqlwrong"
	if szSql == "" {
		return err, rows
	}

	err = "dbsqlerror"
	db, err2 := sql.Open("mysql", client.szConnectStr)
	defer db.Close()
	if Base.CheckErr(err2) == false {
		return err, rows
	}

	rows, err2 = db.Query(szSql)
	if Base.CheckSqlQueryErr(err2) == false {
		return err, rows
	}

	return "success", rows
}

/*
如果数据库中不存在一个表，那么创建它

success

noconnect

tablenameerror

dbsqlerror
*/
func (clinet *DBClient) ifNotExistCreateTable(szTableName string) string {
	err := "tablenameerror"
	if szTableName == "" {
		return err
	}

	err = "noconnect"
	// check for dbclient
	if clinet.szConnectStr == "" {
		return err
	}

	err = "dbsqlerror"
	db, err2 := sql.Open("mysql", client.szConnectStr)
	defer db.Close()
	if Base.CheckErr(err2) == false {
		return err
	}

	// check datebase table name
	szSql := fmt.Sprintf("SELECT `TABLE_NAME` FROM information_schema.`TABLE_NAME` WHERE TABLE_SCHEMA =\"%v\"", szTableName)
	rows, err3 := db.Query(szSql)
	if err3 == sql.ErrNoRows {
		// if don't have table, then create
		// db.Query
	}

	return "success"
}

func (clinet *DBClient) ifNotExistCreateColumn(szTableName string, szColumn string) string {

}
