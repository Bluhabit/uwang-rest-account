FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o uwang-rest-account .

FROM alpine:edge

WORKDIR /app

COPY --from=builder /app/uwang-rest-account .
RUN apk --no-cache add ca-certificates tzdata
EXPOSE 7007
ENTRYPOINT ["/app/uwang-rest-storage"]