package mysql

import (
	"database/sql"
	"katarzynakawala/github.com/coffee-shop/pkg/models"
)


type CoffeeModel struct {
	DB *sql.DB
}

func (m *CoffeeModel) Insert(name, ingredients string) (int, error) {
	stmt := `INSERT INTO coffees (name, ingredients, created)
	VALUES(?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, name, ingredients)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *CoffeeModel) Get(id int) (*models.Coffee, error) {
	return nil, nil
}
func (m *CoffeeModel) Latest() ([]*models.Coffee, error) {
	return nil, nil
}