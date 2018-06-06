package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
	err := errors.New("whoops")
	fmt.Println(err)

	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", router))

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}

func health(s string) string {
	return s
}
