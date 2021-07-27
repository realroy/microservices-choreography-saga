package application

import "order/domain"

type OrderRepositroy interface {
	Create(customerID string) (*domain.Order, error)
	UpdateByID(id string, data *domain.Order) (*domain.Order, error)
}

func CreateOrder(customerID string, r OrderRepositroy) (*domain.Order, error) {
	return r.Create(customerID)
}

func ApproveOrder(orderID string, r OrderRepositroy) (*domain.Order, error) {
	return r.UpdateByID(orderID, &domain.Order{Status: domain.OrderApprovedStatus})
}

func RejectOrder(orderID string, r OrderRepositroy) (*domain.Order, error) {
	return r.UpdateByID(orderID, &domain.Order{Status: domain.OrderRejectedStatus})
}
