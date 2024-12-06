package dto

type AdminUserInfo struct {
	Login string `json:"login"`
}

type AdminUserCreate struct {
	AdminUserInfo
	Password string `json:"password"`
}

type AdminUser struct {
	ID int `json:"id"`
	AdminUserCreate
}
