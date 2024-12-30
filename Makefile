run:
	@go run cmd/main/main.go

up:
	@migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/azamon?sslmode=disable' up

down:
	@migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/azamon?sslmode=disable' down