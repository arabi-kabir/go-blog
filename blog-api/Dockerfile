# -------- Stage 1: Build --------
FROM golang:1.26.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# -------- Stage 2: Run --------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]