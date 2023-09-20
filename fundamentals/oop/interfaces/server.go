package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

func RunSever() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", TableHndler)
	fmt.Println("Listenning on Port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Error Listenning 3000")
	}
}
