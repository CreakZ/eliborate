package domain

type UserInfo struct {
	Login string
	Name  string
}

type UserCreate struct {
	UserInfo
	Password string
}

type User struct {
	ID int
	UserCreate
}
