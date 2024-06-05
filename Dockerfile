FROM golang:1.22.3

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /rest-auth

EXPOSE 8080

CMD ["/rest-auth"]

