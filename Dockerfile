# Builder Stage
FROM golang:1.18.3-alpine3.16 AS builder

RUN apk update
RUN apk add --no-cache curl bash nano

ARG VERSION
ARG CGO_ENABLED

WORKDIR /app/src
COPY . .

RUN go mod tidy
RUN go build -ldflags "-s -w -X main.version=${VERSION}" -o /app/qiscus-caa ./cmd

WORKDIR /app

RUN rm -rf src/
ENTRYPOINT ["/app/qiscus-caa"]


# Production Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/qiscus-caa ./app
CMD ["./app api"]
