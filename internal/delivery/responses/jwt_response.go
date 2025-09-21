package responses

type JwtResponse struct {
	Jwt string `json:"jwt"`
}

func NewJwtResponse(jwt string) JwtResponse {
	return JwtResponse{
		Jwt: jwt,
	}
}
