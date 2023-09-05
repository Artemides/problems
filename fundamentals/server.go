package fundamentals

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func Server() {
	http.HandleFunc("/", handler)
	fmt.Println("Listenning on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q \n", r.URL.Host)
}

var count int

var mu sync.Mutex

func ServerCounter() {
	http.HandleFunc("/", countHandler)
	http.HandleFunc("/requests", countRequestsHandler)
	http.HandleFunc("/metadata", metadataHandler)
	fmt.Println("Listenning on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path %q\n", r.URL.Path)

}

func countRequestsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Total requests %v", count)
	mu.Unlock()

}

//handler
//printm method url protocol
//iterate r.Header and print k v

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Address = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
