.PHONY = all db

GOBIN=${shell pwd}/bin
export GOBIN

all:
	CGO_ENABLED=0 go install ./...

db:
	docker exec -i test_db psql -U postgres < create_db.sql
	docker exec -i test_db psql -U postgres -d video_db < schema.sql

docker_image:
	docker build -t video-api .
