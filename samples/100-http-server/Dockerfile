FROM golang:1.23.3-alpine  as builder

WORKDIR /app

COPY . ./

RUN go build -v -o server
# multi stage builds
FROM debian:bookworm-slim

COPY --from=builder /app/server /app/server

EXPOSE 8080

CMD ["/app/server"]
