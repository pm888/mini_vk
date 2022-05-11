package main

import (
	"database/sql"
	"fmt"
	"log"
	"mymod/cmd/social/service"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	NAME     = "root"
	PASSWORD = "qwerty123"
	HOSTNAME = "127.0.0.1:3306"
	DBNAME   = "BDusers"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", NAME, PASSWORD, HOSTNAME, DBNAME)
}

func OpenSql() *sql.DB {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Println(err)
	}
	return db

}

func main() {
	var port string
	f := os.Args[1:]
	for i := 0; i < len(f); i++ {
		port = f[i]
	}
	fmt.Println("Server run localhost", port)
	fmt.Println("MysQL run")
	DdMySql := OpenSql()
	srv := &service.Server{DdMySql}
	srv.RegisterHandlers()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
