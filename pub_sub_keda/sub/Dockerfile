#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app /app

ENTRYPOINT /app