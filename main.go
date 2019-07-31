package main

import(
	"net/http"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
// Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice book Struct
var books []Book 

// Get all books
func getBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(books)
}

// Get single book
func getBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get Params
	// Loop through books and find matched ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode(&Book{})

}

// Create new book
func createBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock ID
	books = append(books, book)
	json.NewEncoder(res).Encode(book)

}

// Update a book
func updateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get Params
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(req.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(res).Encode(book)
			return
		}
	}
	json.NewEncoder(res).Encode(books)
}

// Delete a book
func deleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get Params
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(books)
}
func main(){
	// Init Mux router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn:"3242343", Title:"Listen And Serve", Author: &Author {Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn:"1233343", Title:"Listen and Protect", Author: &Author {Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "3", Isbn:"5723453", Title:"ProtectAndServe", Author: &Author {Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "4", Isbn:"7654433", Title:"He Protec, He Attac", Author: &Author {Firstname: "John", Lastname: "Doe"}})

	// Route handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}