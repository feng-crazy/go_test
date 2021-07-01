package rabbitmq_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/feng-crazy/go-utils/hmqtt"
	"github.com/sirupsen/logrus"
)

func TestRabbitMqMqtt(t *testing.T) {
	mqttClient := &hmqtt.MqttClient{
		IP:     "tcp://10.122.48.78:1883",
		User:   "guest",
		Passwd: "guest",
	}

	if err := mqttClient.Connect(); err != nil {
		logrus.Error(" mqtt connect failure, ", err)
	}

	i := 0
	for {
		i++
		data := "test  %d " + time.Now().String()
		data = fmt.Sprintf(data, i)
		err := mqttClient.Publish("hdf/test", []byte(data))

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("send success  test...", data)
		time.Sleep(time.Second)
	}
}
