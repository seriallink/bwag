package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":3000", nil)
}

func helloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World")
}
