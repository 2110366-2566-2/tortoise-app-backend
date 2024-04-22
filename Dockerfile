FROM golang:1.22.2-alpine

LABEL MAINTAINER="PetPal"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./cmd/app/main ./cmd/app/main.go

CMD ["./cmd/app/main"]

