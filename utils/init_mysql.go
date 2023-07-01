package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() (err error) {
	fmt.Println("InitMysql....")
	if db == nil {
		db, err = sql.Open("mysql", "ocean_dev:ocean_1314@tcp(42.192.153.239:6633)/blogweb_gin")
		if err != nil {
			log.Fatal(err)
			return
		}
		CreateTableWithUser()
	}
	return
}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		create_time INT(10)
	);`
	ModifyDB(sql)
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
