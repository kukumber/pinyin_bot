# build stage
FROM golang:1.26-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /pinyin_bot cmd/main.go

# runtime stage
FROM alpine:3.21
RUN apk add --no-cache ca-certificates

COPY --from=build /pinyin_bot /usr/local/bin/pinyin_bot

ENTRYPOINT ["pinyin_bot"]
