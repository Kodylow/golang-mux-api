package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get a Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, i := range books {
		if i.ID == params["id"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

//Create a Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var b Book
	_ = json.NewDecoder(r.Body).Decode(&b)
	b.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, b)
	json.NewEncoder(w).Encode(b)
}

//Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for ind, i := range books {
		if i.ID == params["ID"] {
			books = append(books[:ind], books[ind+1:]...)
			var b Book
			_ = json.NewDecoder(r.Body).Decode(&b)
			b.ID = params["ID"]
			books = append(books, b)
			json.NewEncoder(w).Encode(b)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for ind, i := range books {
		if i.ID == params["ID"] {
			books = append(books[:ind], books[ind+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "4487", Title: "Book One", Author: &Author{Firstname: "Kody", Lastname: "Low"}},
		Book{ID: "2", Isbn: "4497", Title: "Book Two", Author: &Author{Firstname: "Kody", Lastname: "Low"}})
	fmt.Println("Starting server...")
	//Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//Run Server
	log.Fatal(http.ListenAndServe(":8000", r))
}
