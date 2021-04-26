package mqtt_client

import (
	"fmt"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

func TestMqttClient(t *testing.T) {
	mqttClient := &MqttClient{
		IP: "tcp://10.229.249.68:1883",
	}

	if err := mqttClient.Connect(); err != nil {
		logrus.Error(" mqtt connect failure, ", err)

	}

	_ = mqttClient.Subscribe("/hdf/test", msgHandle)

	select {}
}

func msgHandle(client mqtt.Client, message mqtt.Message) {
	fmt.Println(message)
	fmt.Println(string(message.Payload()))
}
