FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o backend ./cmd/main.go

CMD ["./backend"]
