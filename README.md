# gRPC 연습

## 프로젝트 설명
이 프로젝트는 주문 도메인을 기준으로 gRPC를 연습하기 위한 목적으로 만들어졌습니다. gRPC를 통해 주문 시스템의 효율적인 통신 구조를 설계하고 구현하는 방법을 탐색합니다.

## 시작하기
이 프로젝트를 시작하기 위해서는 먼저 해당 Repository를 복제(clone)해야 합니다. 아래 명령어를 사용하여 로컬 시스템에 복제할 수 있습니다.

```bash
git clone https://github.com/ShinGyeongseon367/order_grpc.git
```

## 사용 방법
프로젝트를 시작하기 전에 로컬 데이터베이스 설정이 필요합니다. 이 프로젝트는 MySQL을 사용하며, 로컬에서 프로젝트를 시작하기 위한 환경 변수 설정과 함께 아래 명령어를 실행해야 합니다.
```bash
DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3306)/order \ APPLICATION_PORT=3000 \ 
ENV=development \ 
go run cmd/main.go
```


## 빌드 및 테스트
현재 이 프로젝트의 빌드 및 테스트에 대한 부분은 개발 중에 있습니다. 추후 업데이트 예정입니다.

## 기여하기
이 프로젝트에 기여를 희망하시는 분은 메일을 통해 연락 주시기 바랍니다. 
프로젝트에 기여하는 방법에 대한 자세한 정보는 추후 제공될 예정입니다.

## 작성자
- Shin Gyeongseon - [ShinGyeongseon367](https://github.com/ShinGyeongseon367)

---
**감사의 글**  
이 프로젝트는 gRPC 학습 과정에서 얻은 지식을 바탕으로 만들어졌습니다. gRPC 커뮤니티와 모든 기여자에게 감사드립니다.
