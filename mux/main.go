package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Init Books variables as a slice Book Struct

var books []Book

// get All Books

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)

}

// get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params from request

	// Iterate through books

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})

}

// create Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {

	fmt.Println("Server Starting @ http://localhost:8080")
	// Init mux Router
	r := mux.NewRouter()

	//mock Data
	books = append(books, Book{ID: "1", Isbn: "45654", Title: "Hooked", Author: &Author{FirstName: "nil", LastName: "Erial"}})
	books = append(books, Book{ID: "2", Isbn: "45655", Title: "Hooked part-2", Author: &Author{FirstName: "Steve", LastName: "Erial"}})
	// Route Handlers /Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
