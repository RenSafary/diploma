FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/client

RUN go build -o api-gateway

EXPOSE 8080

CMD ["./api-gateway"]