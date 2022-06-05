FROM golang:1.18-buster as build

WORKDIR /go/src/app

COPY . .

RUN go mod download && go build -o app

FROM debian:buster-slim

COPY --from=build /go/src/app/app /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]
