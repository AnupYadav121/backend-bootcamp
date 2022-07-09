// main.go
package main

import (
	"7thJulyQuestion2/Config"
	"7thJulyQuestion2/Models"
	"7thJulyQuestion2/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
