package main

//go get -u github.com/go-sql-driver/mysql
//go run RestApi.go
// cd GinGonic
import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql"
)

// func createConnection() *sql.DB {
// 	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "root", "person_db")

// 	//var err error
// 	db, err := sql.Open("postgres", connectionString)
// 	//db, err := sql.Open("postgres", "postgres://postgres:7046365527@localhost/postgres?sslmode=disable")

// 	if err != nil {
// 		panic(err)
// 	}
// 	// check the connection
// 	err = db.Ping()

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Successfully connected!")
// 	return db
// }
func main() {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "root", "product_db")

	//var err error
	db, err := sql.Open("postgres", connectionString)
	//db, err := sql.Open("postgres", "postgres://postgres:7046365527@localhost/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}
	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	//return db

	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/product_db")
	// if err != nil {
	// 	// fmt.Print(err.Error())
	// 	fmt.Println("Error creating DB:", err)
	// 	fmt.Println("To verify, db is:", db)
	// }
	// defer db.Close()
	// fmt.Println("Successfully  Connected to MYSQl")
	// // make sure connection is available
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Print(err.Error())

	type Person struct {
		Id         int    `db:"ID" json:"id"`
		First_Name string `db:"first_name" json:"first_name"`
		Last_Name  string `db:"last_name" json:"last_name"`
		Age        int    `db:"age" json:"age"`
	}

	router := gin.Default()

	
	

	// GET all persons
	router.GET("/persons", func(c *gin.Context) {
		var (
			person  Person
			persons []Person
		)
		rows, err := db.Query("select * from person;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name, &person.Age)
			persons = append(persons, person)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})

		
	})

	// POST new person details
	router.POST("/person", func(c *gin.Context) {
		var buffer bytes.Buffer
		var person Person
		c.Bind(&person)
		// id, err := strconv.Atoi(c.PostForm("id"))
		// fmt.Println("hello", id)
		// //id := c.PostForm("id")
		// first_name := c.PostForm("first_name")
		// last_name := c.PostForm("last_name")
		// Age, err := strconv.Atoi(c.PostForm("Age"))
		id := person.Id

		//id := c.PostForm("id")
		first_name := person.First_Name
		last_name := person.Last_Name
		Age := person.Age
		//Age := c.PostForm("Age")

		//UPDATED CODE
		// sqlStatement := `INSERT INTO person (id,name, location, age) VALUES ($1, $2, $3,$4) RETURNING ID`
		sqlStatement := `INSERT INTO person (id,first_name, last_name, age) VALUES ($1, $2, $3,$4) `

		// the inserted id will store in this id
		//var id int64

		// execute the sql statement
		// Scan function will save the insert id in the id
		// err := db.QueryRow(sqlStatement, user.ID, user.Name, user.Location, user.Age).Scan(&id)
		err := db.QueryRow(sqlStatement, id, first_name, last_name, Age).Scan(&id)

		if err != nil {
			//log.Fatalf("Unable to execute the query. %v", err)
			fmt.Print(err.Error())
		}

		fmt.Printf("Inserted a single record %v", id)

		// return the inserted id
		//return id

		// stmt, err := db.Prepare("insert into person (id,first_name, last_name,Age) values(?,?,?,?);")
		// if err != nil {
		// 	fmt.Print(err.Error())
		// }
		// // _, err = stmt.Exec(&id, &first_name, &last_name, &Age)
		// _, err = stmt.Exec(id, first_name, last_name, Age)

		// if err != nil {
		// 	fmt.Print(err.Error())
		// }

		// Fastest way to append strings
		//buffer.WriteString(id)
		buffer.WriteString(" ")
		buffer.WriteString(first_name)
		buffer.WriteString(" ")
		buffer.WriteString(last_name)
		buffer.WriteString(" ")

		// buffer.WriteString(strconv.Itoa(Age))
		//buffer.WriteString(Age)
		//defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s %ssuccessfully created", first_name, name),
		})
	})

	
	
	router.Run(":9000")
}
