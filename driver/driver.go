package driver

import (
  "database/sql"
  "fmt"
  
	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectPSQL(host, port, user, password, dbname string) (*DB, error){
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  dbConn.SQL = db
	return dbConn, err
}

