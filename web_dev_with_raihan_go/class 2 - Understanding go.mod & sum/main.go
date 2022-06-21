package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	fmt.Println("Serve running at port: 8080")
	http.ListenAndServe(":8080", r)

}
