package mqtt_client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-basic/uuid"
	log "github.com/sirupsen/logrus"
)

// MqttClient is parameters for Mqtt client.
type MqttClient struct {
	Qos        byte
	Retained   bool
	IP         string
	User       string
	Passwd     string
	CA         string
	Cert       string
	PrivateKey string
	Client     mqtt.Client
}

// newTLSConfig new TLS configuration.
// Only one side check. Mqtt broker check the cert from client.
func (mc *MqttClient) newTLSConfig() (*tls.Config, error) {
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(mc.CA)
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair(mc.Cert, mc.PrivateKey)
	if err != nil {
		return nil, err
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	// fmt.Println(cert.Leaf)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}, nil
}

func (mc *MqttClient) DisConnect() error {
	mc.Client.Disconnect(100)
	return nil
}

func SetReconnectingCb(mqtt.Client, *mqtt.ClientOptions) {
	log.Error("#################MQTT RECONNECTING!!")
}

// Connect connect to the Mqtt server.
func (mc *MqttClient) Connect() error {
	opts := mqtt.NewClientOptions().AddBroker(mc.IP).SetClientID(uuid.New()).
		SetCleanSession(false).SetConnectRetryInterval(5 * time.Second).
		SetReconnectingHandler(SetReconnectingCb)
	if mc.Cert != "" {
		tlsConfig, err := mc.newTLSConfig()
		if err != nil {
			return err
		}
		opts.SetTLSConfig(tlsConfig)
	} else {
		opts.SetUsername(mc.User)
		opts.SetPassword(mc.Passwd)
	}

	mc.Client = mqtt.NewClient(opts)
	// The token is used to indicate when actions have completed.
	if tc := mc.Client.Connect(); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}

	mc.Qos = 1          // At most 1 time
	mc.Retained = false // Not retained
	return nil
}

// Publish publish Mqtt message.
func (mc *MqttClient) Publish(topic string, payload interface{}) error {
	if tc := mc.Client.Publish(topic, mc.Qos, mc.Retained, payload); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	return nil
}

// Subscribe subsribe a Mqtt topic.
func (mc *MqttClient) Subscribe(topic string, onMessage mqtt.MessageHandler) error {
	if tc := mc.Client.Subscribe(topic, mc.Qos, onMessage); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	return nil
}

func NewMqttClient() *MqttClient {
	return &MqttClient{}
}
