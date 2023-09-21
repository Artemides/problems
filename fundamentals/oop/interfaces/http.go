package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (usd dollars) String() string { return fmt.Sprintf("$%2.f", usd) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s : %s\n", item, price)
		}
	case "/price":
		params := r.URL.Query()
		item := params.Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item %s\n", item)
			return
		}

		fmt.Fprintf(w, "%s\n", price)

	default:
		msg := fmt.Sprintf("No such page %s", r.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}

func RunHTTP() {
	db := database{"shoes": 32, "i-phone": 1200}
	log.Fatal(http.ListenAndServe(":3000", db))
}
