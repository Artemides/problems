package fundamentals

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", handler)
	fmt.Println("Listenning on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q \n", r.URL.Host)
}
