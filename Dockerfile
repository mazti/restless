#build stage
FROM golang:alpine AS builder
ADD . /go/src/github.com/mazti/restless/
WORKDIR /go/src/github.com/mazti/restless
RUN apk add --no-cache git
RUN go get ./...
RUN go build -o ./restless ./cmd/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com/mazti/restless/restless /restless
COPY --from=builder /go/src/github.com/mazti/restless/config/config.yml /config.yml

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]
LABEL Name=restless
EXPOSE 8080 8081
