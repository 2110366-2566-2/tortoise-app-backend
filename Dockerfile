FROM golang:1.22.2-alpine

LABEL MAINTAINER="PetPal"

RUN mkdir /petpal
COPY go.mod go.sum /petpal/
WORKDIR /petpal

RUN go mod download
COPY . ./

RUN go mod tidy

# Build in "cmd/app/main.go"
RUN go build -o main cmd/app/main.go

CMD ["/main"]

