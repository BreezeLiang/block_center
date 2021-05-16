package main

import (
	"block_center/config"
	"block_center/db"
	"block_center/router"
	"fmt"
)

func main() {
	fmt.Println("deal error")

	config.Init()

	db.InitSqlx()

	router.StartGin()
}
