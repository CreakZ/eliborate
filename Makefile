# testing book repository
init_book_repo_tests:
	mockgen -source internal/repository/repository.go -destination internal/repository/mocks/mock_repository.go

# testing convertors
convertors_test:
	cd internal/convertors && go test

# testing repository layer
repo_test:
	cd internal/repository/tests && go test

# building Docker environment for API testing
build_test_env:
	cd deploy && sudo docker compose -f compose.test.yaml up --build

# applying migrations on test db
test_migrations_up:
	goose -dir deploy/migrations postgres "postgresql://user:1234@172.17.0.2:5432/lib_test?sslmode=disable" up

# applying migrations on main db TODO
migrations_up:
