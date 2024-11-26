FROM golang:1.23 AS builder

RUN go install github.com/air-verse/air@latest

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -v -o server ./cmd/api

FROM gcr.io/distroless/static-debian11 AS production

COPY --from=builder /build/server .

CMD ["./server"]
