package main

import (
	"block_center/db"
	"block_center/handler"
	"errors"
	"fmt"
	pkgErr "github.com/pkg/errors"
)

func main() {
	fmt.Println("deal error")

	db.InitSqlx()

	var EOF = errors.New("sql: no rows in result set")

	blockFlow, err := handler.GetBlockFlow("xxx")
	if err != nil {
		if pkgErr.Is(pkgErr.Cause(err), EOF) {
			fmt.Printf("记录不存在")
		} else {
			fmt.Printf("查询错误：%v", err)
		}
		return
	}

	fmt.Printf("查询到的结果为：%+v", blockFlow)
}
