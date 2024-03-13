package db

import (
	"fmt"

	"gitHub.com/ShinGyeongseon367/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderId     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSource string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error : %v", openErr.Error())
	}

	migrateError := db.AutoMigrate(Order{}, OrderItem{})
	if migrateError != nil {
		return nil, fmt.Errorf("db migrate error : %v", migrateError)
	}

	return &Adapter{
		db: db,
	}, nil
}

func (a Adapter) Get(ID string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, ID)
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			Quantity:    orderItem.Quantity,
			ProduceCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}

	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProduceCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}
