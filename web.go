package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%v\n", r)
	// fmt.Printf("%+v\n", r)
	fmt.Printf("%#v\n", r)
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main() {
	//routing
	http.HandleFunc("/", handler)

	//server起動
	http.ListenAndServe(":8080", nil)
}
