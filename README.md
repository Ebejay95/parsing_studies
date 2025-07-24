docker build --tag "parsing_studies" .

docker run -it -v $(pwd):/parsing_studies
export DEBUG=0
export DEBUG=1

go build -o "parsing_studies" app.go