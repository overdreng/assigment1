package library

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
}

func (l *Library) BorrowBook(id string) {
	if book, exists := l.Books[id]; exists && !book.IsBorrowed {
		book.IsBorrowed = true
		l.Books[id] = book
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.Books[id]; exists && book.IsBorrowed {
		book.IsBorrowed = false
		l.Books[id] = book
	}
}

func (l *Library) ListBooks() []Book {
	availableBooks := []Book{}
	for _, book := range l.Books {
		if !book.IsBorrowed {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}
