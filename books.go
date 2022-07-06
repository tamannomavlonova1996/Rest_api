package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var books []Book

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//json.NewEncoder(w).Encode(books)

}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return

		}

	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()
	books = append(books, Book{
		ID:    "1",
		Title: "Война и Мир",
		Author: &Author{
			Firstname: "Лев",
			Lastname:  "Толстой",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Title: "Престпудение и наказание",
		Author: &Author{
			Firstname: "Фёдор",
			Lastname:  "Достоевский",
		},
	})
	r.HandleFunc("/books", getBooks).Methods("Get")
	r.HandleFunc("/books/{id}", getBook).Methods("Get")
	r.HandleFunc("/books", createBook).Methods("Post")
	r.HandleFunc("/books/{id}", updateBook).Methods("Put")
	r.HandleFunc("/books/{id}", deleteBook).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}
