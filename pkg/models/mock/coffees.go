package mock

import (
	"katarzynakawala/github.com/coffee-shop/pkg/models"
	"time"
)

var mockCoffee = &models.Coffee{
	ID:          1,
	Name:        "Coffee",
	Ingredients: "coffee and water",
	Created:     time.Now(),
}

type CoffeeModel struct{}

func (m *CoffeeModel) Insert(name, ingredients string) (int, error) {
	return 2, nil
}

func (m *CoffeeModel) Get(id int) (*models.Coffee, error) {
	switch id {
	case 1:
		return mockCoffee, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *CoffeeModel) Latest() ([]*models.Coffee, error) {
	return []*models.Coffee{mockCoffee}, nil
}