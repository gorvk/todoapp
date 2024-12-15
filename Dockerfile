FROM golang:1.23.3

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod tidy

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]  
