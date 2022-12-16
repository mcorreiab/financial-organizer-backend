FROM golang:1.19-alpine as build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /financial-organizer ./cmd/financial-organizer

FROM alpine:3.17.0

WORKDIR /

COPY --from=build /financial-organizer /financial-organizer

EXPOSE 8080

ENTRYPOINT ["/financial-organizer"]
