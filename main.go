package main

import (
	"fmt"
	"time"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path != "" {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	} else {
		fmt.Fprint(w, "Hello World!\n")
		fmt.Fprint(w, "The time is", time.Now())
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Fprint(w, sum)
	}
}
