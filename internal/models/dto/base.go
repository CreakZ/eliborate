package dto

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type PasswordUpdate struct {
	Password string `json:"password"`
}
