package store

import (
    "errors"
)

var (
    ErrNotFound = errors.New("not found")
    ErrExist    = errors.New("exist")
)

// Book
type Book struct {
    Id      string   `json:"id"`      // 图书 ISBN ID
    Name    string   `json:"name"`    // 图书名称
    Authors []string `json:"authors"` // 图书作者
    Press   string   `json:"press"`   // 出版社
}

// Store
type Store interface {
    Create(*Book) error
    Update(*Book) error
    Get(string) (Book, error)
    GetAll() ([]Book, error)
    Delete(string) error
}
