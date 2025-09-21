package responses

import "eliborate/internal/models/dto"

type Categories struct {
	Categories []dto.Category `json:"categories"`
}
