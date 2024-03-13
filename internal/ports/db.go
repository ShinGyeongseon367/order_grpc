package ports

import "gitHub.com/ShinGyeongseon367/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(ID string) (domain.Order, error)
	Save(*domain.Order) error
}
