package responses

type MessageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(message string) MessageResponse {
	return MessageResponse{
		Message: message,
	}
}

func NewSuccessMessageResponse() MessageResponse {
	return NewMessageResponse("success")
}
