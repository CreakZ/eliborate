package domain

import (
	"database/sql"

	"github.com/lib/pq"
)

type BookInfo struct {
	Title       string
	Authors     pq.StringArray
	Description sql.NullString
	Category    string
	CoverUrls   pq.StringArray
}

type BookPlacement struct {
	Rack  int
	Shelf int
}

type BookCreate struct {
	BookInfo
	BookPlacement
}

type Book struct {
	ID int
	BookInfo
	BookPlacement
}

type UpdateBookInfo struct {
	Title       sql.NullString
	Authors     pq.StringArray
	Description sql.NullString
	Category    string
	CoverUrls   pq.StringArray
}

type BookSearch struct {
	ID          int      `json:"id" search:"id"`
	Title       string   `json:"title" search:"title"`
	Authors     []string `json:"authors" search:"authors"`
	Description string   `json:"description" search:"description"`
	Category    string   `json:"category" search:"category"`
}
