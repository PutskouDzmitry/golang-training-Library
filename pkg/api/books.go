package api

import (
	"encoding/json"
	"log"
	"net/http"
	//"strconv"

	"github.com/PutskouDzmitry/golang-training-Library/pkg/data"

	"github.com/gorilla/mux"
)

type bookAPI struct {
	data *data.BookData
}

func ServeUserResource(r *mux.Router, data data.BookData) {
	api := &bookAPI{data: &data}
	r.HandleFunc("/books", api.getAllBooks).Methods("GET")
	r.HandleFunc("/book{name}", api.getOneBook).Methods("GET")
	r.HandleFunc("/books", api.createBook).Methods("POST")
	//r.HandleFunc("/books{id}/{number}", api.updateBook).Methods("PUT")
	//r.HandleFunc("/books{id}", api.deleteBook).Methods("DELETE")
}

func (a bookAPI) getAllBooks(writer http.ResponseWriter, request *http.Request) {
	users, err := a.data.ReadAll()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users "))
		if err != nil {
			log.Println(err)
		}
	}
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a bookAPI) getOneBook(writer http.ResponseWriter, request *http.Request) {
	idRequest := mux.Vars(request)
	name := idRequest["name"]
	user, err := a.data.Read(name)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if user.NameOfBook != "" {
		err = json.NewEncoder(writer).Encode(user)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a bookAPI) createBook(writer http.ResponseWriter, request *http.Request) {
	book := new(data.Book)
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if book == nil {
		log.Printf("failed empty JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.Add(*book)
	if err != nil {
		log.Println("user hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

//func (a bookAPI) updateBook(writer http.ResponseWriter, request *http.Request) {
//	idRequest := mux.Vars(request)
//	id, err := strconv.Atoi(idRequest["id"])
//	if err != nil {
//		log.Println(err)
//		writer.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	strNumber := idRequest["number"]
//	number, err := strconv.Atoi(strNumber)
//	if err != nil {
//		log.Println("book hasn't been updated, because number doesn't equal int:", number)
//		writer.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	err = a.data.Update(id, number)
//	if err != nil {
//		log.Println("book hasn't been updated")
//		writer.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	writer.WriteHeader(http.StatusCreated)
//}
//
//func (a bookAPI) deleteBook(writer http.ResponseWriter, request *http.Request) {
//	idRequest := mux.Vars(request)
//	id, err := strconv.Atoi(idRequest["id"])
//	if err != nil {
//		log.Println(err)
//		writer.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	err = a.data.Delete(id)
//	if err != nil {
//		log.Println("book hasn't been deleted")
//		writer.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	writer.WriteHeader(http.StatusCreated)
//}
