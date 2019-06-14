package main

import (
  "os"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "strconv"
  "github.com/JamesAndresCM/book_api_go/controllers"
	"github.com/subosito/gotenv"
	"github.com/JamesAndresCM/book_api_go/driver"
)



func logFatal(err error){
  if err != nil {
    log.Fatal(err)
  }
}

func main(){

  gotenv.Load()

  dbName := os.Getenv("DB_NAME")
  dbPass := os.Getenv("DB_PASS")
  dbHost := os.Getenv("DB_HOST")
  dbPort := os.Getenv("DB_PORT")

  portdb, _ := strconv.Atoi(dbPort)
	db, err := driver.ConnectPSQL(dbHost, "snake", dbPass, dbName, portdb)
	logFatal(err)
  router := mux.NewRouter()
  controller := controllers.Controller{}
  router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
  router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
  router.HandleFunc("/books/{id}", controller.DestroyBook(db)).Methods("DELETE")
  /*router.HandleFunc("/books", addBook).Methods("POST")
  router.HandleFunc("/books", updateBook).Methods("PUT")
  router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
*/
//  fmt.Println(http.ListenAndServe(":8080", nil))
  logFatal(http.ListenAndServe(":3000", router))
}
