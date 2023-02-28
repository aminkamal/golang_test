# Golang take-home test

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

## Sample cURL commands

### Create a video

```
curl -H "Authorization: hunter2" \
     -X POST http://localhost:8080/v1/videos \
     -d "{\"name\": \"Baby Shark!\", \"description\": \"doo doo doo doo\", \"url\": \"https://bucket.storage.com/abc123/video.mp4\"}"
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
     -d "{\"note\": \"this contains an advertisement\", \"type\": \"different_language\", \"start\": \"00:59:04\", \"end\": \"01:00:00\"}"
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
