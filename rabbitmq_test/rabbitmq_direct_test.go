package rabbitmq_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/feng-crazy/go-utils/rabbitmq"
)

func TestDirectSubMqtt(t *testing.T) {
	ex := &rabbitmq.Exchange{
		Name:       "amq.direct",
		Kind:       rabbitmq.ExchangeDirect,
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
		IsDeclare:  false,
	}

	mq, err := rabbitmq.New(&rabbitmq.Config{
		Addr:     "amqp://guest:guest@10.122.48.78:5672/",
		Exchange: ex,
	})
	if err != nil {
		panic(err)
	}

	queue := &rabbitmq.Queue{Name: "hdf.test.queue0", Key: "hdf.test.direct"}
	// err = mq.Pub(queue, &rabbitmq.Message{Data: []byte("{\"hello\":\"world\"}")})
	// if err != nil {
	// 	panic(err)
	// }

	msgch, err := mq.Sub(queue)
	if err != nil {
		panic(err)
	}
	for msg := range msgch {
		// fmt.Printf("%+v\n", msg)
		fmt.Printf("queue0 %s\n", string(msg.Data))
	}
}

func TestDirectSubMqtt1(t *testing.T) {
	ex := &rabbitmq.Exchange{
		Name:       "amq.direct",
		Kind:       rabbitmq.ExchangeDirect,
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
		IsDeclare:  false,
	}

	mq, err := rabbitmq.New(&rabbitmq.Config{
		Addr:     "amqp://guest:guest@10.122.48.78:5672/",
		Exchange: ex,
	})
	if err != nil {
		panic(err)
	}

	queue := &rabbitmq.Queue{Name: "hdf.test.queue1", Key: "hdf.test.direct"}
	// err = mq.Pub(queue, &rabbitmq.Message{Data: []byte("{\"hello\":\"world\"}")})
	// if err != nil {
	// 	panic(err)
	// }

	msgch, err := mq.Sub(queue)
	if err != nil {
		panic(err)
	}
	for msg := range msgch {
		// fmt.Printf("%+v\n", msg)
		fmt.Printf("queue1 %s\n", string(msg.Data))
	}
}

func TestDirectPubMqtt(t *testing.T) {
	ex := &rabbitmq.Exchange{
		Name:       "amq.direct",
		Kind:       rabbitmq.ExchangeDirect,
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
		IsDeclare:  false,
	}

	mq, err := rabbitmq.New(&rabbitmq.Config{
		Addr:     "amqp://guest:guest@10.122.48.78:5672/",
		Exchange: ex,
	})
	if err != nil {
		panic(err)
	}

	i := 0
	for {
		i++
		data := "test  %d " + time.Now().String()
		data = fmt.Sprintf(data, i)
		err = mq.Pub("hdf.test.direct", &rabbitmq.Message{Data: []byte(data)})
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
		time.Sleep(time.Second)
	}

}
