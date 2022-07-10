// main.go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"july8Files/Config"
	"july8Files/Models"
	"july8Files/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Order{})
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Transaction{})
	Config.DB.AutoMigrate(&Models.Customer{})
	Config.DB.AutoMigrate(&Models.OrderUpdated{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}
