FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o cmd/main cmd/main.go

EXPOSE 8080

ENTRYPOINT [ "./cmd/main" ]