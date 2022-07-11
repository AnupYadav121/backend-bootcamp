// main.go
package main

import (
	"7thJulyQuestion3/Config"
	"7thJulyQuestion3/Models"
	"7thJulyQuestion3/Routes"
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
