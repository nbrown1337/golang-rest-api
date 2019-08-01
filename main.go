package main

import(
	"net/http"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
// Datasets Struct (Model)
type Datasets struct {
	ID string `json:"id"`
	Label string `json:"label"`
	Data []string `json:"data"`
}

type Label struct {
	Labels []string `json:"labels"`
}

type Combined struct {
	Datasets []Datasets
	Label []Label
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice book Struct
var books []Datasets 
var labels []Label
var combined []Combined

// Get all books
func getDatasets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(combined)
}

// Get single book
func getDataset(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get Params
	// Loop through books and find matched ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode(&Datasets{})

}

// Create new book
func createDatasets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var book Datasets
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock ID
	books = append(books, book)
	json.NewEncoder(res).Encode(book)

}

// Update a book
func updateDatasets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get Params
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Datasets
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
func deleteDatasets(res http.ResponseWriter, req *http.Request) {
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
	books = append(books, Datasets{ID: "1", Label:"Subs", Data:[]string{"14", "15", "21", "0", "12", "4", "1"}})
	books = append(books, Datasets{ID: "2", Label:"Views", Data:[]string{"8", "12", "11", "4", "1", "14", "8"}})
	books = append(books, Datasets{ID: "3", Label:"Likes", Data:[]string{"11", "15", "5", "15", "42", "4", "14"}})
	books = append(books, Datasets{ID: "4", Label:"Dislikes", Data:[]string{"4", "5", "1", "10", "32", "2", "12"}})
	//books = append(books, Datasets{Labels:[]string{"1", "2", "3", "4", "5"}})
	labels = append(labels, Label{Labels:[]string{"1", "2", "3", "4", "5"}})
	combined = append(combined, Combined{books,labels});

	// Route handlers / Endpoints
	router.HandleFunc("/api/datasets", getDatasets).Methods("GET")
	router.HandleFunc("/api/dataset/{id}", getDataset).Methods("GET")
	router.HandleFunc("/api/datasets", createDatasets).Methods("POST")
	router.HandleFunc("/api/datasets/{id}", updateDatasets).Methods("PUT")
	router.HandleFunc("/api/datasets/{id}", deleteDatasets).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}