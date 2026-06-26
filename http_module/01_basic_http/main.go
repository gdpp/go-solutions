package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.Write([]byte("Hello from Go"))
}

func main(){
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Try going to http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}

