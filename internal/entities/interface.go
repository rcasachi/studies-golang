package entities

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotalTransactions() (int, error)
}
