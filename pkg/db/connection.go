package db

import (
	"fmt"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/constDb"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//GetConnection it's return a new connection in db
func GetConnection(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(constDb.AddInfoForConnection,
		host, port, user, dbname, password, sslmode)
	connection, err :=  gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil, fmt.Errorf(constDb.TroubleWithConnection, err)
	}
	return connection, nil
}
