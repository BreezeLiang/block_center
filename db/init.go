package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"runtime/debug"
	"time"
)

var (
	Mysqlx *sqlx.DB
)

func InitSqlx() {
	db, err := sqlx.Open("mysql", "root"+":"+"smmdb2016"+"@tcp("+"182.254.211.181"+":"+"3306"+")/"+"flow")
	if err != nil {
		log.Panic(fmt.Sprintf("mysql session 启动失败: %v\n堆栈信息: %v", err, string(debug.Stack())))
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf(fmt.Sprintf("mysql session  启动失败: %v\n堆栈信息: %v", err, string(debug.Stack())))
		os.Exit(0)
		return
	}

	Mysqlx = db

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Second * 300)
	log.Printf("InitSqlx Success.")
}
