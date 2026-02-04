FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/auth-service

RUN go build -o auth-service

EXPOSE 50051

CMD ["./auth-service"]