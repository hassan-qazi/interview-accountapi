FROM golang:1.20-alpine

WORKDIR /app/form3test

COPY ./form3 /app/form3

COPY ./form3test/go.mod ./
COPY ./form3test/go.sum ./

RUN go mod download

COPY ./form3test/integration_test.go ./

RUN go test -c -o form3.test

CMD [ "./form3.test -test.v" ]