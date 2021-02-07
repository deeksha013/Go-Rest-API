package models

import (
  _ "github.com/mattn/go-sqlite3"
  "database/sql")

var Db *sql.DB

func ConnectDatabase() {

  db,err1:= sql.Open("sqlite3","/home/ds/go/src/project2/test.db")
  if err1 != nil {
   panic("Failed to connect to database test.db!")
  }

  Db = db


}
