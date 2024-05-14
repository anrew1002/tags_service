
FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/

FROM alpine
EXPOSE 1090

WORKDIR /build

ADD .env /build

COPY --from=builder /build/server /build/server

CMD ["./server"]



