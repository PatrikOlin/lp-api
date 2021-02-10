FROM golang:1.15 as build

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o lp-api

FROM debian:buster

WORKDIR /app
COPY --from=build /app/lp-api /app


EXPOSE 8125

CMD ["./lp-api"]
