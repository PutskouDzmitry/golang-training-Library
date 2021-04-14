package main

import (
	"../pkg/constDb"
	"../pkg/data"
	"../pkg/db"
	"fmt"
	"log"
	"os"
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
		host = constDb.Host
	}
	if port == "" {
		port = constDb.Port
	}
	if user == "" {
		user = constDb.User
	}
	if dbname == "" {
		dbname = constDb.DbName
	}
	if password == "" {
		password = constDb.Password
	}
	if sslmode == "" {
		sslmode = constDb.Sslmode
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
		BookId:            12,
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
	changeNumber := 123
	err = bookDate.Update("number", 2, changeNumber)
	if err != nil {
		log.Fatal(err)
	}
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