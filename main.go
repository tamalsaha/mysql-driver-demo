package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
}
