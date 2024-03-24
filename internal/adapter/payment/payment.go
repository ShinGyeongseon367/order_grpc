package payment

import (
	"context"

	"gitHub.com/ShinGyeongseon367/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices-proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentURL string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentURL, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	paymentClient := payment.NewPaymentClient(conn)

	return &Adapter{
		payment: paymentClient,
	}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(),
		&payment.CreatePaymentRequest{
			UserId:     order.CustomerID,
			OrderId:    order.ID,
			TotalPrice: order.TotalPrice(),
		})

	return err
}
