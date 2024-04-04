package api

import (
	"strings"

	"gitHub.com/ShinGyeongseon367/microservices/order/internal/application/core/domain"
	"gitHub.com/ShinGyeongseon367/microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db             ports.DBPort
	paymentAdapter ports.PaymentPort
}

func NewApplication(db ports.DBPort, paymentAdapter ports.PaymentPort) *Application {
	return &Application{
		db:             db,
		paymentAdapter: paymentAdapter,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	paymentErr := a.paymentAdapter.Charge(&order)
	if paymentErr != nil {
		// st, _ := status.FromError(paymentErr)
		st := status.Convert(paymentErr)
		var allErrors []string
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				for _, violation := range t.GetFieldViolations() {
					allErrors = append(allErrors, violation.Description)
				}
			}
		}
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: strings.Join(allErrors, "\n"),
		}

		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetail, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetail.Err()
	}

	return order, nil
}
