package client

import (
	"log"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

type MQTTClient struct {
	client *client.Client
}

func NewMQTTClient(address, token string) *MQTTClient {
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			log.Println(err)
		},
	})

	// Connect to the MQTT Server.
	err := cli.Connect(&client.ConnectOptions{
		Network: "tcp",
		//Address:  `192.168.99.100:30514`,
		Address: address,
		//UserName: []byte("RASPBERRY_PI_DEMO_TOKEN"),
		UserName: []byte(token),
		// CleanSession is the Clean Session of the CONNECT Packet.
		CleanSession: true,
	})
	if err != nil {
		panic(err)
	}

	return &MQTTClient{client: cli}
}

func (this MQTTClient) PublishData(json string) error {
	return this.client.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte("v1/devices/me/telemetry"),
		//Message:   []byte(`{"temperature":21, "humidity":55.0, "active": false}`),
		Message:   []byte(json),
	})
}

func (this MQTTClient) PublishDeviceInfo(json string) error {
	return this.client.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte("v1/devices/me/attributes"),
		//Message:   []byte(`{"firmware_version":"1.0.1", "serial_number":"SN-001"}`),
		Message:   []byte(json),
	})
}

func (this *MQTTClient) Terminate() {
	this.client.Terminate()
}

func (this *MQTTClient) Disconnect() error {
	return this.client.Disconnect()
}