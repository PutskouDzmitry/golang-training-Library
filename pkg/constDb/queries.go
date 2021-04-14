package constDb

//Const for queries
const (
	ReadBookWithJoin       = "RIGHT JOIN books on books.publisher_id=publishers.publisher_id"
	SelectBookAndPublisher = "books.book_id, books.name_of_book ,publishers.name_of_publisher"
	BookId = "book_id = ?"
)
