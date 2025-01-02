FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go build main.go

FROM alpine:3.19

COPY --from=builder /app .

EXPOSE 5000

CMD ["./main"]
