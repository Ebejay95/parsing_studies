FROM golang:1.19

WORKDIR /json_validator

COPY . .

RUN go mod tidy
RUN go build -o json_validator app.go

CMD ["/bin/bash"]