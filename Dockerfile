FROM golang:latest

WORKDIR /app
COPY . .

RUN go get -d -v ./... \
    && go install -v ./... \
    && go build -v -o main

CMD ["/app/main"]
