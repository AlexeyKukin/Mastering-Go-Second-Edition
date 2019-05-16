package main

import (
	"fmt"
	"net/http"
)

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is an example server.\n")
	fmt.Fprintf(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/", Default)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}