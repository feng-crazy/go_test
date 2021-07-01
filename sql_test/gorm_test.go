package sql_test

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func (d *Device) TableName() string {
	return "device"
}

func TestSqlite3Gorm(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Device{})
	if err != nil {
		fmt.Println(db.Error)
	}
}

func TestCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Logger.LogMode(1)

	dt := db.Table("device").Save(&Device{
		DeviceName: "test1",
		Value:      1,
		Expire:     1,
		Input:      []byte{0, 1, 2},
	})
	fmt.Println(dt.Error)
}

func TestReplace(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Logger.LogMode(1)
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	devices := make([]Device, 0)
	db.Table("devices").Find(&devices)
	fmt.Println(devices)
}

func TestFindOne(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	device := Device{
		DeviceName: "test1",
	}
	db.Table("devices").Where(&device).Find(&device)
	fmt.Println(device)

	diff := time.Now().Unix() - device.UpdatedAt.Unix()
	fmt.Println(diff)
}

func TestDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.Delete(&Device{DeviceName: "test2"}).Error
	fmt.Println(err)
}

func TestUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	dt := db.Table("devices").Where(&Device{ID: 1}).Updates(&Device{Expire: 3})
	fmt.Println(dt.Error)

}

func TestBitchUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	type DTest struct {
		ID   uint   `gorm:"primarykey"`
		Name string `json:"name" gorm:"uniqueIndex"`
		Dr   string `json:"dr"`
		P1   string `json:"p1"`
		P2   int    `json:"p2"`
		P3   bool   `json:"p3"`
	}
	err = db.AutoMigrate(&DTest{})
	if err != nil {
		fmt.Println(db.Error)
	}

	arg := []DTest{
		{
			Name: "dev1",
			Dr:   "0",
			P1:   "",
			P2:   0,
			P3:   false,
		},
		{
			Name: "dev2",
			Dr:   "0",
			P1:   "",
			P2:   0,
			P3:   false,
		}}
	fmt.Println(arg)

	var dt *gorm.DB
	// dt = db.Table("d_tests").CreateInBatches(&arg, 3)
	// if dt.Error != nil {
	// fmt.Println(dt.Error)
	// 	return
	// }

	dt = db.Exec("replace into d_tests (name,dr,p1,p2,p3) values ('dev1','4','p1 test1', 1, true),('dev2','5','p1 test2', 2, true)")
	fmt.Println(dt.Error)

}

func TestBitchDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	type DTest struct {
		ID   uint   `gorm:"primarykey"`
		Name string `json:"name" gorm:"uniqueIndex"`
		Dr   string `json:"dr"`
		P1   string `json:"p1"`
		P2   int    `json:"p2"`
		P3   bool   `json:"p3"`
	}

	dt := db.Table("d_tests").Where("p3 = ?", 1).Delete(DTest{})
	fmt.Println(dt.Error)
}
