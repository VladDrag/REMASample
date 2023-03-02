package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

//   struct Human {
// 	Name string
// 	Email string
//   }

func main() {

	postgreSQLConnString := "postgres://catAdmin:passKT99,@postktsrvr.postgres.database.azure.com/postgres?sslmode=require"
	db, err := sql.Open("postgres", postgreSQLConnString)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")

	rows, err := db.Query("SELECT * FROM \"People\".persons;")
	if err != nil {
		fmt.Println("Error getting persons")
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id    (int)
			name  (string)
			email (string)
		)	
		if err := rows.Scan(&id, &name, &email); err != nil {
			panic(err)
		}
		fmt.Println("The name is " + name)
	}

	http.Handle("/", http.FileServer(http.Dir("")))
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}