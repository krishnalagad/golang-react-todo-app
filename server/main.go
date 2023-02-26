package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/krishnalagad/golang-react-todo/router"
)

// main function
func main() {
	fmt.Println("Starting project on port 9000")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":9000", r))
}
