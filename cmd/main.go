package main

import (
	"log"
	"net/http"

	"github.com/devopsmi/rad/cmd/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root@/employees?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.LogMode(true)

	api.DB = db

	r := api.ConfigRouter()
	log.Fatal(http.ListenAndServe(":8899", r))
}
