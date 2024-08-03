# Build stage
FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

# Runtime stage
FROM alpine:3.20

COPY --from=builder /app/apk-crawler /app/apk-crawler

CMD ["/app/apk-crawler"]