package utils

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(left, right []byte) bool {
	return bcrypt.CompareHashAndPassword(left, right) == nil
}
