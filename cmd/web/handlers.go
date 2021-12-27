package main

import (
	"errors"
	"fmt"
	"katarzynakawala/github.com/coffee-shop/pkg/models"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
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

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	ingredients := r.PostForm.Get("ingredients")

	errors := make(map[string]string)

	if strings.TrimSpace(name) == "" {
		errors["name"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(name) > 100 {
		errors["name"] = "This field is too long - max is 100 characters"
	}

	if strings.TrimSpace(ingredients) == "" {
		errors["ingredients"] = "This field cannot be blank"
	}

	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	id, err := app.coffees.Insert(name, ingredients)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/coffee/%d", id), http.StatusSeeOther)
}
