FROM golang:1.24.1
WORKDIR /app
COPY go.mod go.sum
COPY . .
RUN go mod download
RUN go build -o main ./cmd/app
EXPOSE 8081
CMD ["./main"]
