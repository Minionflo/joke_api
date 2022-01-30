FROM golang:alpine
RUN mkdir -p /app
WORKDIR /app
COPY . /app/
RUN go build main.go
CMD ["./main"]
