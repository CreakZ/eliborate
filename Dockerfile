FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/main cmd/main.go


FROM scratch
COPY --from=builder /app/cmd/main /app/cmd/main
COPY --from=builder /app/docs/ /app/docs/

WORKDIR /app
CMD [ "./cmd/main" ]
