package account

import (
	"database/sql"
	"fmt"
	"log"

	// to use mysql db
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	username = "root"
	passward = "123456"
	network  = "tcp"
	host     = "113.31.106.132"
	port     = 3306
	database = "shop"
)

func Mysql() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, passward, network, host, port, database)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
