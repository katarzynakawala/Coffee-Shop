package main

import (
	"errors"
	"fmt"
	"katarzynakawala/github.com/coffee-shop/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	
	c, err := app.coffees.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Coffees: c,
	})
}

func (app *application) displayCoffee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	c, err := app.coffees.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

app.render(w, r, "display.page.tmpl", &templateData{
	Coffee: c,
	})	
}

func (app *application) createCoffeeForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createCoffee(w http.ResponseWriter, r *http.Request) {

	name := "black coffee"
	ingredients := "water coffee"

	id, err := app.coffees.Insert(name, ingredients)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/coffee/id=%d", id), http.StatusSeeOther)
}