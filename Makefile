DB_URL=postgresql://root:secret@localhost:5432/investify?sslmode=disable
#migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force VERSION

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpines

migrateup:
	go run cmd/migrations/init/init.sql.go up




migratedown:
	go run cmd/migrations/init/init.sql.go down


sqlc:
	sqlc generate


server:
	go run cmd/app/main.go
