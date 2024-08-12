FROM golang:1.22-alpine

RUN mkdir /app

COPY go.mod /app
COPY go.sum /app

WORKDIR /app

RUN go mod download

COPY . .

RUN go build -o myapp ./main.go

EXPOSE 8080

CMD ["./myapp"]
