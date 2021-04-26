package eventbus_test

import (
	"fmt"
	"testing"

	"github.com/asaskevich/EventBus"
)

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func TestEventBus1(t *testing.T) {
	bus := EventBus.New()
	_ = bus.Subscribe("main:calculator", calculator)
	bus.Publish("main:calculator", 20, 40)
	_ = bus.Unsubscribe("main:calculator", calculator)
}
