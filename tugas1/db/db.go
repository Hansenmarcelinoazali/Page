package db

import (
	"tugas/config"
	// "ECHO-GORM/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func DbManager() *gorm.DB {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err := gorm.Open("mysql", connect_string)
	// defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic("DB Connection Error")
	}

	fmt.Println("DB Connected ( ͡❛ ͜ʖ͡❛ )")
	return db
}
