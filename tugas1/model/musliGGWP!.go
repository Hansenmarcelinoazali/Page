package model

// "github.com/jinzhu/gorm"

type Books struct {
	ID      int    `json:"ID"`
	Books   string `json:"Books"`
	Authors string `json:"Authors"`
}
