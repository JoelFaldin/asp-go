package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hey there!")
}

func main() {
	// Bindear funciones a rutas específicas:
	http.HandleFunc("/", test)

	http.ListenAndServe(":8080", nil)
}
