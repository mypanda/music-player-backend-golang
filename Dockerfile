#get a base image
FROM golang:1.16-buster

MAINTAINER anaiya raisinghani <anaiya.raisinghani@mongodb.com>

WORKDIR /go/src/app
COPY ./src .

RUN go get -d -v
RUN go build -v

CMD ["./docker-golang-example"]




FROM golang as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build/api
COPY go.mod ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o api
# post build stage
FROM alpine
WORKDIR /root
COPY --from=builder /build/api/api .
EXPOSE 8080
CMD ["./api"]