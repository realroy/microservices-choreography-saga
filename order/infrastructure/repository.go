package infrastructure

import (
	"order/domain"
	"strconv"
)

var orderId = float64(1)
var orders = make(map[string]domain.Order)

type OrderRepositroy struct{}

func (r OrderRepositroy) Create(customerID string) (*domain.Order, error) {
	order := &domain.Order{
		Id:         strconv.Itoa(int(orderId)),
		Status:     domain.OrderPendingStatus,
		Amount:     10,
		CustomerID: customerID,
	}

	orders[order.Id] = *order

	orderId += 1

	return order, nil
}

func (r OrderRepositroy) FindMany(condition map[string]interface{}) ([]domain.Order, error) {
	result := make([]domain.Order, 0)

	for _, order := range orders {
		result = append(result, order)
	}

	return result, nil
}

func (r OrderRepositroy) UpdateByID(id string, data *domain.Order) (*domain.Order, error) {
	oldOrder := orders[id]

	newOrder := &domain.Order{
		Id:         id,
		Status:     data.Status,
		CustomerID: oldOrder.CustomerID,
		Amount:     oldOrder.Amount,
	}

	orders[id] = *newOrder

	return newOrder, nil
}
