package store

import (
    "sync"

    storeStruct "github.com/book-store/store"
    "github.com/book-store/store/factory"
)

func init() {
    factory.Register("mem", &MemStore{
        books: make(map[string]*storeStruct.Book),
    })
}

type MemStore struct {
    sync.RWMutex
    books map[string]*storeStruct.Book
}

// Create creates a new Book in the store.
func (ms *MemStore) Create(book *storeStruct.Book) error {
    ms.Lock()
    defer ms.Unlock()

    if _, ok := ms.books[book.Id]; ok {
        return storeStruct.ErrExist
    }

    nBook := *book
    ms.books[book.Id] = &nBook

    return nil
}

// Update updates the existed Book in the store.
func (ms *MemStore) Update(book *storeStruct.Book) error {
    ms.Lock()
    defer ms.Unlock()

    oldBook, ok := ms.books[book.Id]
    if !ok {
        return storeStruct.ErrNotFound
    }

    nBook := *oldBook
    if book.Name != "" {
        nBook.Name = book.Name
    }

    if book.Authors != nil {
        nBook.Authors = book.Authors
    }

    if book.Press != "" {
        nBook.Press = book.Press
    }

    ms.books[book.Id] = &nBook

    return nil
}

// Get retrieves a book from the store, by id. If no such id exists. an
// error is returned.
func (ms *MemStore) Get(id string) (storeStruct.Book, error) {
    ms.RLock()
    defer ms.RUnlock()

    t, ok := ms.books[id]
    if ok {
        return *t, nil
    }
    return storeStruct.Book{}, storeStruct.ErrNotFound
}

// Delete deletes the book with the given id. If no such id exist. an error
// is returned.
func (ms *MemStore) Delete(id string) error {
    ms.Lock()
    defer ms.Unlock()

    if _, ok := ms.books[id]; !ok {
        return storeStruct.ErrNotFound
    }

    delete(ms.books, id)
    return nil
}

// GetAll returns all the books in the store, in arbitrary order.
func (ms *MemStore) GetAll() ([]storeStruct.Book, error) {
    ms.RLock()
    defer ms.RUnlock()

    allBooks := make([]storeStruct.Book, 0, len(ms.books))
    for _, book := range ms.books {
        allBooks = append(allBooks, *book)
    }
    return allBooks, nil
}
