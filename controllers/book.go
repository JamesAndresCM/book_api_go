package controllers

import (
  "encoding/json"
  "log"
  "net/http"
  "fmt"
  "github.com/gorilla/mux"
  "strconv"
  "github.com/JamesAndresCM/book_api_go/models"
  "github.com/JamesAndresCM/book_api_go/driver"
)

type Controller struct{}

var books []models.Book

func logFatal(err error){
  if err != nil {
    log.Fatal(err)
  }
}

func (c Controller) GetBooks(db *driver.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    books = []models.Book{}
    rows, err := db.SQL.Query("SELECT *  FROM books")
    logFatal(err)

    defer rows.Close()
    for rows.Next(){
      err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
      logFatal(err)
      books = append(books, book)
    }
    fmt.Println(books)
    json.NewEncoder(w).Encode(books)
  }
}

func (c Controller) GetBook(db *driver.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    data_not_found := make(map[string]string)
    params := mux.Vars(r)
    bookid, err := strconv.Atoi(params["id"])
    logFatal(err)
    rows := db.SQL.QueryRow("SELECT * FROM books where id=$1", bookid)
    error := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
    if error == nil {
      json.NewEncoder(w).Encode(book)
    } else {
      data_not_found["error"] = "book not found"
      json.NewEncoder(w).Encode(data_not_found)
    }
  }
}


func (c Controller) DestroyBook(db *driver.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    params := mux.Vars(r)
    bookid, err := strconv.Atoi(params["id"])
    logFatal(err)
    rows := db.SQL.QueryRow("SELECT * FROM books where id=$1", bookid)
    error := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
    if error == nil {
      db.SQL.QueryRow("DELETE FROM books where id=$1", bookid)
      json.NewEncoder(w).Encode("book destroyed")
    } else {
      json.NewEncoder(w).Encode("book not found")
    }
  }
}

func (c Controller) AddBook(db *driver.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request){
    var book models.Book
    var bookID int
    book_message := make(map[string]string)
    json.NewDecoder(r.Body).Decode(&book)
    error := db.SQL.QueryRow("INSERT INTO books (title, author, year) values ($1,$2,$3) RETURNING id;",
      book.Title, book.Author, book.Year).Scan(&bookID)
      logFatal(error)
    id := strconv.Itoa(bookID)
    book_message["success"] = "book" + id + "has been added"
    json.NewEncoder(w).Encode(book_message)
  }
}
