package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PutskouDzmitry/golang-training-Library/pkg/const_db"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/data"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/db"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = const_db.Host
	}
	if port == "" {
		port = const_db.Port
	}
	if user == "" {
		user = const_db.User
	}
	if dbname == "" {
		dbname = const_db.DbName
	}
	if password == "" {
		password = const_db.Password
	}
	if sslmode == "" {
		sslmode = const_db.Sslmode
	}
}

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	bookDate := data.NewBookData(conn)
	books, err := bookDate.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)
	newBook := data.Book{
		BookId:            122,
		AuthorId:          4,
		PublisherId:       2,
		NameOfBook:        "Lord of the Rings",
		YearOfPublication: "2017-12-5",
		BookVolume:        50,
		Number:            10,
	}
	err = bookDate.Add(newBook)
	if err != nil {
		log.Fatal(err)
	}
	books, err = bookDate.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)
	books, err = bookDate.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)
	err = bookDate.Delete(11)
	books, err = bookDate.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)
}
