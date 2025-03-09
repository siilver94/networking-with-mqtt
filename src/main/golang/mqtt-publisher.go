package main

import (
    "fmt"
    "github.com/eclipse/paho.mqtt.golang"
    "time"
)

func main() {
    opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    for i := 1; i <= 5; i++ {
        text := fmt.Sprintf("Message %d", i)
        token := client.Publish("test/topic", 0, false, text)
        token.Wait()
        fmt.Println("Published:", text)
        time.Sleep(1 * time.Second)
    }

    client.Disconnect(250)
}
