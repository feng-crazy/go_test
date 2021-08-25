package mqtt_client

import (
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

func TestMqttClient(t *testing.T) {
	mqttClient := &MqttClient{
		IP:         "ssl://10.122.48.78:1883",
		CA:         "myCA\\ca.crt",
		Cert:       "myCA\\client-crt.pem",
		PrivateKey: "myCA\\client-key.pem",
	}

	if err := mqttClient.Connect(); err != nil {
		logrus.Error(" mqtt connect failure, ", err)
	}

	_ = mqttClient.Subscribe("/hdf/test", msgHandle)

	select {}
}

func TestMqttClient1(t *testing.T) {
	mqttClient := &MqttClient{
		IP: "tcp://10.122.48.5:1883",
	}

	if err := mqttClient.Connect(); err != nil {
		logrus.Error(" mqtt connect failure, ", err)
	}

	_ = mqttClient.Subscribe("/hdf/test", msgHandle)

	time.Sleep(5 * time.Second)
	err := mqttClient.DisConnect()
	fmt.Println(err)
	select {}
}

func TestMqttClient2(t *testing.T) {
	mqttClient := &MqttClient{
		IP: "tcp://10.122.48.5:1883",
	}

	if err := mqttClient.Connect(); err != nil {
		logrus.Error(" mqtt connect failure, ", err)
	}
	// i := 0
	for {
		_ = mqttClient.Publish("/hdf/test", []byte("test"))
		time.Sleep(1 * time.Second)
	}

}

func msgHandle(client mqtt.Client, message mqtt.Message) {
	fmt.Println(message)
	fmt.Println(string(message.Payload()))
}

func TestMqttClient3(t *testing.T) {
	mqttClient := &MqttClient{
		IP: "tcp://10.122.48.5:1883",
	}

	if err := mqttClient.Connect(); err != nil {
		logrus.Error(" mqtt connect failure, ", err)
	}

	_ = mqttClient.Subscribe("/hdf/test", msgHandle1)
	_ = mqttClient.Subscribe("/hdf/test", msgHandle2)

	// time.Sleep(5*time.Second)
	// err := mqttClient.DisConnect()
	// fmt.Println(err)
	select {}
}

func msgHandle1(client mqtt.Client, message mqtt.Message) {
	fmt.Println("msgHandle1.................")
	fmt.Println(message)
	fmt.Println(string(message.Payload()))
}

func msgHandle2(client mqtt.Client, message mqtt.Message) {
	fmt.Println("msgHandle2....................")
	fmt.Println(message)
	fmt.Println(string(message.Payload()))
}

func TestMqttClient4(t *testing.T) {
	mqttClient := &MqttClient{
		IP: "tcp://localhost:1883",
	}

	if err := mqttClient.Connect(); err != nil {
		logrus.Error(" mqtt connect failure, ", err)
	}

	err := mqttClient.Subscribe("/hdf/test/+/111/+", msgHandle)

	fmt.Println(err)
	select {}
}
