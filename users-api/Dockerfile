# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY users-api/go.mod .
RUN go mod download

COPY users-api/*.go ./

RUN go build -o /users-api

EXPOSE 8020

CMD [ "/users-api" ]