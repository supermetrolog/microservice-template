FROM golang:1.18-alpine

WORKDIR /app

COPY ./ ./

RUN go build -o build/web -v cmd/main.go

EXPOSE 9000

CMD ["build/web"]