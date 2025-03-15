package controllers

import (
	"encoding/json"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Failed to marshal books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)

	book := newBook.CreateBook()
	if book == nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteBook(ID)
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = updateBook.UpdateBook(ID)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	updatedBook, _ := models.GetBookById(ID)

	res, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}