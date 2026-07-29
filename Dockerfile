FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest
WORKDIR /app
ENTRYPOINT [ "./app" ]
CMD [ "--file", "servers.json" ]
COPY --from=builder /app/app .