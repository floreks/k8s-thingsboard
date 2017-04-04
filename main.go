package main

import (
	"os"
	"os/signal"

	mqtt "github.com/floreks/k8s-thingsboard/client"
	"github.com/floreks/k8s-thingsboard/service/dht"
)

func main() {
	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	// Create an MQTT Client.
	client := mqtt.NewMQTTClient("192.168.99.100:30514", "RASPBERRY_PI_DEMO_TOKEN")

	// Terminate the Client.
	defer client.Terminate()

	dht11Service := dht.NewDHT11Service(client)
	go dht11Service.ReadAndPublish()

	// Wait for receiving a signal.
	<-sigc

	// Disconnect the Network Connection.
	if err := client.Disconnect(); err != nil {
		panic(err)
	}
}