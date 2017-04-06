// The MIT License
//
// Copyright (c) 2016 Sebastian Florek
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sensor

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/d2r2/go-dht"
)

const (
	GPIO_PIN        = 17
	RETRIES         = 5
	BOOST_PERF_FLAG = false
)

var (
	gpioPinOverride = GPIO_PIN
)

// TODO add doc
type DHT11Reader struct{}

// TODO add doc
func (DHT11Reader) ReadFromSensor() (*DHTResponse, error) {
	log.Printf("Reading from sensor %s on GPIO pin %d", dht.DHT11, gpioPinOverride)
	temp, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, gpioPinOverride, BOOST_PERF_FLAG,
		RETRIES)
	if err != nil {
		if strings.Contains(err.Error(), "C.dial_DHTxx_and_read") {
			return nil, fmt.Errorf("Could not read from sensor.")
		}

		return nil, err
	}

	return &DHTResponse{DHTTemperature{Temperature: temp}, DHTHumidity{Humidity: humidity}}, nil
}

type DHT11ReaderMock struct{}

// TODO add doc
func (DHT11ReaderMock) ReadFromSensor() (*DHTResponse, error) {
	temperature := random(10, 25)
	humidity := random(40, 80)
	return &DHTResponse{DHTTemperature{Temperature: temperature}, DHTHumidity{Humidity: humidity}}, nil
}

func random(min, max int) float32 {
	rand.Seed(time.Now().Unix())
	return float32(rand.Intn(max-min) + min)
}
