# Golang take-home test

Sample RESTful API for a video resource with many annotation subresources.

Implemented using Golang, Gin, GORM and Postgres.

Because this is a sample, only three tests are included which require a real database. This can improved upon by utilizing test database (that are reset on each run) or by using mocks (mockery, a Golang mock generator).

## How to run

Ensure docker and docker-compose (or compose plugin) are installed, then run:

```
make
```

This will start Postgres on port 5432, import the database schema and finally start the webserver on port 8080.

To shutdown containers (and delete database volume) use:

```
make compose_down
```

## Tests

Three sample test are provided, to run them do:

```
make test
```

Note that the database must be up (and migrations have run), if you ran `make` previously then that should have already been done for you.

## Resetting the database

This will destroy the local database, create a new one and import the schema.

```
make compose_down
make db_start
make db_setup
```

## Sample cURL commands

### Create a video

```
curl -H "Authorization: hunter2" \
     -X POST http://localhost:8080/v1/videos \
     -d '{"name": "Baby Shark", "description": "doo doo doo doo", "url": "https://bucket.storage.com/abc123/video.mp4"}'
```

### List videos

```
curl -H "Authorization: hunter2" http://localhost:8080/v1/videos
```

### Get a video

```
curl -H "Authorization: hunter2" http://localhost:8080/v1/videos/<VIDEO_UUID>
```

### Delete a video

```
curl -H "Authorization: hunter2" \
     -X DELETE \
     http://localhost:8080/v1/videos/<VIDEO_UUID>
```

### Create an annotation

```
curl -H "Authorization: hunter2" \
     -X POST http://localhost:8080/v1/videos/<VIDEO_UUID>/annotations \
     -d '{"note": "this contains an advertisement", "type": "different_language", "start": "00:59:04", "end": "01:00:00"}'
```

### Update an annotation

```
curl -H "Authorization: hunter2" \
     -X PUT http://localhost:8080/v1/videos/<VIDEO_UUID>/annotations/<ANNOTATION_UUID> \
     -d '{"note": "this contains an advertisement", "type": "advertisement", "start": "00:27:15", "end": "00:35:17"}'
```

### Get annotations for a video

```
curl -H "Authorization: hunter2" http://localhost:8080/v1/videos/<VIDEO_UUID>/annotations
```

### Get a specific annotation

```
curl -H "Authorization: hunter2" http://localhost:8080/v1/videos/<VIDEO_UUID>/annotations/<ANNOTATION_UUID>
```

### Delete an annotation

```
curl -H "Authorization: hunter2" \
     -X DELETE \
     http://localhost:8080/v1/videos/<VIDEO_UUID>/annotations/<ANNOTATION_UUID>
```
