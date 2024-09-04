run:
	go run .

migrate:
	go run ./database/migrate

migrate-fresh:
	go run ./database/migrate --fresh
