// main.go
package main

import (
	"Day4-5_Codebase/config"
	"Day4-5_Codebase/models"
	"Day4-5_Codebase/routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			panic("Error Occurred in closing DB")
		}
	}(config.DB)

	config.DB.AutoMigrate(&models.Order{})
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.Customer{})
	config.DB.AutoMigrate(&models.Retailer{})

	r := routes.SetupRouter()
	//running
	err := r.Run()
	if err != nil {
		panic("Error occurred in handing start router")
	}
}
