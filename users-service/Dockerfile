FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/users-service

RUN go build -o users-service

EXPOSE 50053

CMD ["./users-service"]