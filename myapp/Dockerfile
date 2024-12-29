# Use official Go image as a base
FROM golang:1.23.4

WORKDIR /app

# Install Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@v1.52.0

# debugging port
EXPOSE 2345

CMD ["dlv", "debug", "--headless", "--listen=0.0.0.0:2345", "--api-version=2", "--log", "--accept-multiclient", "--continue", "./main.go"]
