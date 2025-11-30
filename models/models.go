package models

import (
	"time"

	"gorm.io/datatypes"
)

type Projects struct {
	ID        		uint      			`gorm:"primaryKey;autoIncrement" json:"id"`
	Title     		string    			`json:"title"`
	Category  		string   			`json:"category"`
	ImageClass 		string    			`json:"imageClass"`
	Technologies 	datatypes.JSON 		`json:"technologies"`
	Description 	string   			`json:"description"`
	Link       		datatypes.JSON    	`json:"link"`
	Featured    	bool    			`json:"featured"`
	Thumbnail      	string    			`json:"thumbnail"`
	CreatedAt 		time.Time 			`json:"createdAt"`
	UpdatedAt 		time.Time 			`json:"updatedAt"`
	ImagesURL      	datatypes.JSON    	`json:"imagesUrl"`
}

type User struct {
	ID			uint				`gorm:"primaryKey;autoIncrement" json:"id"`
	Username	string				`json:"username"`
	Password	string				`json:"password"`
	CreatedAt	time.Time			`json:"createdAt"`
	UpdatedAt	time.Time			`json:"updatedAt"`
}