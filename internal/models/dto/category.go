package dto

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryCreateUpdate struct {
	Name string `json:"name"`
}
