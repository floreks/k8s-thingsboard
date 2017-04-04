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
)

type SensorType int

const (
	DHT11 SensorType = iota + 1
	DHT11_MOCK
)

// TODO add doc
type DHTResponse struct {
	DHTTemperature
	DHTHumidity
}

// TODO add doc
type DHTTemperature struct {
	// Temperature
	Temperature float32 `json:"temperature"`
}

// TODO add doc
type DHTHumidity struct {
	// Humidity
	Humidity float32 `json:"humidity"`
}

// TODO add doc
type DHTReader interface {
	ReadFromSensor() (*DHTResponse, error)
}

func NewDHTReader(sensorType SensorType) (DHTReader, error) {
	switch sensorType {
	case DHT11:
		return DHT11Reader{}, nil
	case DHT11_MOCK:
		return DHT11ReaderMock{}, nil
	default:
		return nil, fmt.Errorf("Unsupported sensor type provided.")
	}
}