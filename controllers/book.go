package controllers

import (
	"encoding/json"
  "log"
  "net/http"
  //"strconv"
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

func (c Controller) getBooks(db *driver.DB) http.HandlerFunc {
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
			json.NewEncoder(w).Encode(books)
		}
}

  
  
  
  /*
  func getBook(db *driver.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	  var book models.Book
	  params := mux.Vars(r)
	  bookid, err := strconv.Atoi(params["id"])
	  logFatal(err)
	  rows := db.SQL.QueryRow("SELECT * FROM books where id=$1", bookid)
	  error := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	  logFatal(error)
	  json.NewEncoder(w).Encode(book)
	}
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
	*/
  