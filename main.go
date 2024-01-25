package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	cfg, err := mysql.ParseDSN("tcp(mysql-demo.mysql.svc:3306)/")
	if err != nil {
		panic(err)
	}
	cfg.User = "user"
	cfg.Passwd = "password"
	// cfg.DBName = "adb"
	fmt.Println(cfg.FormatDSN())
	//
	//_, err := sql.Open("mysql", "tcp(mysql-demo.mysql.svc:3306)/")
	//if err != nil {
	//	panic(err)
	//}
}
