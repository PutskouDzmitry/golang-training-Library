package main

import (
	"github.com/PutskouDzmitry/golang-training-Library/api"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/constDb"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/data"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/db"
	"gorilla/mux"
	"log"
	"net"
	"net/http"
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
	// 2. create router that allows to set routes
	r := mux.NewRouter()
	// 3. connect to data layer
	userData := data.NewBookData(conn)
	// 4. send data layer to api layer
	api.ServeUserResource(r, *userData)
	// 5. cors for making requests from any domain
	r.Use(mux.CORSMethodMiddleware(r))
	// 6. start server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server Listen port...", err)
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}