FROM golang:latest

WORKDIR /app
COPY . .

RUN go get -u github.com/gin-gonic/gin
RUN go build -o main .

CMD ["./main"]

