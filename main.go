package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

//   struct Human {
// 	Name string
// 	Email string
//   }

type myHandler struct {
}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

		postgreSQLConnString := os.Getenv("POSTGRESQLCONNSTR_PostGre");
		db, err := sql.Open("postgres", postgreSQLConnString)

		if err != nil {
			panic(err)
		}
		defer db.Close()

		if err = db.Ping(); err != nil {
			panic(err)
		}
		fmt.Println("The database is connected")
	
		rows, err := db.Query("SELECT * FROM \"People\".persons;")
		if err != nil {
			fmt.Println("Error getting persons")
			panic(err)
		}
	
		defer rows.Close()
		var  (
			id    (int)
			name  (string)
			email (string)
		)
		for rows.Next() {
			if err := rows.Scan(&id, &name, &email); err != nil {
				panic(err)
			}
		}
		w.Write([]byte("Name is " + name + " " + "email is " + email))
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("")))
	http.Handle("/person", myHandler{})

	http.ListenAndServe(":8080", nil)
}
