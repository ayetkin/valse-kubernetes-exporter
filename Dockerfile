FROM golang:1.17-alpine3.15 as builder
WORKDIR /opt
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .

FROM alpine:3.15
RUN apk add tzdata
ENV TZ="Europe/Istanbul"
WORKDIR /app
COPY --from=builder /opt/main /app/main
CMD ["/app/main"]