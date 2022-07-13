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

	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			panic("Error Occurred in closing DB")
		}
	}(Config.DB)

	Config.DB.AutoMigrate(&Models.Order{})
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{})
	Config.DB.AutoMigrate(&Models.Retailer{})

	r := Routes.SetupRouter()
	//running
	err := r.Run()
	if err != nil {
		panic("Error occurred in handing start router")
	}
}
