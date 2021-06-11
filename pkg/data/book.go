package data

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
	//dbConst "github.com/PutskouDzmitry/golang-training-Library/pkg/const_db"
)

//Entity in database
type Book struct {
	BookId            primitive.ObjectID `bson:"_id"`
	AuthorId          int                `bson:"author_id"`
	PublisherId       int                `bson:"publisher_id"`
	NameOfBook        string             `bson:"name_of_book"`
	YearOfPublication string             `bson:"year_of_publication"`
	BookVolume        int                `bson:"book_volume"`
	Number            int                `bson:"number"`
}

//ReadAll output all data with table books
func (B BookData) ReadAll() ([]Book, error) {
	var books []Book
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := B.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book Book
		err = cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, mongo.ErrNoDocuments
	}
	return books, nil
}

//Read read data in db
func (B BookData) Read(name string) (Book, error) {
	var book Book
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := B.collection.FindOne(ctx, &Book{NameOfBook: name}).Decode(&book)
	if err != nil {
		return book, err
	}
	return book, nil
}

//Add add data in db
func (B BookData) Add(book Book) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := B.collection.InsertOne(ctx, book)
	logrus.Info(result)
	if result == nil {
		return fmt.Errorf("error")
	}
	return nil
}

////Update update number of books by the id
//func (B BookData) Update(id int, value int) error {
//	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
//	result, _ := B.collection.UpdateOne(ctx,bson.M{"number":value}, bson.D{
//		{"$set",bson.D{{"number":value}}},
//	})
//	result := B.db.Table(dbConst.Books).Where(dbConst.BookId, id).Update("number", value)
//	if result.Error != nil {
//		return fmt.Errorf(dbConst.CantUpdateDataError, result.Error)
//	}
//	return nil
//}

//String output data in console
func (B Book) String() string {
	return fmt.Sprintln(B.BookId, B.AuthorId, B.PublisherId, strings.TrimSpace(B.NameOfBook), B.YearOfPublication, B.BookVolume, B.Number)
}

//BookData create a new connection
type BookData struct {
	collection *mongo.Collection
}

//NewBookData it's imitation constructor
func NewBookData(collection *mongo.Collection) *BookData {
	return &BookData{collection: collection}
}

//Don't ready(

//
//
////Update update number of books by the id
//func (B BookData) Update(id int, value int) error {
//	result := B.db.Table(dbConst.Books).Where(dbConst.BookId, id).Update("number", value)
//	if result.Error != nil {
//		return fmt.Errorf(dbConst.CantUpdateDataError, result.Error)
//	}
//	return nil
//}
//
////Delete delete data in db
//func (B BookData) Delete(id int) error {
//	result := B.db.Where(dbConst.BookId, id).Delete(&Book{})
//	if result.Error != nil {
//		return fmt.Errorf(dbConst.CantDeleteDataError, result.Error)
//	}
//	return nil
//}
