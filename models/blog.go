package models

import (
	"github.com/dhairyajoshi/go-rest-api/helpers/database"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title   string
	Content string
	UserId  int64
}

func migrateBlogs() {

	conn := database.GetDbConnection()

	conn.AutoMigrate(&Blog{})
}

func (blog Blog) Save() {
	conn := database.GetDbConnection()

	conn.Create(&blog)

}

func (blog Blog) Get(id ...int64) []Blog {
	conn := database.GetDbConnection()
	var blogs []Blog

	if len(id) == 1 {

		conn.First(&blogs, "id = ?", id[0])

		return blogs
	}

	conn.Find(&blogs)

	return blogs
}
