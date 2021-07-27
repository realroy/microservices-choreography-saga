package domain

type OrderStatus string

const (
	OrderPendingStatus  OrderStatus = "pending"
	OrderApprovedStatus             = "approved"
	OrderRejectedStatus             = "rejected"
	OrderPaidStatus                 = "paid"
)

type Order struct {
	Id         string      `json:"id"`
	Status     OrderStatus `json:"status"`
	Amount     float64     `json:"amount"`
	CustomerID string      `json:"customer_id"`
}

type OrderCreated struct {
	Id         string  `json:"id"`
	Amount     float64 `json:"amount"`
	CustomerID string  `json:"customer_id"`
}
