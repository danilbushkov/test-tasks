FROM golang:1.22.5

WORKDIR /usr/src/app


COPY . .


RUN go build -v -o /usr/local/bin/app ./cmd/app/main.go

CMD ["app"]
