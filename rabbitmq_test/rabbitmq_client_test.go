package rabbitmq_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/feng-crazy/go-utils/rabbitmq"
)

func TestAmqp(t *testing.T) {
	queue := &rabbitmq.Queue{Name: "hdf.test"}
	// exchange := &Exchange{Name: "hdf.exchange.test"}

	msg := &rabbitmq.Message{
		Data: []byte("{\"seqno\":\"1563541319\",\"cmd\":\"44\",\"data\":{\"mid\":1070869}}"),
	}

	mq, err := rabbitmq.New(&rabbitmq.Config{
		Addr:         "amqp://guest:guest@10.122.48.78:5672/",
		ExchangeName: "hdf.test",
	})
	if err != nil {
		panic(err)
	}

	testCount := 1000
	wg := sync.WaitGroup{}

	startTime := time.Now()
	si := 0
	for ; si < testCount; si++ {
		err := mq.Push(queue, msg)
		if err != nil {
			panic(err)
		}
	}
	t.Logf("发送 %d 条数据, 耗时 %d 纳秒 \n", si, time.Since(startTime))

	startTime1 := time.Now()
	wg.Add(testCount)
	go func() {
		msgs, err := mq.Sub(queue)
		if err != nil {
			panic(err)
		}
		for msg := range msgs {
			fmt.Println(string(msg.Data))
			wg.Done()
		}
	}()

	wg.Wait()
	t.Logf("消费 %d 条数据, 耗时 %d 纳秒 \n", testCount, time.Since(startTime1))
}

func TestSubMqtt(t *testing.T) {
	ex := &rabbitmq.Exchange{
		Name:       "amq.topic",
		Kind:       rabbitmq.ExchangeTopic,
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

	queue := &rabbitmq.Queue{Name: "hdf.test", Key: "hdf.test"}
	// err = mq.Pub(queue, &rabbitmq.Message{Data: []byte("{\"hello\":\"world\"}")})
	// if err != nil {
	// 	panic(err)
	// }

	msgch, err := mq.Sub(queue)
	if err != nil {
		panic(err)
	}
	for msg := range msgch {
		fmt.Printf("%+v\n", msg)
		fmt.Printf("%s\n", string(msg.Data))
	}
}

func TestSubMqtt1(t *testing.T) {
	ex := &rabbitmq.Exchange{
		Name:       "amq.topic",
		Kind:       rabbitmq.ExchangeTopic,
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

	queue := &rabbitmq.Queue{Name: "hdf.test1111", Key: "hdf.test"}
	// err = mq.Pub(queue, &rabbitmq.Message{Data: []byte("{\"hello\":\"world\"}")})
	// if err != nil {
	// 	panic(err)
	// }

	msgch, err := mq.Sub(queue)
	if err != nil {
		panic(err)
	}
	for msg := range msgch {
		fmt.Printf("%+v\n", msg)
		fmt.Printf("%s\n", string(msg.Data))
	}
}

func TestPubMqtt(t *testing.T) {
	ex := &rabbitmq.Exchange{
		Name:       "amq.topic",
		Kind:       rabbitmq.ExchangeTopic,
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

	queue := &rabbitmq.Queue{Name: "hdf.test1", Key: "hdf.test1"}
	for {
		err = mq.Push(queue, &rabbitmq.Message{Data: []byte("{\"hello\":\"world\"}")})
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}

}
