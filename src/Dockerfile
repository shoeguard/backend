FROM golang:1.16.5-alpine3.13 as builder

WORKDIR /usr/src/app

RUN apk update; apk add git ca-certificates

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o bin/main cmd/server/main.go

FROM alpine

COPY --from=builder /usr/src/app/bin/main ./main

ENV PORT=80
EXPOSE 80

ENTRYPOINT ["./main"]