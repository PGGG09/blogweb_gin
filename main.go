package main

import (
	"blogweb_gin/routers"
	"blogweb_gin/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.InitMysql()
	router := routers.InitRouter()
	// static resources
	router.Static("/static", "./static")
	router.Run(":8081")
}

//func main() {
//	dsn := "ocean_dev:ocean_1314@tcp(42.192.153.239:6633)/blogweb_gin"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	err = db.Ping()
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	log.Fatal("Test OK!")
//}
