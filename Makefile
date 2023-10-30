.PHONY: dev-start dev-down run

dev-start:
	docker-compose up -d --build

dev-stop:
	docker-compose down

run:
	go run main.go
