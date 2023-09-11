package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
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

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func Get() {
	terms := []string{"term"}
	respose, err := searchIssues(terms)
	if err != nil {
		log.Fatalf("error...%s", err)
	}
	fmt.Println(respose)
}

func searchIssues(terms []string) (*IssuesSearchResult, error) {
	const httpUrl = "https://api.github.com/search/issues"
	query := url.QueryEscape(strings.Join(terms, " "))

	response, err := http.Get(httpUrl + "?q=" + query)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("searching query failed %s", response.Status)
	}

	var searchResult IssuesSearchResult

	if err := json.NewDecoder(response.Body).Decode(&searchResult); err != nil {
		return nil, err
	}
	return &searchResult, nil
}
