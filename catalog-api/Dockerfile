# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY catalog-api/go.mod .
RUN go mod download

COPY catalog-api/*.go ./

RUN go build -o /catalog-api

EXPOSE 8010

CMD [ "/catalog-api" ]