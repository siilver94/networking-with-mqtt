package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	//import the Paho Go MQTT library
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// 구독하기위한 변수
var xfileHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("Topic %s registered...\n", msg.Payload())
}

func main() {
	// 시그널 채널 생성
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// 토픽 생성
	const TOPIC = "kopo/ict/1916110385/time"

	// 클라이언트 옵션 생성 (브로커 설정)
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://test.mosquitto.org:1883")

	// 클라이언트 생성
	client := MQTT.NewClient(opts)

	// 구독
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	} else {
		fmt.Printf("Connected to server\n")
		client.Subscribe("kopo/ict/1916110385/time", 0, xfileHandler)
	}

	// 전송할 데이터와 데이터 폼
	payload := fmt.Sprintf("1916110385: 현재 시간은 %s 입니다", time.Now().String())

	// 오류 발생시 인터럽트가 아닌 오류로 인한 강제 종료 실행
	if token := client.Publish(TOPIC, 0, false, payload); token.Wait() && token.Error() != nil {
		os.Exit(1)
	}

	<-c

}
