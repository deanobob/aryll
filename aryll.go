package main

import (
	"database/sql"
	"log"
	"net/http"
)

var DB *sql.DB

func main() {
	DB, _ = Connect("test.db")

	router := CreateRouter(
		Routes{
			Route{"Index", "GET", "/", Index},
			Route{"Users", "GET", "/users", GetUsers},
			Route{"CreateUser", "PUT", "/users/{username}", CreateUser},
			Route{"UserShow", "GET", "/users/{userID}", GetUserDetails},
			Route{"UserTest", "GET", "/resources/{userID}", Resources},
		})

	log.Fatal(http.ListenAndServe(":8080", router))
}
