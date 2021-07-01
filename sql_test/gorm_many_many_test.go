package sql_test

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func TestCreateTable(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db = db.AutoMigrate(&User{})
	fmt.Println(db.Error)
}

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	langEN := Language{Name: "EN"}
	db.Table("user_languages").Create(&langEN)
	langCN := Language{Name: "CN"}
	db.Table("user_languages").Create(&langCN)

	u1 := &User{
		Name: "user1",

		Languages: []Language{
			langEN,
			langCN,
		},
	}
	db.Create(u1)

	u2 := &User{
		Name: "user2",

		Languages: []Language{
			langCN,
		},
	}
	db.Create(u2)
}
