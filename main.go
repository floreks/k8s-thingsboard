package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/floreks/k8s-thingsboard/client"
	"github.com/floreks/k8s-thingsboard/sensor/dht"
	"github.com/floreks/k8s-thingsboard/service/dht"
	"github.com/spf13/pflag"
)

var (
	argThingsboardHost = pflag.String("thingsboard-host", "", "The address of thingsboard to report on")
	argDeviceToken = pflag.String("token", "", "Token of the device to report on")
)

func main() {
	// Set logging out to standard console out.
	log.SetOutput(os.Stdout)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if len(*argThingsboardHost) == 0 || len(*argDeviceToken) == 0 {
		log.Fatal("Thingsboard host and device token have to be defined.")
	}

	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	// Create an MQTT Client.
	client := mqtt.NewMQTTClient(*argThingsboardHost, *argDeviceToken)

	// Terminate the Client.
	defer client.Terminate()

	// Change to sensor.DHT11 to read from an actual dht sensor.
	dhtReader, _ := sensor.NewDHTReader(sensor.DHT11)
	dht11Service := dht.NewDHT11Service(client, dhtReader)
	go dht11Service.ReadAndPublish()

	// Wait for receiving a signal.
	<-sigc

	// Disconnect the Network Connection.
	if err := client.Disconnect(); err != nil {
		panic(err)
	}
}
