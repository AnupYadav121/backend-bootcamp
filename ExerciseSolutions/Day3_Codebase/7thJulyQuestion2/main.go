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

	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			panic("Error occurred in closing DB connection")
		}
	}(Config.DB)

	Config.DB.AutoMigrate(&Models.Student{})
	Config.DB.AutoMigrate(&Models.SubjectMarks{})
	r := Routes.SetupRouter()
	//running
	err := r.Run()
	if err != nil {
		panic("Error occurred during network router start")
	}
}
