package db

//DB：是一个数据库操作句柄，代表一个具有零到多个底层连接的连接池，它可以安全的被多个go程序同时使用。

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //数据库(MySQL)驱动 有该驱动golang才能操作mysql
)

var (
	Db  *sql.DB
	err error
)

func init() {
	//Open函数可能只是验证其参数，而不创建与数据库的连接。如果要检查数据源的名称是否合法，应调用返回值的Ping()方法。
	//Ping()检查与数据库的连接是否仍有效，如果需要会创建连接。
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
}
