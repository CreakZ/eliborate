package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"yurii-lib/internal/models/domain"
	mock_repository "yurii-lib/internal/repository/mocks"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repository.NewMockBookRepo(ctrl)

	book1 := domain.BookPlacement{
		BookInfo: domain.BookInfo{
			Title:   "Моби Дик, или Белый Кит",
			Authors: []string{"Герман Мелвилл"},
			Description: sql.NullString{
				String: "Книга о выживании в нечеловеческих условиях",
				Valid:  true,
			},
			Category:  4,
			IsForeign: true,
			CoverURL: sql.NullString{
				Valid: false,
			},
		},
		Rack:  1,
		Shelf: 1,
	}

	book2 := domain.BookPlacement{
		BookInfo: domain.BookInfo{
			Title:   "Теория удвоенного времени, или почему я опоздал в Ригу",
			Authors: []string{"Майкл Наки"},
			Description: sql.NullString{
				String: "Пособие по теории удвоенного времени. Аналогичные задачи в пособии не рассматриваются",
				Valid:  true,
			},
			Category:  0,
			IsForeign: false,
			CoverURL: sql.NullString{
				Valid: false,
			},
		},
		Rack:  0,
		Shelf: 14,
	}

	book3 := domain.BookPlacement{
		BookInfo: domain.BookInfo{
			Title: "Зарождение бутсы Вконтакте",
			Authors: []string{
				"Николай Горлангер",
				"Самара Бутсович",
				"Даниил Тикток",
				"Григорий Котлинович"},
			Description: sql.NullString{
				String: "Книга, обязательная для прочтения любому senior PHP-разработчику",
				Valid:  true,
			},
			Category:  11,
			IsForeign: false,
			CoverURL: sql.NullString{
				Valid: false,
			},
		},
		Rack:  9,
		Shelf: 11,
	}

	ctx := context.TODO()

	repo.EXPECT().CreateBook(ctx, book1).Return(1, nil)
	id1, err := repo.CreateBook(ctx, book1)
	require.NoError(t, err)

	repo.EXPECT().CreateBook(ctx, book2).Return(2, nil)
	id2, err := repo.CreateBook(ctx, book2)
	require.NoError(t, err)

	repo.EXPECT().CreateBook(ctx, book3).Return(3, nil)
	id3, err := repo.CreateBook(ctx, book3)
	require.NoError(t, err)

	fmt.Println("created ids:", id1, id2, id3)
}
