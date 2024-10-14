build:
	docker build --rm -t rate-my-media:dev .

run:
	docker run -p 8080:8080 --name rate-my-media rate-my-media:dev