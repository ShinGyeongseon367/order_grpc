package main

import (
	"log"

	"gitHub.com/ShinGyeongseon367/microservices/order/config"
	"gitHub.com/ShinGyeongseon367/microservices/order/internal/adapter/db"
	"gitHub.com/ShinGyeongseon367/microservices/order/internal/adapter/grpc"
	"gitHub.com/ShinGyeongseon367/microservices/order/internal/adapter/payment"
	"gitHub.com/ShinGyeongseon367/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
