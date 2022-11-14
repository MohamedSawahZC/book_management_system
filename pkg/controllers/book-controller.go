package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MohamedSawahZC/book_management_system/pkg/models"
	"github.com/MohamedSawahZC/book_management_system/pkg/utils"
	"github.com/gorilla/mux"
)


var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request){
	newBooks := models.GetAllBooks()
	w,_ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(w)
}

func GetBookById(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("err while parsing book")
	}
	bookDetails , _ := models.GetBookById(ID)
	w,_ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(w)

}

func CreateBook(res http.ResponseWriter, req *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(req,CreateBook)
	b := CreateBook.CreateBook()
	w , _ := json.Marshal(b)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(w)
}

func DeleteBook(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("err while parsing book")
	}
	book  := models.DeleteBook(ID)
	w,_ := json.Marshal(book)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(w)

}

func UpdateBook(res http.ResponseWriter, req *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(req,updateBook)
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("err while parsing book")
	}
	bookDetails , db := models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	w,_ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(w)
}