.PHONY = all build db_setup db_start compose_down docker_image

GOBIN=${shell pwd}/bin
export GOBIN

all: db_start db_setup
	docker compose up webserver -d

build:
	CGO_ENABLED=0 go install ./...

db_setup:
	docker exec -i test_db psql -U postgres < create_db.sql
	docker exec -i test_db psql -U postgres -d video_db < schema.sql

db_start:
	docker compose up db -d
	sleep 3

compose_down:
	docker compose down -v

docker_image:
	docker build -t video-api .
