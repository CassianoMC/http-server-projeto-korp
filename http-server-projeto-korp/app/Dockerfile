FROM golang:1.25.0

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go build -o http-server-projeto-korp . 

EXPOSE 8080

CMD ["./http-server-projeto-korp"]