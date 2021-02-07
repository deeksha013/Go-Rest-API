package models
// import "database/sql"

type Book struct{
  Id int `json:id`
  Name string `json:name`
  Author string `json:author`
  Cost int `json:cost`
}
