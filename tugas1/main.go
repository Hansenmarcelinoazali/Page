package main

import (
	"tugas/routes"
)

// "tugas/db"
// "github.com/jinzhu/gorm"

// type Books struct {
// 	// gorm.Model
// 	Books   string
// 	Authors string
// }

func main() {

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
