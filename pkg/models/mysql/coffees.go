package mysql

import (
	"database/sql"
	"katarzynakawala/github.com/coffee-shop/pkg/models"
)


type CoffeeModel struct {
	DB *sql.DB
}

func (m *CoffeeModel) Insert(name, ingredients string) (int, error) {
	return 0, nil
}

func (m *CoffeeModel) Get(id int) (*models.Coffee, error) {
	return nil, nil
}
func (m *CoffeeModel) Latest() ([]*models.Coffee, error) {
	return nil, nil
}