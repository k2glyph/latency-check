FROM golang:1.12.16-alpine3.10 as builder
LABEL maintainer="Dinesh Katwal <medineshkatwal@gmail.com>"
RUN apk update && apk --no-cache add ca-certificates && apk add git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o latency .

FROM alpine:3.10  

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/latency .
COPY --from=builder /app/run.sql .
COPY --from=builder /app/drain.sql .

