package dto

type BookInfo struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description *string  `json:"description"`
	Category    string   `json:"category"`
	CoverUrls   []string `json:"cover_urls"`
}

type BookPlacement struct {
	Rack  int `json:"rack"`
	Shelf int `json:"shelf"`
}

type BookCreate struct {
	BookInfo
	BookPlacement
}

type Book struct {
	ID int `json:"id"`
	BookInfo
	BookPlacement
}

type UpdateBookInfo struct {
	Title       *string  `json:"title"`
	Authors     []string `json:"authors"`
	Description *string  `json:"description"`
	Category    *string  `json:"category"`
	CoverUrls   []string `json:"cover_urls"`
}

type BookSearch struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
}
