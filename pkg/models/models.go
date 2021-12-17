package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Coffee struct {
	ID int 
	Name string
	Ingredients string
	Created time.Time
}