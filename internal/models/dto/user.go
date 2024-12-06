package dto

type UserInfo struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type UserCreate struct {
	UserInfo
	Password string `json:"password"`
}

type User struct {
	ID int `json:"id"`
	UserCreate
}
