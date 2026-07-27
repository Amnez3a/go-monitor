FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

FROM alpine:alpine
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/servers.json .
ENTRYPOINT [ "./app" ]
CMD [ "--file", "servers.json" ]

