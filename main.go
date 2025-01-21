package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my Go Website</h1>");
}

func main() {

	// Behind the scenes of HandleFunc
	// mux := http.NewServeMux();
	// mux.HandleFunc("/", handlerFunc);
	http.HandleFunc("/", handlerFunc);

	time.Sleep(3 * time.Second);


	fmt.Println("Starting up server on port :3000");
	http.ListenAndServe(":3000", nil);
}

