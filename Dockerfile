FROM golang:1.18.2

COPY . /
WORKDIR /

RUN go mod download
RUN go build -o backend main.go

CMD ["./backend"]
EXPOSE 8080


