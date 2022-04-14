package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/akto-api-security/gomiddleware"
	"github.com/gorilla/mux"
)

// book struct (Model)
type Book struct {
	ID        string  `json:"id"`
	Isbn      string  `json:"isbn"`
	Title     string  `json:"title"`
	Author    *Author `json:"author"`
	Timestamp int64   `json:"timestamp"`
}

type Car struct {
	ID     string            `json:"id"`
	Number string            `json:"number"`
	Model  string            `json:"model"`
	A      string            `json:"a"`
	B      string            `json:"b"`
	C      string            `json:"c"`
	D      string            `json:"d"`
	E      string            `json:"e"`
	F      string            `json:"f"`
	G      string            `json:"g"`
	H      string            `json:"h"`
	I      string            `json:"i"`
	X      string            `json:"x"`
	Y      string            `json:"y"`
	Z      string            `json:"z"`
	Tyre   map[string]string `json:"tyre"`
}

type Toy struct {
	ID     string            `json:"id"`
	Number string            `json:"number"`
	Model  string            `json:"model"`
	Wheels string            `json:"wheels"`
	ABC    map[string]string `json:"abc"`
}

type Football struct {
	ID  string         `json:"id"`
	Map map[int]string `json:"map"`
}

type Cricket struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	A    string `json:"a"`
	B    string `json:"b"`
	C    string `json:"c"`
}

type Something struct {
	Basketball Basketball `json:"basketball"`
	Hockey     Hockey     `json:"hockey"`
}

type Tesla struct {
	Elon Anything `json:"elon"`
}

type Anything struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	A    string `json:"a"`
	B    string `json:"b"`
	C    string `json:"c"`
}

type Basketball struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Hockey struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json: "email"`
	Phone     string `json: "phone"`
	CC        string `json: "cc"`
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
	now := time.Now()
	books = append(books, Book{ID: "1", Isbn: "3223", Title: "Book 1", Author: &Author{
		Firstname: "Avneesh", Lastname: "Hota", Email: "avneesh@akto.io", Phone: "+917021916328", CC: "378282246310005"}, Timestamp: now.Unix()})
	books = append(books, Book{ID: "2", Isbn: "2323", Title: "Book 2", Author: &Author{
		Firstname: "Ankush", Lastname: "Jain", Email: "avneesh@akto.io", Phone: "+917021916328", CC: "378282246310005"}, Timestamp: now.Unix()})

	cars = append(cars, Car{ID: "1", Number: "2323", Model: "Car 1",
		A: "1", B: "1", C: "1", D: "1", E: "1", F: "1", G: "1", H: "1", I: "1", X: "1", Y: "2", Z: "3"})

	cars = append(cars, Car{ID: "2", Number: "23423", Model: "Car 2",
		A: "1", B: "1", C: "1", D: "1", E: "1", F: "1", G: "1", H: "1", I: "1", X: "1", Y: "2", Z: "3"})

	cars = append(cars, Car{ID: "3", Number: "9203", Model: "Car 3",
		A: "1", B: "1", C: "1", D: "1", E: "1", F: "1", G: "1", H: "1", I: "1", X: "1", Y: "2", Z: "3"})

	toys = append(toys, Toy{ID: "1", Number: "2323", Model: "Toy 1", Wheels: "32"})
	toys = append(toys, Toy{ID: "2", Number: "1131", Model: "Toy 2", Wheels: "32"})
	toys = append(toys, Toy{ID: "3", Number: "8888", Model: "Toy 3", Wheels: "32"})
	toys = append(toys, Toy{ID: "4", Number: "2312", Model: "Toy 4", Wheels: "32"})

	anything = append(anything, Anything{ID: "4", Name: "2312", A: "Toy 4", B: "32", C: "we"})
	something = append(something, Something{
		Basketball: Basketball{ID: "1", Name: "nameeee"},
		Hockey:     Hockey{A: "1", B: "2", C: "23"},
	})
	tesla = append(tesla, Tesla{Elon: Anything{ID: "4", Name: "2312", A: "Toy 4", B: "32", C: "we"}})

	c := map[int]string{}
	c[1] = "Ankush"
	c[2] = "Avneesh"
	c[3] = "Ankita"
	footballs = append(footballs, Football{ID: "4", Map: c})

	// Route endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("POST")
	r.HandleFunc("/api/books/1", getBooks1).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/cars", getCars).Methods("POST")
	r.HandleFunc("/api/cars/{id}", getCarsId).Methods("GET")
	r.HandleFunc("/api/toys", getToys).Methods("POST")
	r.HandleFunc("/api/toys/{id}", getToysNew).Methods("GET")
	r.HandleFunc("/api/games", getBooks).Methods("POST")
	r.HandleFunc("/api/football", getFootball).Methods("POST")
	r.HandleFunc("/api/cricket", getBooks).Methods("POST")
	r.HandleFunc("/api/auth/signin", signIn).Methods("GET")
	r.HandleFunc("/api/latest/meta-data/local-ipv4", asdf).Methods("GET")
	r.HandleFunc("/api/something/{id}", getNothing).Methods("GET")
	r.HandleFunc("/api/ankush/pepsi/{id}", getPepsi).Methods("GET")
	r.HandleFunc("/api/ankush/coke/{id}", getCoke).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}

// init books var as slice Book struct
var books []Book
var cars []Car
var toys []Toy
var anything []Anything
var something []Something
var tesla []Tesla
var footballs []Football
var counter = 0

func getBooks1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(cars[0])
}

func getNothing(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	counter++
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(tesla[0])
}

func getPepsi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(anything[0])
}

func getCoke(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	counter++
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(something[0])
}

func getCarsId(w http.ResponseWriter, r *http.Request) {
	counter++
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	a := cars[0]
	a.Tyre = map[string]string{"b_" + strconv.Itoa(counter): "c"}
	json.NewEncoder(w).Encode(a)
}

func getToysNew(w http.ResponseWriter, r *http.Request) {
	counter++
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	a := cars[1]
	a.Tyre = map[string]string{
		"b_" + strconv.Itoa(counter): "c",
		"c_" + strconv.Itoa(counter): "d",
		"d_" + strconv.Itoa(counter): "e",
	}
	json.NewEncoder(w).Encode(a)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(books[0])
}

func getCars(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(cars)
}

func getToys(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(toys[0])
}

func getFootball(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(footballs[0])
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
