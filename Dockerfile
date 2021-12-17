FROM golang:alpine AS build
ENV GOPROXY=https://proxy.golang.org
WORKDIR /app


COPY go.mod go.sum /app/
RUN go mod download
COPY  . /app/
RUN go build -o webapi

FROM alpine 

WORKDIR /app
COPY --from=build /app/webapi /app/
COPY ./infrastructure/datastore/migrations /app/test-beer/infrastructure/datastore/migrations

ENTRYPOINT  ["/app/webapi"]
