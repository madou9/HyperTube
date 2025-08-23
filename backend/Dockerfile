FROM golang:1.22-alpine

RUN apk add --no-cache git curl

# Install Air for live reload
RUN go install github.com/cosmtrek/air@v1.43.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

# Use Air instead of go run
CMD ["air", "-c", ".air.toml"]
