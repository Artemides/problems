package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func Run() {
	movies := []Movie{
		{Title: "The Lord of the rings", Year: 2001, Color: true, Actors: []string{"Candalf", "Frodo", "Sam", "Pipin"}},
		{Title: "The Lord of the rings I", Year: 2004, Color: true, Actors: []string{"Candalf", "Frodo", "Sam", "Pipin"}},
		{Title: "The Lord of the rings II", Year: 2005, Color: true, Actors: []string{"Candalf", "Frodo", "Sam", "Pipin"}},
		{Title: "The Hobbit I", Year: 2014, Color: true, Actors: []string{"Candalf", "Legolas", "Sam", "Pipin"}},
		{Title: "The Hobbit II", Year: 2016, Color: true, Actors: []string{"Candalf", "Frodo", "Sam", "Pipin"}},
		{Title: "The Hobbit III", Year: 2019, Color: true, Actors: []string{"Candalf", "Frodo", "Sam", "Pipin"}},
	}
	data, err := json.Marshal(movies)
	dataIndented, err2 := json.MarshalIndent(movies, "", "   ")
	if err != nil || err2 != nil {
		log.Fatalf("Error Json Marshaling... %s", err)
	}

	var decoded []struct {
		Title  string
		Actors []string
	}
	if err := json.Unmarshal(data, &decoded); err != nil {
		log.Fatalf("Error Decoding... %s", err)
	}
	fmt.Printf("%s\n", decoded)
	fmt.Printf("%s\n%s\n", dataIndented, data)
	// plot.WriteToFile("movies.json", data, true)
}
