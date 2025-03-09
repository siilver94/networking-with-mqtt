#  Golang 기반 MQTT 통신 프로젝트  

##  프로젝트 소개  
**"Golang과 MQTT를 활용한 실시간 메시징 및 REST API 연동"**  

이 프로젝트는 **Golang**을 사용하여 **MQTT 프로토콜**을 기반으로 한 통신 환경을 구축하는 것을 목표로 합니다.  
또한, **REST API(Kakao Map API)**를 활용하여 데이터를 주고받으며, 실시간 메시지 전송 및 위치 데이터를 처리하는 방식을 학습합니다.  

---

##  주요 기술 및 개념  
✔ **Golang** – 빠르고 가벼운 네트워크 프로그래밍 지원  
✔ **MQTT** – IoT 및 M2M을 위한 경량 메시징 프로토콜  
✔ **Mosquitto (MQTT Broker)** – 메시지 중계를 담당하는 브로커  
✔ **REST API** – Kakao Map API를 사용하여 좌표 및 주소 변환  

---

##  프로젝트 구조  
 **구성 요소 설명**  
- **MQTT Broker**: Mosquitto를 사용하여 메시지 중계  
- **Publisher**: 특정 Topic에 메시지를 발행  
- **Subscriber**: Topic을 구독하여 메시지를 수신  
- **REST API 연동**: Kakao Map API를 활용하여 좌표 변환  

---

##  MQTT란?  
**MQTT(Message Queuing Telemetry Transport)**는 M2M 및 IoT 환경을 위한 경량 메시징 프로토콜입니다.  
TCP/IP 위에서 동작하며, 저전력 환경에서 효율적인 통신을 제공하는 것이 특징입니다.  

MQTT의 주요 구성 요소:  
- **Broker** 🏢: 메시지를 중계하는 서버 (예: Mosquitto)  
- **Publisher** 📤: 메시지를 특정 Topic에 발행하는 클라이언트  
- **Subscriber** 📥: Topic을 구독하여 메시지를 수신하는 클라이언트  
