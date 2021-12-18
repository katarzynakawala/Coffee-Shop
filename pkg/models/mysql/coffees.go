package mysql

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, name, ingredients, created FROM coffees
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)
	s := &models.Coffee{}

	err := row.Scan(&s.ID, &s.Name, &s.Ingredients, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}


func (m *CoffeeModel) Latest() ([]*models.Coffee, error) {
	stmt := `SELECT id, name, ingredients, created FROM coffees
	ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	coffees := []*models.Coffee{}

	for rows.Next() {
		s := &models.Coffee{}
		err = rows.Scan(&s.ID, &s.Name, &s.Ingredients, &s.Created)
		if err != nil {
			return nil, err
		}
		coffees = append(coffees, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return coffees, nil
}