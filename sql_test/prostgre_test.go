package sql_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/bmizerany/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var postgresDB *gorm.DB

func initPostgres() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		"10.122.48.17", "5437", "cloudedge", "cloudedge", "cloudedge", "Asia/Shanghai")

	fmt.Println(dsn)
	postgresDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
}

type Student struct {
	Id    int `gorm:"column:id;not null;primaryKey;type:bigserial;commnet:'主键'"`
	Name  string
	Age   int
	Score string
}

func TestPostgres(t *testing.T) {
	initPostgres()
	if err := postgresDB.AutoMigrate(&Student{}); err != nil {
		logrus.Errorf("system AutoMigrate err: %s", err.Error())
		return
	}

	dt := postgresDB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"remark", "is_delete"}),
	}).Create(&Student{})

	// dt := postgresDB.Save(&Student{
	// 	Name:  "test2",
	// 	Age:   13,
	// 	Id:    "2",
	// 	Score: "60",
	// })

	fmt.Println(dt.Error)
}

func TestPostgres1(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)

	// 查询数据
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var uid int
		var name string
		var age int
		var score string
		err = rows.Scan(&name, &age, &uid, &score)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(uid)
		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(score)
	}
}

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=10.122.48.17 port=5437 user=cloudedge password=cloudedge dbname=cloudedge sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		return nil, err
	}

	fmt.Println(db, err)
	return db, err
}

func TestPostgres3(t *testing.T) {
	initPostgres()
	if err := postgresDB.AutoMigrate(&Student{}); err != nil {
		logrus.Errorf("system AutoMigrate err: %s", err.Error())
		return
	}

	dt := postgresDB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		// DoUpdates: clause.AssignmentColumns([]string{}),
		DoUpdates: clause.AssignmentColumns([]string{"name", "age", "id", "score"}),
	})

	students := make([]Student, 10)
	for i := 0; i < 10; i++ {
		students[i] = Student{
			Name:  "test555",
			Age:   i + 555,
			Id:    i + 10,
			Score: "555",
		}
	}

	students = append(students, Student{
		Name:  "test777",
		Age:   777,
		Id:    777,
		Score: "777",
	})
	dt = dt.Save(&students)

	fmt.Println(dt.Error)
}

func TestPostgres4(t *testing.T) {
	initPostgres()

	type GatewayAccessControl struct {
		Name  string
		Age   int
		Id    int `gorm:"column:id;not null;primaryKey;type:bigserial;commnet:'主键'"`
		Score string
	}
	if err := postgresDB.AutoMigrate(&GatewayAccessControl{}); err != nil {
		logrus.Errorf("system AutoMigrate err: %s", err.Error())
		return
	}

}

func TestPostgres5(t *testing.T) {
	initPostgres()
	if err := postgresDB.AutoMigrate(&Student{}); err != nil {
		logrus.Errorf("system AutoMigrate err: %s", err.Error())
		return
	}

	dt := postgresDB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		// DoUpdates: clause.AssignmentColumns([]string{}),
		DoUpdates: clause.AssignmentColumns([]string{"name", "age", "id", "score"}),
	})

	students := make([]Student, 10)
	for i := 0; i < 10; i++ {
		students[i] = Student{
			Name:  "test555",
			Age:   i + 555,
			Id:    i + 10,
			Score: "555",
		}
	}

	students = append(students, Student{
		Name:  "test777",
		Age:   777,
		Id:    777,
		Score: "777",
	})
	dt = dt.Save(&students)

	fmt.Println(dt.Error)
}
