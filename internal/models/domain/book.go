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
	CategoryID  int
	CoverUrls   []string
}

type BookSearch struct {
	ID          int      `json:"id" search:"id"`
	Title       string   `json:"title" search:"title"`
	Authors     []string `json:"authors" search:"authors"`
	Description string   `json:"description" search:"description"`
	Category    string   `json:"category" search:"category"`
}
