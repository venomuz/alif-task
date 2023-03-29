FROM golang:1.18.2-alpine

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/app/main.go

EXPOSE 8080

CMD ./main