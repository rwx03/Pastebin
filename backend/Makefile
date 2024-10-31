run:
	go run cmd/app/main.go -config="configs/config.yml"

migrate:
	migrate -path ./schema -database "postgres://postgres:123123@localhost:5432/pbin?sslmode=disable" up