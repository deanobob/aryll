package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]

	Result, err := DB.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, password)
	if err != nil {
		fmt.Fprintln(w, "Failed to insert user:", username)
		fmt.Println(err)
	} else {
		userID, _ := Result.LastInsertId()
		fmt.Fprintln(w, "User created", userID)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
		return
	}
	var username string
	for rows.Next() {
		rows.Scan(&username)
		fmt.Fprintln(w, username)
	}
}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	var username string
	err := DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
	if err != nil {
		fmt.Fprintln(w, "Unknown userID:", userID)
	} else {
		fmt.Fprintln(w, "Username:", username)
	}
}

func Resources(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	rows, err := DB.Query("SELECT name, value FROM resources INNER JOIN resourceDetails ON resourceDetails.resourceID = resources.resourceID WHERE userID = ?", userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	var name string
	var value int
	for rows.Next() {
		err := rows.Scan(&name, &value)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Fprintln(w, name, value)
		}
	}
}
