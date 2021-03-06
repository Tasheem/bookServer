package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tasheem/bookServer/models"
	"github.com/Tasheem/bookServer/services"
)

func getBooks(res http.ResponseWriter, req *http.Request) {
	books, err := services.GetAllBooks()
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Error Fetching Books.", http.StatusInternalServerError)
	}

	res.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(books)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Error Fetching Books", http.StatusInternalServerError)
	}
}

func postBook(res http.ResponseWriter, req *http.Request) {
	var b models.Book

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("JSON Object: %v\n", b)

	id, err := services.CreateBook(b)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Error Creating Book.", http.StatusInternalServerError)
		return
	}

	resource := fmt.Sprintf("/api/books/id=%s", id)
	res.Header().Set("Location", resource)
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("Book Successfully Created"))
}

// Client should send book object with updated price and existing id.
func updatePrice(res http.ResponseWriter, req *http.Request) {
	var b models.Book

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Book: %v\n", b)
	err = services.UpdatePrice(b)
	if err != nil {
		if err.Error() == "row not found" {
			http.Error(res, "Resource not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Error Updating Price.", http.StatusInternalServerError)
		return
	}

	res.Write([]byte("Book Successfully Updated."))
}

func deleteBook(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	fmt.Printf("ID: %s\n", id)
	if len(id) == 0 {
		http.Error(res, "Error with query string.", http.StatusBadRequest)
		return
	}

	err := services.DeleteBook(id)
	if err != nil {
		http.Error(res, "Error Deleting Book.", http.StatusInternalServerError)
		return
	}

	res.Write([]byte("Book Successfully Deleted."))
}

func handleBooks(res http.ResponseWriter, req *http.Request) {
	origin := req.Header.Get("Origin")
	fmt.Printf("Origin: %v\n", origin)

	// Prevent any client from access except for authServer.
	if origin != "localhost:4000" {
		http.Error(res, "Unauthorized Origin", http.StatusForbidden)
		return
	}

	method := req.Method
	if method == "POST" {
		if req.Header.Get("Content-Type") != "application/json" {
			http.Error(res, "Invalid Media Type", http.StatusUnsupportedMediaType)
			return
		}
		postBook(res, req)
	} else if method == "GET" {
		getBooks(res, req)
	} else if method == "PUT" {
		if req.Header.Get("Content-Type") != "application/json" {
			http.Error(res, "Invalid Media Type", http.StatusUnsupportedMediaType)
			return
		}
		updatePrice(res, req)
	} else if method == "DELETE" {
		deleteBook(res, req)
	} else {
		http.Error(res, "Invalid Media Type", http.StatusMethodNotAllowed)
		return
	}
}

func root(res http.ResponseWriter, req *http.Request) {
	origin := req.Header.Get("Origin")
	fmt.Printf("Origin: %v\n", origin)

	// Prevent any client from access except for authServer.
	if origin != "localhost:4000" {
		http.Error(res, "Unauthorized Origin", http.StatusForbidden)
		return
	}

	res.Write([]byte("Book Store API"))
}

func main() {
	http.HandleFunc("/api", root)
	http.HandleFunc("/api/books", handleBooks)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err)
	}

	/*var book *models.Book = models.NewBook("1", "1984", "George", "Orwell")
	fmt.Printf("Book: %v", *book)*/
}
