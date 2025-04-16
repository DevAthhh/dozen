.SILENT:

run-file:
	go mod tidy && go run cmd/app/main.go

docker-compose:
	docker-compose up --remove-orghans