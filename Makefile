.PHONY: sqlc serve

sqlc:
	sqlc generate

serve:
	go run main.go
