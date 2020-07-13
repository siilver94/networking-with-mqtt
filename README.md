# Networking With MQTT

## 프로젝트 소개

개발 과정에 있어서 필요에 의해 가장 중요할 수 있는 요소중 하나가 통신 입니다.

이번 과제는 다양한 통신 프로토콜중 **MQTT프로토콜**을 사용하여 프로토콜의 특징과 구축, 연동방법을 학습하는 것 입니다.
더 나아가 통신 프로토콜 뿐만아니라 **restAPI** 를 사용하여 자원을 주고 받을 것입니다.
이번 프로젝트에선 **Golang** 언어와 **MQTT** 프로토콜을 연동하여 통신 것이 목표입니다.


## 프로젝트 구조

**MQTT** 중계자 역할인 **Broker**를 설정하고 본인과 상대방과의 통신을 합니다.

작은 코드 공간이 필요하거나 네트워크 대역폭이 제한되는 원격 위치와의 연결을 위해 설계된 **MQTT 메세징 프로토콜** 과 **Golang** 을 연동하여 
통신환경을 구축 할 것입니다.

그 후, 네트워크 기반 아키텍처인 **RestAPI** 를 사용하여 데이터를 주고 받습니다.
**GoLang** 과 **MQTT**를 연동하여 간단한 메세지, 혹은 함수를 보내보고, **KAKAO MAP RestAPI** 에서 좌표를 얻고 그 좌표로 도로명과 주소를 구해 원하는 정보를 얻습니다.



### MQTT 란?
**MQTT**는 M2M, IOT를 위한 프로토콜로서, 최소한의 전력과 패킷량으로 통신하는 프로토콜입니다. 따라서 IOT와 모바일 어플리케이션 등의 통신에 매우 적합한 프로토콜입니다.
MQTT는 HTTP, TCP등의 통신과 같이 클라이언트-서버 구조로 이루어지는 것이 아닌, Broker, Publisher, Subscriber 구조로 이루어집니다.

![1_lKWgSNIYc1Pil5FFoAHMkA](https://user-images.githubusercontent.com/57824945/87271082-2426f100-c50d-11ea-9c56-443b2ff4034c.png)


Publisher는 Topic을 발행(publish) 하고, Subscriber는 Topic에 구독(subscribe)합니다. 
Broker는 이들을 중계하는 역할을 하며, 단일 Topic에 여러 Subscriber가 구독할 수 있기 때문에, 1:N 통신 구축에도 매우 유용합니다.

출처 : https://medium.com/@jspark141515/mqtt%EB%9E%80-314472c246ee


## 리뷰
