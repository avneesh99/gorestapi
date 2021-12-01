package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akto-api-security/gomiddleware"
	"github.com/gorilla/mux"
)

// book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	// Init Router
	r := mux.NewRouter()
	url := os.Getenv("AKTO_CONNECT_IP")
	fmt.Println("URL: " + url + ":8080")
	config, _ := gomiddleware.GetConfigFromDashboard(url + ":8080")
	kafkaWriter := gomiddleware.GetKafkaWriter(url+":9092", "akto.api.logs", 100, 1*time.Second)
	r.Use(gomiddleware.Middleware(kafkaWriter, config, 1111))

	books = append(books, Book{ID: "1", Isbn: "3223", Title: "Book 1", Author: &Author{
		Firstname: "Avneesh", Lastname: "Hota"}})
	books = append(books, Book{ID: "2", Isbn: "2323", Title: "Book 2", Author: &Author{
		Firstname: "Ankush", Lastname: "Jain"}})

	// Route endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/cars", getBooks).Methods("GET")
	r.HandleFunc("/api/auth/signin", signIn).Methods("GET")
	r.HandleFunc("/api/latest/meta-data/local-ipv4", asdf).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// init books var as slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(books)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode("wefwe")
}

func asdf(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("127.0.0.1"))
}
