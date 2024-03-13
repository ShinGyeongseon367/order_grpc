package ports

import "gitHub.com/ShinGyeongseon367/microservices/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
