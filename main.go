package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/login", login).Methods("POST")
	log.Fatal(http.ListenAndServe(":6900", handlers.CORS(
		handlers.AllowCredentials())(router)))
	fmt.Println("Server started at port 6500")
}

type LoginRequest struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	var loginRequestObj LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequestObj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	response := struct {
		Status       string `json:"status"`
		Account_type string `json:"account_type"`
		Token        string `json:"token"`
		Expires      string `json:"expires"`
	}{
		"success",
		"desk",
		"c1b5ebd45281f0480cb9c6de963e5dd420c577d8b5abdc0e45944fddc988c67d",
		"2017-12-15 23:59:59"}

	sendResponse, _ := json.Marshal(response)
	w.Write(sendResponse)
}
