FROM golang:1.17-alpine

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY /kv-store-project/go.mod /kv-store-project/go.sum ./
RUN go mod download

COPY /kv-store-project ./

RUN go build -o app .

EXPOSE 8000

CMD ["./app"]