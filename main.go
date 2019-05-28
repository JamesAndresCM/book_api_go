package main

import (
	"encoding/json"
	"os"
	"fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "strconv"

	"github.com/subosito/gotenv"
	"github.com/JamesACM/book_api_go/driver"
)


type Book struct{
  ID int `json:id`
  Title string `json:title`
  Author string `json:author`
  Year int `json:year`
}



func main(){
  gotenv.Load()

  dbName := os.Getenv("DB_NAME")
  dbPass := os.Getenv("DB_PASS")
  dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	
	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

  router := mux.NewRouter()
  router.HandleFunc("/books", getBooks).Methods("GET")
  router.HandleFunc("/books/{id}", getBook).Methods("GET")
  router.HandleFunc("/books", addBook).Methods("POST")
  router.HandleFunc("/books", updateBook).Methods("PUT")
  router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
  log.Println("Get all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Get book")
}

func addBook(w http.ResponseWriter, r *http.Request) {
  log.Println("Add one book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
  log.Println("update book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
  log.Println("remove book")
}
