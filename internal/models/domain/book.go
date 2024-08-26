package domain

import (
	"database/sql"

	"github.com/lib/pq"
)

type BookInfo struct {
	Title       string
	Authors     pq.StringArray
	Description sql.NullString
	Category    int
	IsForeign   bool
	CoverURL    sql.NullString
}

type UpdateBookInfo struct {
	Title                 sql.NullString
	Authors               pq.StringArray
	Description, Category sql.NullString
	IsForeign             sql.NullBool
	CoverURL              sql.NullString
}

type BookPlacement struct {
	BookInfo
	Rack  int
	Shelf int
}

type Book struct {
	ID int
	BookPlacement
}
