package controllers

import ("github.com/gin-gonic/gin"
       "project2/models"
       "net/http"
        "fmt")


     func GetBook(c *gin.Context){
       id := c.Param("id")
       //statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
       //statement.Exec()
       //statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
       //statement.Exec("Nic", "Raboy")
      sqlStatement1 := "SELECT Name,Author,Cost FROM books WHERE Id="+id
      fmt.Println(sqlStatement1)
       rows, err2 := models.Db.Query(sqlStatement1)
       if err2 != nil {
         fmt.Println(err2)
        panic("Failed to query database test.db!")
      }
      fmt.Println(rows)
       var cost int
       var author string
       var name string
       for rows.Next() {
             rows.Scan(&name,&author,&cost)
             fmt.Println(name,author,cost)
         }

        c.JSON(http.StatusOK,gin.H{"name":name,"author":author,"cost":cost})


     }


     func GetAllBooks(c *gin.Context){
       var books = []models.Book{}
       var book = models.Book{}
       sqlStatement := "SELECT Id,Name,Author,Cost from books"
        rows,err:=models.Db.Query(sqlStatement)
        if err!=nil{
         panic("db query to fetch all books failed")
       }
       defer rows.Close()
       for rows.Next() {
         rows.Scan(&book.Id,&book.Name,&book.Author,&book.Cost)
         books= append(books,book)
         fmt.Println(book)
       }
       c.JSON(http.StatusOK,gin.H{"books":books})
     }



     func CreateBook(c *gin.Context){
      fmt.Println("inside creation...........")

      //extracting data from the request
      type requestBody struct{   //this struct is defined to validate user's input to prevent ourselves from invalid data
        Name string `json:"name" binding:"required"`
        Author string `json:"author" binding:"required"`
        Cost int `json:"cost" binding:"required"`
      }
      var input = []requestBody{}
      if err := c.ShouldBindJSON(&input) ; err != nil {
        panic("unable to bind json")
      }
      fmt.Println(input)


     //using the extracted data to form sql statement
     sqlStatement := "INSERT INTO books (Name,Author,Cost) VALUES "
     var sqlVals = []interface{}{}
     for _,val:= range input{
       sqlStatement = sqlStatement +  "(?, ?, ?),"
       sqlVals = append(sqlVals,val.Name,val.Author,val.Cost)
     }
    fmt.Println(sqlStatement,sqlVals)
    sqlStatement = sqlStatement[0:len(sqlStatement)-1] //trim the last ','


    //executing sql statemets
    stmt,err := models.Db.Prepare(sqlStatement)
    if err!=nil{
      fmt.Println(err)
      panic("failed to prepare stmt")
    }
    response,err := stmt.Exec(sqlVals...)
    if err!=nil{
      fmt.Println(err)
      panic("failed to execute stmt")
    }
    fmt.Println(response)

    c.JSON(http.StatusOK,gin.H{"response":response})

     }



     func UpdateBookCost(c *gin.Context){

      //extracting data from request requestBody
      type requestBody struct{
        Id int `json:"id"  required:"binding"`
        Cost int `json:"cost" required:"binding"`
      }

      var input = []requestBody{}
      if err := c.ShouldBindJSON(&input); err!=nil{
        panic("unable to bind json")
        }
      fmt.Println(input)

      //create sql sqlStatement
      sqlStatement := "UPDATE books SET Cost = CASE Id"
      sqlVals := []interface{}{}
      for _,val := range input{
        sqlStatement = sqlStatement + " WHEN ? THEN ?"
        sqlVals = append(sqlVals,val.Id,val.Cost)
      }
      sqlStatement = sqlStatement + " END WHERE Id IN ("
      for index,val := range input{
        sqlVals = append(sqlVals,val.Id)
        if index == len(input)-1{
         sqlStatement = sqlStatement + "?)"
       }else{
         sqlStatement = sqlStatement + "?,"}
      }
      fmt.Println(sqlStatement)
      fmt.Println(sqlVals)

      //execute sqlStatement
      stmt,err := models.Db.Prepare(sqlStatement)
      if err != nil{
        fmt.Println(err)
        panic("unable to prepare sql stmt")
      }

      response,err := stmt.Exec(sqlVals...)
      if err != nil{
        panic("unable to prepare sql stmt")
      }

      c.JSON(http.StatusOK,gin.H{"response":response})
     }
