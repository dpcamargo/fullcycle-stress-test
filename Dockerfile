FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stress-test ./cmd/main.go

ENTRYPOINT [ "./stress-test" ]