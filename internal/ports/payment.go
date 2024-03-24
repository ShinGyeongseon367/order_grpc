package ports

import "gitHub.com/ShinGyeongseon367/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
