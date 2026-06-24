FROM golang:1.26-alpine AS builder
WORKDIR /app 
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-monitor .

FROM alpine:3.20
WORKDIR /root/
COPY --from=builder /go-monitor .
COPY servers.json .
CMD [ "./go-monitor" ]
