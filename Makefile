.PHONY = all db

GOBIN=${shell pwd}/bin
export GOBIN

all:
	go install ./...

db:
	docker exec -i test_db psql -U postgres < create_db.sql
	docker exec -i test_db psql -U postgres -d video_db < schema.sql
