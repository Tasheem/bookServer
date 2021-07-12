package main

import (
	"bookServer/models"
	"bookServer/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getBooks(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request Origin: %s", req.RemoteAddr)
	books, err := services.GetAllBooks()
	if err != nil {
		http.Error(res, "Error Fetching Books.", http.StatusInternalServerError)
	}

	res.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(books)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
	}
}

func postBook(res http.ResponseWriter, req *http.Request) {
	var b models.Book

	if req.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(req.Body).Decode(&b)

		if err != nil {
			http.Error(res, "Invalid JSON", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(res, "Invalid Media Type", http.StatusUnsupportedMediaType)
		return
	}

	go func() {
		fmt.Printf("JSON Object: %v\n", b)
	}()

	err := services.CreateBook(b)
	if err != nil {
		http.Error(res, "Error Creating Book.", http.StatusInternalServerError)
		return
	} else {
		res.Write([]byte("Book Successfully Created."))
	}
}

func handleBooks(res http.ResponseWriter, req *http.Request) {
	method := req.Method
	if method == "POST" {
		postBook(res, req)
	} else if method == "GET" {
		getBooks(res, req)
	}
}

func root(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Book Store API"))
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/books", handleBooks)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err)
	}

	/*var book *models.Book = models.NewBook("1", "1984", "George", "Orwell")
	fmt.Printf("Book: %v", *book)*/
}
