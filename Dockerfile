FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /server ./main.go

FROM gcr.io/distroless/static-debian12

COPY --from=builder /server /server

# Cloud Run はデフォルトで PORT 環境変数を渡してくる
ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/server"]