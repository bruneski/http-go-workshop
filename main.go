package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=uft-8");
	fmt.Fprint(w, "<h1>Welcome to my Go Website</h1>");
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><br/><p>To get in touch, please email <a href=\"\">admin@matthewrbrune.com</a></p>");
	fmt.Println(r.URL.Path);
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w,r)
// 	case "/contact":
// 		contactHandler(w,r)
// 	default:
// 		//TODO: Handle 404 page
// 		http.Error(w, "Page Does Not Exist", http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound);
// 		// fmt.Fprint(w, "Page Not Found");
// 	}

// 	// if(r.URL.Path == "/") {
// 	// 	homeHandler(w, r)
// 	// } else if (r.URL.Path == "/contact") {
// 	// 	contactHandler(w,r)
// 	// }
// 	fmt.Println(r.URL.Path);
// 	fmt.Println(r.URL.RawPath);

// }


// type Server struct {
// 	DB *sql.DB
// }

// func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {

// }
type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w,r)
	case "/contact":
		contactHandler(w,r)
	default:
		//TODO: Handle 404 page
		http.Error(w, "Page Does Not Exist", http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound);
		// fmt.Fprint(w, "Page Not Found");
	}
}

func main() {

	// Behind the scenes of HandleFunc
	// mux := http.NewServeMux();
	// mux.HandleFunc("/", homeHandler);
	//http.HandleFunc("/", pathHandler);
	//http.HandleFunc("/contact", contactHandler);

	//time.Sleep(3 * time.Second);

	// http.Handler - interface with the ServeHTTP method
	// http.HandlerFunc - a function that accepts the same Args as ServeHTTP method and also implements http.Handlerhttp.Handler

	// http.Handle("/", http.Handler)
	// http.HandlerFunc("/", pathHandler)

	// var router http.HandlerFunc;
	// router = pathHandler
	fmt.Println("Starting up server on port :3000");
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler));
}

