package usecases

import "github.com/rcasachi/studies-golang/internal/entities"

type OrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entities.OrderRepositoryInterface
}

func NewCalculateFinalPrice(repository entities.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderRepository: repository,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entities.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
