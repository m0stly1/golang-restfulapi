FROM golang:latest

LABEL maintainer="Anders"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build

CMD ["./playground1"]