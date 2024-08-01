package convertors_test

import (
	"testing"
	"yurii-lib/internal/convertors"

	"github.com/stretchr/testify/assert"
)

func TestCategoryToInt(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, convertors.CategoryToInt("Библиотека всемирной литературы"))
	a.Equal(50, convertors.CategoryToInt("Медицина"))
	a.Equal(-1, convertors.CategoryToInt("Другое"))
}

func TestCategoryToString(t *testing.T) {
	a := assert.New(t)

	a.Equal("Медицина", convertors.CategoryToString(50))
	a.Equal("Библиотека отечественной общественной мысли", convertors.CategoryToString(28))
	a.Equal("Не определено", convertors.CategoryToString(0))
}
