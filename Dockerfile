# Stage 1
FROM golang:alpine as builder
RUN apk update && apk add --no-cache
RUN mkdir /build 
ADD . /build/
WORKDIR /build

RUN go mod download
RUN go build -o app ./cmd/app
# Stage 2
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./app"]