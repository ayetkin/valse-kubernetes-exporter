FROM golang:1.17.1-buster as builder
WORKDIR /opt
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .

FROM debian:buster-slim
ENV TZ="Europe/Istanbul"
WORKDIR /app
COPY --from=builder /opt/main /app/main
CMD ["/app/main"]