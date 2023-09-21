package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (usd dollars) String() string { return fmt.Sprintf("$%2.f", usd) }

type database map[string]dollars
type Item struct {
	Item  string  `json:"item"`
	Price dollars `json:"price"`
}

var db = database{"shoes": 32, "i-phone": 1200}

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

func (db database) list(w http.ResponseWriter, r *http.Request) {
	var items []Item
	for item, price := range db {
		items = append(items, Item{item, price})
	}

	json, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "error encoding items as json", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		msg := fmt.Sprintf("Unsuported Method: %s", r.Method)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}

	params := r.URL.Query()
	item := params.Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %s\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}
func (db database) AddItem(item string, price dollars) error {
	if item == "" {
		return fmt.Errorf("empty item or price not allowed")

	}
	_, ok := db[item]
	if ok {
		return fmt.Errorf("item already created")
	}

	err := validatePrice(price)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	db[item] = dollars(price)
	return nil
}

func validatePrice(price dollars) error {

	if price < 0 {
		return fmt.Errorf("item cannot cost less than 0")
	}

	return nil
}
func (db database) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := fmt.Sprintf("Unsuported Method: %s", r.Method)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	var item Item

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("item", item.Item, item.Price)

	err := db.AddItem(item.Item, item.Price)
	if err != nil {
		msg := fmt.Sprintf("item: %s , price: %s", item.Item, item.Price)
		http.Error(w, msg, http.StatusConflict)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		msg := fmt.Sprintf("Parsing Item to JSON failed %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (db database) readHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	params := r.URL.Query()
	item := params.Get("item")
	price, exists := db.exists(item)
	if !exists {
		msg := fmt.Sprintf("item: %s. not found", item)
		http.Error(w, msg, http.StatusOK)
		return
	}
	var response Item = Item{item, price}
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error encoding Item as json", http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func (db database) exists(item string) (dollars, bool) {
	price, ok := db[item]
	if !ok {
		return 0, ok
	}
	return price, ok
}

func RunHTTP() {
	log.Fatal(http.ListenAndServe(":3000", db))
}
func RunHandlerMux() {
	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.readHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
