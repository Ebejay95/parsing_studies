FROM golang:1.19 AS build-stage

WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

WORKDIR /
