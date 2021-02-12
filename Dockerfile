FROM golang:alpine AS build

RUN apk add --no-cache curl git alpine-sdk

WORKDIR $GOPATH/src/github.com/bradmccoydev/teamwork-cli

COPY go.mod go.sum $GOPATH/src/github.com/bradmccoydev/teamwork-cli/

RUN go mod tidy

COPY . .

RUN go build -o /teamwork-cli

FROM alpine:latest

WORKDIR /teamwork-cli

COPY --from=build /teamwork-cli teamwork-cli

ENTRYPOINT [ "./teamwork-cli" ]
