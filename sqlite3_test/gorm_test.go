package sqlite3_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type EnergyInput struct {
	EventId    string `json:"eventId"`
	DeviceType string `json:"deviceType"`
	DeviceName string `json:"deviceName"`
	Timestamp  int64  `json:"timestamp"`
	ValidFlag  bool   `json:"ValidFlag"`

	Realvalue     float32 `json:"realvalue"`
	Offsetval     float32 `json:"offsetval"`
	Multiple      int     `json:"multiple"`
	Collectionrat int     `json:"collectionrat"`
}

type Device struct {
	ID uint `gorm:"primary_key"`

	DeviceName string `gorm:"uniqueIndex"`
	Value      int
	Expire     int

	Input []byte

	CreatedAt time.Time `gorm:"default:null"`
	UpdatedAt time.Time `gorm:"default:null"`
}

func (d *Device) Table() string {
	return "device"
}

func TestSqlite3Gorm(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	if db.HasTable(&Device{}) {
		db := db.AutoMigrate(&Device{})
		if db.Error != nil {
			fmt.Println(db.Error)
		}
	} else {
		db := db.CreateTable(&Device{})
		if db.Error != nil {
			fmt.Println(db.Error)
		}
	}

}

func TestCreate(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	dt := db.Table("devices").Create(&Device{
		DeviceName: "test1",
		Value:      1,
		Expire:     1,
		Input:      []byte{0, 1, 2},
	})
	fmt.Println(dt.Error)
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	devices := make([]Device, 0)
	db.Table("devices").Find(&devices)
	fmt.Println(devices)
}

func TestFindOne(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	device := Device{
		DeviceName: "test1",
	}
	db.Table("devices").Where(&device).Find(&device)
	fmt.Println(device)

	diff := time.Now().Unix() - device.UpdatedAt.Unix()
	fmt.Println(diff)
}

func TestDelete(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	err = db.Delete(&Device{DeviceName: "test2"}).Error
	fmt.Println(err)
}

func f(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	dt := db.Table("devices").Where(&Device{ID: 1}).Update(&Device{Expire: 3})
	fmt.Println(dt.Error)

}
