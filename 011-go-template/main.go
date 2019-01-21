// https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
// https://www.w3schools.com/html/html5_serversentevents.asp

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type (
	BooksPage struct {
		BooksListing []Book
	}
	Book struct {
		Isbn    string
		Author  string
		Title   string
		PubDate string
		Price   Currency
	}
	Currency float64
)

func main() {

	http.HandleFunc("/books/", booksHandler)
	// register static files handle '/index.html -> client/index.html'
	http.Handle("/", http.FileServer(http.Dir("client")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	booksPage := &BooksPage{
		loadBooks("books.json"),
	}
	tmpl := template.Must(template.ParseFiles("server/books/index.html"))

	err := tmpl.Execute(w, booksPage)
	if err != nil {
		log.Fatal(err)
	}
}
func loadBooks(fn string) []Book {
	var books []Book
	f, err := os.Open(fn)
	if nil != err {
		log.Fatal(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.Decode(&books)
	return books
}
func (c Currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}
