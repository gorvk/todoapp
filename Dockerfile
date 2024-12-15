FROM golang:1.23.3

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

RUN alias air='$(go env GOPATH)/bin/air'

CMD ["air", "-c", ".air.toml"]  
