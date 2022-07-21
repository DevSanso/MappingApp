# README

## 의사 매칭 앱
---

어플로 통해 비대면으로 의사한테 상담을 받을수 있으며, ai 매칭을 통하여 현재상황에 적합한 의사를 자동으로 찾아주는 서비스 입니다.

## Probably Requirement Technology

---

- Rust
    - tungstenite
- Golang
    - torch(pytorch binding for go)
- Python
    - fastapi
- Typescript
    - express
- Dart
    - flutter
- Kafka
- Elasticsearch
- Protobuf
- Pytorch

## 요구사항(서버)

---

- 모든 서비스는 일정 시간의 시그널에 맞쳐 rw상태가 바끼는 only read/ only write 버퍼 방식을 채택하며,전송하고자하는 모든 메세지를 write 버퍼에 저장하며, read 버퍼로 자동으로 변환되면 해당 버퍼를 다른 서비스로 전송하도록 설계
- 모든 db 서비스는 write db 커넥터, read db 커넥터를 별도로 사용한다
- 데이터베이스의 직접 접근으로 읽는 서비스는 *db 서비스 그리고 검색엔진 서비스만 허용한다.
- 개인 정보랑 관련된 db 서비스는 개인 정보 암호화 키 생성 서비스를 이용하여 암호화키를 생성, 해당키로 정보를 암호화 한다.
- 정적 서비스는 다른 서비스로 연동이 안된다, 오로지 본인만의 기능만 작동해야한다.
- cache db  서비스의 db 구조는 현재 사용중인 서비스가 정의한다.
- 모든 서비스 설정 파일은 json 포멧을 쓰며, 프로세스 실행시 매개변수를 통하여 전달한다.

## RoadMap

---

![mappingApp 서비스 흐름.png](README%20ec939c255e68427787c9c27d382da869/mappingApp_%25EC%2584%259C%25EB%25B9%2584%25EC%258A%25A4_%25ED%259D%2590%25EB%25A6%2584.png)

### 모듈

---

- 서비스 공통 모듈 제작
    - [x]  rw 버퍼 모듈
    - [x]  프로세스 매개변수 설정 파싱 모듈
    - [x]  카프카 프로듀서, 컨슈머 객체 생성 wrapper 모듈

### 서비스

---

- [ ]  유저 정보 db 서비스 : (유저의 개인정보 보관 서비스)
- [ ]  medical history db 서비스 : (진료 기록 보관 서비스)
- [ ]  유저 access db 서비스 : (유저 oauth 보관 서비스)
- [ ]  make priacy key 서비스 : (유저 정보 암호화 키 생성 서비스)
- [ ]  검색 엔진 서비스 : ( db 정보 검색 서비스)
- [ ]  ai mapping 서비스 : (의사 매칭 서비스)
- [ ]  oauth 서비스 : (google,그리고 다른 oauth 로그인,회원 관련 처리 서비스)
- [ ]  cache db 서비스 : (임시 저장 인 메모리 db 서비스)
- [ ]  chat 서비스 : (환자 ,의사의 1대1 진료 채팅 관리 서비스)
- [ ]  정적 리소스 서버 : (html,markdown등 정적 파일들을 어플로 보내는 서비스)
- [ ]  채팅 서버  : (어플의 채팅 관련 요청을 처리하는 서버)
- [ ]  restapi 서비 : (정적 리소스, 채팅 서버의 처리외의 모든 처리 당담 서버)

### 제약사항

---

- 서비스 프로젝트 언어는 모두 go(1.18)로 통일
- restapi 서버는 파이썬 + fastapi 사용
- 채팅 서버 프로그래밍 언어는 rust 사용