FROM golang:1.19-alpine3.16 as builder
WORKDIR /app

# COPY go.mod .
# COPY go.sum .

ADD . .
RUN go mod download
RUN go build -o main .

FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/main .
CMD [ "/app/main" ]