package main

import (
    "fmt"
    "github.com/eclipse/paho.mqtt.golang"
)

func main() {
    opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    client.Subscribe("test/topic", 0, func(client mqtt.Client, msg mqtt.Message) {
        fmt.Printf("Received message: %s\n", msg.Payload())
    })

    select {} // Keep running
}
