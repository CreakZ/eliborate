package domain

type BookInfo struct {
	Title       string
	Authors     []string
	Description string
	CoverUrls   []string
}

type BookPlacement struct {
	Rack, Shelf int
}

type BookCreate struct {
	BookInfo
	CategoryID int
	BookPlacement
}

type Book struct {
	ID int
	BookInfo
	Category string
	BookPlacement
}

type UpdateBookInfo struct {
	Title       *string
	Authors     []string
	Description *string
	CategoryID  *int
	CoverUrls   []string
}

type UpdateBookPlacement struct {
	Rack, Shelf *int
}
