build:
	go build -o ./bin/rate-my-media-mac

run:
	go run .

migrate:
	go run ./database/migrate

migrate-fresh:
	go run ./database/migrate --fresh

docker-build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/rate-my-media-linux
	docker build . -t rate-my-media:latest

docker-run:
	docker run -p 8000:8000 --env-file .env --expose 8000 rate-my-media