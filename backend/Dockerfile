FROM golang:1.16-alpine

WORKDIR /

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /main

EXPOSE 8080

CMD [ "/main" ]