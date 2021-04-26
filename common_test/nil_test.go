package common

import (
	"log"
	"testing"
)

type Base interface {
	do()
}

type App struct {
}

func TestNil(t *testing.T) {
	var base Base
	base = GetApp()

	log.Println(base)
	log.Println(base == nil)
}

func GetApp() *App {
	return nil
}
func (a *App) do() {}

type MyErr struct {
	Msg string
}

func TestErrNil(t *testing.T) {
	var e error
	log.Println(e == nil)
	e = GetErr()
	log.Println(e == nil)

	err := GetErr()
	log.Println(err == nil)
}

func GetErr() *MyErr {
	return nil
}

func (m *MyErr) Error() string {
	return "脑子进煎鱼了"
}
