package dto

type BookInfo struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
	CoverUrls   []string `json:"cover_urls"`
}

type BookPlacement struct {
	Rack  int `json:"rack"`
	Shelf int `json:"shelf"`
}

type BookCreate struct {
	BookInfo
	CategoryID int `json:"category_id"`
	BookPlacement
}

type Book struct {
	ID int `json:"id"`
	BookInfo
	Category string `json:"category"`
	BookPlacement
}

type UpdateBookInfo struct {
	Title       *string  `json:"title"`
	Authors     []string `json:"authors"`
	Description *string  `json:"description"`
	CategoryID  *int     `json:"category_id"`
	CoverUrls   []string `json:"cover_urls"`
}

type UpdateBookPlacement struct {
	Rack  *int `json:"rack"`
	Shelf *int `json:"shelf"`
}
