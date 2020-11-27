# MQTT

Simple mqtt package

## Usage

This simple example code.

```golang
package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	mqtt "github.com/kimbugs/simple-connector/mqtt"
)

const (
	url      = "localhost:1883"
	clientID = "test2"
	id       = ""
	password = ""

	topic = "/test"
)

var payload = "ABCD"

var handler = func(payload []byte) {
	fmt.Println(string(payload))
}

func main() {
	c := mqtt.NewClient(url, clientID, id, password)
	c.OnConnect(topic, func(client MQTT.Client, msg MQTT.Message) {
		go handler(msg.Payload())
	})

	if err := c.Connet(); err != nil {
		fmt.Println(err.Error())
		panic("Mqtt not connected!!")
	}

	for {
		t := time.NewTicker(time.Second * 1)
		for tc := range t.C {
			fmt.Println(tc.String())
			c.Publish(topic, payload)
		}
	}

}

```
