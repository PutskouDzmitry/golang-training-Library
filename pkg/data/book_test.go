package data

import (
	"database/sql"
	"errors"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/constDb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func NewGorm(db *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DriverName:           "postgres",
		Conn:                 db,
	})
	gormDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gormDb
}

var testBook = Book{
	BookId:            1,
	AuthorId:          2,
	PublisherId:       3,
	NameOfBook:        "LordoftheRings",
	YearOfPublication: "2017-12-5",
	BookVolume:        50,
	Number:            10,
}

var testResult = Result{
	BookId:          16,
	NameOfBook:      "TestBook",
	NameOfPublisher: "TestBook2",
}

func TestBookData_ReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	rows := sqlmock.NewRows([]string{"book_id", "author_id", "publisher_id", "name_of_book", "year_of_publication", "book_volume", "number"}).
		AddRow(testBook.BookId, testBook.AuthorId, testBook.PublisherId, testBook.NameOfBook, testBook.YearOfPublication, testBook.BookVolume, testBook.Number)
	mock.ExpectQuery(constDb.SelectAllBooks).WillReturnRows(rows)
	products, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(products)
	assert.Equal(products[0], testBook)
	assert.Len(products, 1)
}


func TestBookData_Read(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	rows := sqlmock.NewRows([]string{"book_id", "name_of_book", "name_of_publisher"}).
		AddRow(testResult.BookId, testResult.NameOfBook, testResult.NameOfPublisher)
	mock.ExpectQuery(constDb.SelectFromBooksWithID).WithArgs(1).WillReturnRows(rows)
	users, err := data.Read(1)
	assert.NoError(err)
	assert.NotEmpty(users)
	assert.Equal(users[0], testResult)
	assert.Len(users, 1)
}

func TestBookData_Add(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(constDb.InsertBook).WithArgs(testBook.BookId, testBook.AuthorId, testBook.PublisherId, testBook.NameOfBook, testBook.YearOfPublication, testBook.BookVolume, testBook.Number).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	id, err := data.Add(testBook)
	assert.NoError(err)
	assert.Equal(id, testBook.BookId)
}

func TestBookData_Update(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(constDb.Update).
		WithArgs(testBook.NameOfBook, testBook.BookId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := data.Update("name_of_book", testBook.BookId, testBook.NameOfBook)
	assert.NoError(err)
}

func TestBookData_Delete(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(constDb.Delete).
		WithArgs(testBook.BookId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := data.Delete(testBook.BookId)
	assert.NoError(err)
}

func TestBookData_UpdateErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(constDb.Delete).
		WithArgs(testBook.BookId).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	err := data.Delete(testBook.BookId)
	assert.Error(err)
}

func TestBookData_ReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	mock.ExpectQuery(constDb.SelectAllBooks).WillReturnError(errors.New("something went wrong..."))
	products, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(products)
}

func TestBookData_ReadErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	mock.ExpectQuery(constDb.ReadBookWithJoin).WillReturnError(errors.New("something went wrong..."))
	products, err := data.Read(0)
	assert.Error(err)
	assert.Empty(products)
}

func TestBookData_AddErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(constDb.InsertBook).WithArgs(testBook.BookId, testBook.AuthorId, testBook.PublisherId, testBook.NameOfBook, testBook.YearOfPublication, testBook.BookVolume, testBook.Number).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	id, err := data.Add(testBook)
	assert.Error(err)
	assert.Equal(id, -1)
}

func TestBookData_DeleteErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewBookData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(constDb.Delete).
		WithArgs(testBook.BookId).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	err := data.Delete(testBook.BookId)
	assert.Error(err)
}