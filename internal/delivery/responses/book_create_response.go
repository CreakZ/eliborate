package responses

type BookCreateResponse struct {
	ID int `json:"id"`
}

func NewBookCreateResponse(id int) BookCreateResponse {
	return BookCreateResponse{
		ID: id,
	}
}
