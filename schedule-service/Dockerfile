FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/schedule-service

RUN go build -o schedule-service

EXPOSE 50054

CMD ["./schedule-service"]