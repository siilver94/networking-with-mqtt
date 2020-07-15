// 통신 과정
// 과정1. 카카오 REST API를 이용해 지역명으로 좌표를 구하고
// 과정2. 카카오 REST API를 이용해 좌표로 주소와 도로명 주소를 구해서
// 과정3. mqtt를 이용하여 통신

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// 구독하기위한 변수
var xfileHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("Topic %s registered...\n", msg.Payload())
}

// 주소와 도로명 주소를 얻기 위한 구조체
// 주소와 도로명 이외에는 걸러진다.
type Geo2AddrResult struct {
	Meta struct {
		TotalCount int `json:"total_count"`
	} `json:"meta"`
	Documents []struct {
		RoadAddress struct {
			AddressName string `json:"address_name"`
		} `json:"road_address"`
		Address struct {
			AddressName string `json:"address_name"`
		} `json:"address"`
	} `json:"documents"`
}

// 좌표를 얻기 위한 구조체
// 좌표 정보 이외에는 다 걸러진다.
type addr2GeoResult struct {
	Meta struct {
		TotalCount    int    `json:"total_count"`
		PageableCount string `json:"pageable_count"`
		IsEnd         bool   `json:"is_end"`
	} `json:"meta"`
	Documents []struct {
		AddressName string `json:"address_name"`
		Y           string `json:"y"`
		X           string `json:"x"`
	} `json:"documents"'`
}

// 과정1. 지역명을 매개변수로 좌표 값을 얻어낸다.
func test1(input string) (x string, y string) {
	// GET 방식 통신 설정
	req, _ := http.NewRequest("GET", "https://dapi.kakao.com/v2/local/search/address.json", nil)

	// HTML 헤더 설정
	req.Header.Set("Content-Type", "charset=UTF-8")
	req.Header.Add("Authorization", "KakaoAK 9861434039c320a19371333c565f7c36")

	// HTML 쿼리 정보 설정 & 추가 & 인코딩
	q := req.URL.Query()          // 설정
	q.Add("query", input)         // 추가
	req.URL.RawQuery = q.Encode() // 인코딩

	// 통신 개시
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	// 좌표를 저장할 구조체 생성
	var r addr2GeoResult
	err := json.NewDecoder(resp.Body).Decode(&r)
	_ = err

	// 좌표값 반환
	x = r.Documents[0].X
	y = r.Documents[0].Y

	return
}

// 과정2. 좌표를 매개변수로 주소와 도로명 주소를 얻어낸다.
func test2(x string, y string) (res1 string, res2 string) {
	// GET 방식 통신 설정
	req, _ := http.NewRequest("GET", "https://dapi.kakao.com/v2/local/geo/coord2address.json", nil)

	// HTML 헤더 설정
	req.Header.Set("Content-Type", "charset=UTF-8")
	req.Header.Add("Authorization", "KakaoAK 9861434039c320a19371333c565f7c36")

	// HTML 쿼리 정보 설정 & 추가 & 인코딩
	q := req.URL.Query()          // 설정
	q.Add("x", x)                 // 추가
	q.Add("y", y)                 // 추가
	req.URL.RawQuery = q.Encode() // 인코딩

	// 통신 개시
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	// 주소와 도로명 주소를 저장할 구조체 생성
	var r Geo2AddrResult
	json.NewDecoder(resp.Body).Decode(&r)
	fmt.Println(r) // 확인용

	// 주소와 도로명 주소 반환
	res1 = r.Documents[0].RoadAddress.AddressName
	res2 = r.Documents[0].Address.AddressName

	return
}

// 과정3. 주소와 도로명 주소를 매개변수로 이 정보를 mqtt를 통해 전송한다.
func send(add1, add2 string) {
	// 시그널 채널 생성
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// 토픽 생성
	const TOPIC = "kopo/ict/1916110385/kakao"

	// 연결할 브로커 설정
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://test.mosquitto.org:1883")

	// 클라이언트 생성
	client := MQTT.NewClient(opts)

	// 브로커 측과 구독
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	} else {
		client.Subscribe("kopo/ict/1916110385/kakao", 0, xfileHandler)
	}

	// 전송할 데이터와 데이터 폼
	payload := fmt.Sprintf("1916110385: %s %s", add1, add2)

	// 오류 발생시 인터럽트가 아닌 오류로 인한 강제 종료 실행
	if token := client.Publish(TOPIC, 0, false, payload); token.Wait() && token.Error() != nil {
		os.Exit(1)
	}

	<-c
}

func main() {
	var add string
	fmt.Scanln(&add)
	send(test2(test1(add)))
}
