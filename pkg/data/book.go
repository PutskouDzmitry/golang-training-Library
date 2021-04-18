package data

import (
	dbConst "example.com/m/pkg/constDb"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

//Entity in database
type Book struct {
	BookId int // primary key
	AuthorId int // foreign key
	PublisherId int // foreign key
	NameOfBook string // name of book
	YearOfPublication string // year of publication of the book
	BookVolume int // book volume
	Number int // number of book
}

//String output data in console
func (B Book) String() string {
	return fmt.Sprintln(B.BookId, B.AuthorId, B.PublisherId, strings.TrimSpace(B.NameOfBook), B.YearOfPublication, B.BookVolume, B.Number)
}

//BookData create a new connection
type BookData struct {
	db *gorm.DB // connection in db
}

//NewBookData it's imitation constructor
func NewBookData(db *gorm.DB) *BookData {
	return &BookData{db: db}
}

func (B BookData) ReadAll() ([]Book, error) {
	var books []Book
	result := B.db.Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read users from database, error: %w", result.Error)
	}
	return books, nil
}

//Read read data in db
func (B BookData) Read() ([]Result, error) {
	var results []Result
	result := B.db.Table(dbConst.Publishers).Select(dbConst.SelectBookAndPublisher).
		Joins(dbConst.ReadBookWithJoin).
		Find(&results)
	if result.Error != nil {
		return nil, result.Error
	}
	return results, nil
}

//Add add data in db
func (B BookData) Add(book Book) (int, error)  {
	result := B.db.Create(&book)
	if result.Error != nil 	{
		return -1, fmt.Errorf(dbConst.CantAddDataError, result.Error)
	}
	return book.BookId, nil
}

//Update update data in db
func (B BookData) Update(colume string, id int, value interface{}) error {
	result := B.db.Table(dbConst.Books).Where(dbConst.BookId, id).Update(colume, value)
	if result.Error != nil {
		return fmt.Errorf(dbConst.CantUpdateDataError, result.Error)
	}
	return nil
}

//Delete delete data in db
func (B BookData) Delete(id int) error {
	result := B.db.Where(dbConst.BookId, id).Delete(&Book{})
	if result.Error != nil {
		return fmt.Errorf(dbConst.CantDeleteDataError, result.Error)
	}
	return nil
}