package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"readinglist/internal/data"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		books := []data.Book{
			{
				ID:        1,
				Title:     "The Darkening of Tristam",
				Published: 1998,
				Pages:     300,
				Genres:    []string{"Fiction", "Thriller"},
				Rating:    4.5,
				Version:   1,
			},
			{
				ID:        2,
				Title:     "The Legecy of Deckard Cain",
				Published: 2007,
				Pages:     432,
				Genres:    []string{"Fiction", "Adventure"},
				Rating:    4.9,
				Version:   1,
			},
		}

		js, err := json.Marshal(books)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		js = append(js, '\n')
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if r.Method == "POST" {
		fmt.Fprintln(w, "Added a new book to the reading list.")
	}
}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			app.getBook(w, r)
		}
	case http.MethodPut:
		{
			app.updateBook(w, r)
		}
	case http.MethodDelete:
		{
			app.deleteBook(w, r)
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	book := data.Book{
		ID:        idInt,
		CreatedAt: time.Now(),
		Title:     "Echoes in the Darkness",
		Published: 2019,
		Pages:     300,
		Genres:    []string{"Fiction", "Thriller"},
		Rating:    4.5,
		Version:   1,
	}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Update the details of the book with ID: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Delete the details of the book with ID: %d", idInt)
}
