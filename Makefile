run:
	docker compose exec app go run cmd/main/main.go

codegen:
	docker compose exec app go run github.com/99designs/gqlgen generate

# Docker

container-start:
	docker compose up -d

container-stop:
	docker compose down
