package map_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/devfeel/mapper"
)

type (
	User struct {
		Name string
		Age  int
		Id   string `mapper:"_id"`
		AA   string `json:"Score"`
		Time time.Time
	}

	Student struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Score string
	}

	Teacher struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Level string
	}
)

func init() {
	mapper.Register(&User{})
	mapper.Register(&Student{})
}

func TestMapper(t *testing.T) {
	user := &User{}
	userMap := &User{}
	teacher := &Teacher{}
	student := &Student{Name: "test", Age: 10, Id: "testId", Score: "100"}
	valMap := make(map[string]interface{})
	valMap["Name"] = "map"
	valMap["Age"] = 10
	valMap["_id"] = "x1asd"
	valMap["Score"] = 100
	valMap["Time"] = time.Now()

	_ = mapper.Mapper(student, user)
	_ = mapper.AutoMapper(student, teacher)
	_ = mapper.MapperMap(valMap, userMap)

	fmt.Println("student:", student)
	fmt.Println("user:", user)
	fmt.Println("teacher", teacher)
	fmt.Println("userMap:", userMap)
}
