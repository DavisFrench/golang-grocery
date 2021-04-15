FROM golang:latest AS builder

COPY . /go/src/golang-grocery

WORKDIR /go/src/golang-grocery

RUN CGO_ENABLED=0 GOOS=linux go build -o grocery cmd/app/main.go

FROM golang:alpine

COPY --from=builder /go/src/golang-grocery/grocery .

EXPOSE 8888

CMD ./grocery
