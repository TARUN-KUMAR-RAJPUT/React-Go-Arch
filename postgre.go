package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)


type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
	Mobile string `json:"mobile"`
}

func main() {
	// Set up the database connection
	r := gin.Default()

	pgConnString := "" 

	db, err := sql.Open("postgres", pgConnString)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r.Use(cors.Default())

    // Define a GET route to retrieve all users from the "users" table
    r.GET("/users", func(c *gin.Context) {
		
		// Perform a query
		rows, err := db.Query("select * from test_table_uchiha")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

        fmt.Println(rows)
        // Create a slice of User structs to store the retrieved users
        var users []User

        // Loop through the result set and create a User struct for each row
        for rows.Next() {
            var user User
            err := rows.Scan(&user.ID, &user.Name, &user.Mobile)
            if err != nil {
                fmt.Println(err)
            }
            users = append(users, user)
        }

        // Check for any errors that may have occurred while iterating through the result set
        err = rows.Err()
        if err != nil {
			fmt.Println(err)
        }

        // Return the slice of User structs as JSON
        c.JSON(200, users)
    })

	r.POST("/addusers", func(c *gin.Context) {
        // Bind the request body to a User struct
        var user User
        err := c.BindJSON(&user)
        if err != nil {
		fmt.Println(err)
        }

        // Execute an INSERT query to insert the new user into the "users" table
        _, err = db.Exec("INSERT INTO test_table_uchiha (id, name, mobile) VALUES ($1, $2, $3)", user.ID, user.Name, user.Mobile)
        if err != nil {
		fmt.Println(err)
        }

        // Return a success message
        c.JSON(200, gin.H{
            "message": "User created",
        })
    })


	// Start the Gin server
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
