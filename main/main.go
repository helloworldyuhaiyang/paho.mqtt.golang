package main

import (
	"crypto/tls"
	"fmt"
	mqtt "github.com/helloworldyuhaiyang/paho.mqtt.golang"
	"time"
)

func main() {
	opts := mqtt.NewClientOptions().
		AddBroker("quic://106.54.175.237:8861").
		SetClientID("sample").
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	c.Subscribe("/test", 0x01, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("%s recv:%s\n", message.Topic(), message.Payload())
	})

	c.Publish("/test", 0x01, false, "nihao")
	c.Publish("/test", 0x01, false, "hello")

	time.Sleep(time.Second * 2)
}
