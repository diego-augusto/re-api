build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

dc:
	docker build -t re-api .

dc-run: dc
	docker run -d -p ${PORT}:${PORT} -e PORT=${PORT} re-api

dc-stop:
	docker stop $(shell docker ps -q --filter ancestor=re-api)