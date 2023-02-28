FROM golang:1.20-alpine

WORKDIR /app/form3usage

COPY ./form3 /app/form3

COPY ./hello/go.mod ./

RUN go mod download

COPY ./hello/hello.go ./

RUN go build

CMD [ "./hello" ]