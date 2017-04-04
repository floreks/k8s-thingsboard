package dht

import (
	"encoding/json"
	"log"
	"time"

	"github.com/floreks/k8s-thingsboard/client"
	"github.com/floreks/k8s-thingsboard/sensor/dht"
)

const defaultPublishInterval = time.Second

type DHT11Service struct {
	reader          sensor.DHTReader
	client          *client.MQTTClient
	publishInterval time.Duration

	dataChan chan string
}

func (this DHT11Service) Configure(publishInterval time.Duration) {
	this.publishInterval = publishInterval
}

func (this DHT11Service) ReadAndPublish() {
	go func() {
		for {
			response, err := this.reader.ReadFromSensor()
			if err != nil {
				log.Println(err.Error())
				return
			}

			responseJson, err := json.Marshal(*response)
			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Read from sensor: %s", responseJson)
			this.dataChan <- string(responseJson)
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		select {
		case dataJson := <-this.dataChan:
			log.Printf("Publishing data: %s", dataJson)
			err := this.client.PublishData(string(dataJson))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

}

func NewDHT11Service(client *client.MQTTClient) DHT11Service {
	return DHT11Service{reader: sensor.DHT11ReaderMock{}, client: client, publishInterval: defaultPublishInterval,
		dataChan: make(chan string)}
}
