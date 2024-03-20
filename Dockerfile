FROM golang:1.21-alpine as build
RUN apk update && apk add git gcc g++

COPY . /app
COPY .env /app/.env

WORKDIR /app

RUN go build -tags musl main.go

FROM alpine 
WORKDIR /app

COPY --from=build /app/main ./
COPY --from=build /app/.env ./
CMD ["./main", "serve"]
