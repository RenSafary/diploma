FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/admin-service

RUN go build -o admin-service

EXPOSE 50052

CMD ["./admin-service"]