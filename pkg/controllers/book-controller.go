package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var NewBook models.Book 

func GetBooks(w http.ResponseWriter,r *http.Request) {
	newBooks := models.GetAllBooks()
	res,_ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetail,_ := models.GetBookById(ID)
	res,_ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r,Book)
	b := Book.CreateBook()
	res,_ := json.Marshal(b)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r,updatedBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetail,db := models.GetBookById(ID)
	if updatedBook.Name != "" {
		bookDetail.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetail.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetail.Publication = updatedBook.Publication
	}
	db.Save(&bookDetail)
	res,_ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}