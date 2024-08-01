package dto

type BookInfo struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description *string  `json:"description"`
	Category    string   `json:"category"`
	IsForeign   bool     `json:"is_foreign"`
	Logo        *string  `json:"logo"`
}

type UpdateBookInfo struct {
	Title       *string   `json:"title"`
	Authors     *[]string `json:"authors"`
	Description *string   `json:"description"`
	Category    *string   `json:"category"`
	IsForeign   *bool     `json:"is_foreign"`
	Logo        *string   `json:"logo"`
}

type BookPlacement struct {
	BookInfo
	Rack  int `json:"rack"`
	Shelf int `json:"shelf"`
}

type Book struct {
	ID int `json:"id"`
	BookPlacement
}
