FROM golang:1.19

WORKDIR /parsing_studies

COPY . .

RUN go mod tidy
RUN go build -o parsing_studies app.go

CMD ["/bin/bash"]