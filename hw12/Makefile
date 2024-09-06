run:
	go run cmd/main.go

clear-data:
	rm -rf tokens.json users.json

gen_docs:
	swag init -g internal/handler/httphandler/handler.go