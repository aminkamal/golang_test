FROM golang:1.19-bullseye as builder
COPY . /var/build
WORKDIR /var/build
RUN make


FROM alpine:latest
RUN mkdir -p /var/app/
COPY --from=builder /var/build/bin/video-api /var/app/
WORKDIR /var/app
CMD ["./video-api"]
