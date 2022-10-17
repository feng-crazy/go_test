package rabbitmq

import (
	"fmt"
	"testing"

	"github.com/feng-crazy/go-utils/rabbitmq"
)

func TestDataCenter(t *testing.T) {
	queueMeasurement := &rabbitmq.Queue{Name: "hdftest111.BatchMeasurementData"}

	ex := &rabbitmq.Exchange{
		Name:       "edge.amq.direct",
		Kind:       rabbitmq.ExchangeDirect,
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
		IsDeclare:  false,
	}
	mq, err := rabbitmq.New(&rabbitmq.Config{
		Addr:     "amqp://admin:admin@10.229.249.109:5672/",
		Exchange: ex,
	})
	if err != nil {
		panic(err)
	}

	msgs, err := mq.Sub(queueMeasurement)
	if err != nil {
		panic(err)
	}
	for msg := range msgs {
		fmt.Println(string(msg.Data))
	}

}
