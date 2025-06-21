package servutils

func CountOffset(page, limit int) int {
	return (page - 1) * limit
}
