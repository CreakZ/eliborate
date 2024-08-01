package domain

type AdminUserInfo struct {
	Login string
}

type AdminUserCreate struct {
	AdminUserInfo
	Password string
}

type AdminUser struct {
	ID int
	AdminUserInfo
}
