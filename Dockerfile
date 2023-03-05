FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .
RUN swag init

EXPOSE 8181

CMD ["./main"]