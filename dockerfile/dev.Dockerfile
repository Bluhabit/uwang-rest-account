FROM golang:alpine as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o uwang-rest-account .

FROM alpine:edge

WORKDIR /app

COPY --from=builder /app/uwang-rest-account .
RUN apk --no-cache add ca-certificates tzdata
EXPOSE 7009
ENTRYPOINT ["/app/uwang-rest-storage"]