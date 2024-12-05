package collect_test

import (
	"eliborate/internal/models/domain"
	"eliborate/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectTagValuesFromStructByTag(t *testing.T) {
	a := assert.New(t)

	strct := struct {
		Fimoz   string `some:"long"`
		Hamster int    `some:"kombat"`
	}{}
	tag := "some"

	someMore := domain.BookSearch{}

	if !a.Equal(
		[]string{"long", "kombat"},
		utils.CollectTagValuesFromStructByTag(tag, strct),
	) {
		t.Error("test 1 failed")
	}

	if !a.Equal(
		[]string{"id", "title", "authors", "description", "category"},
		utils.CollectTagValuesFromStructByTag("search", someMore),
	) {
		t.Error("test 2 failed")
	}
}
