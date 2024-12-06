# генерация моков для репозиторного слоя книг
init_book_repo_tests:
	mockgen -source internal/repository/repository.go -destination internal/repository/mocks/mock_repository.go

# тестирование функций-конверторов
convertors_test:
	cd internal/convertors && go test

# генерация документации
docs:
	swag init --dir ./cmd,./internal/delivery/handlers,./internal/delivery/responses --parseDependency
