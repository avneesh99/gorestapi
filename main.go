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
	kafka_url := os.Getenv("AKTO_CONNECT_IP") + ":9092"
	fmt.Println(kafka_url)
	// dashboard_url := "http://" + os.Getenv("AKTO_DASHBOARD_IP") + ":8080"
	// fmt.Println(dashboard_url)
	// config, _ := gomiddleware.GetConfigFromDashboard(dashboard_url)
	// fmt.Println(config.BlackList)
	kafkaWriter := gomiddleware.GetKafkaWriter(kafka_url, "akto.api.logs", 100, 1*time.Second)
	r.Use(gomiddleware.Middleware(kafkaWriter, 1000000))

	books = append(books, Book{ID: "1", Isbn: "3223", Title: "Book 1", Author: &Author{
		Firstname: "Avneesh", Lastname: "Hota"}})
	books = append(books, Book{ID: "2", Isbn: "2323", Title: "Book 2", Author: &Author{
		Firstname: "Ankush", Lastname: "Jain"}})

	// Route endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/cars", getBooks).Methods("POST")
	r.HandleFunc("/api/toys", getBooks).Methods("POST")
	r.HandleFunc("/api/games", getBooks).Methods("POST")
	r.HandleFunc("/api/football", getBooks).Methods("POST")
	r.HandleFunc("/api/cricket", getBooks).Methods("POST")
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
	json.NewEncoder(w).Encode(books[0])
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
