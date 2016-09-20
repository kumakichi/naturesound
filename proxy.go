package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Listen local at :80")
	http.ListenAndServe(":80", nil)
}
