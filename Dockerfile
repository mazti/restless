#build stage
FROM golang:alpine AS builder
ADD . /go/src/gitlab.com/phypass_server/go/tvsion/
WORKDIR /go/src/gitlab.com/phypass_server/go/tvsion/tvsion-channel
RUN apk add --no-cache git
RUN go get ./...
RUN go build -o ./tvsion-channel ./cmd/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/gitlab.com/phypass_server/go/tvsion/tvsion-channel/tvsion-channel /tvsion-channel
COPY --from=builder /go/src/gitlab.com/phypass_server/go/tvsion/tvsion-channel/config/config.yml /config.yml

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]
LABEL Name=tvsion-channel
EXPOSE 8080 8081
