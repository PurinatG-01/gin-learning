FROM golang:1.20-alpine as builder
WORKDIR /app
RUN apk update && apk add --no-cache git
ENV GIN_MODE=release
COPY . .
RUN go mod download
RUN go build -o main .
FROM golang:1.20-alpine
WORKDIR /app
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]