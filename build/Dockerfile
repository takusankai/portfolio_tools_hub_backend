FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod ./
COPY go.sum* ./
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./app

# 実行ステージ
FROM alpine:3.18

# 必要なランタイム依存関係のインストール
RUN apk --no-cache add ca-certificates

WORKDIR /app

# ビルドステージからバイナリをコピー
COPY --from=builder /build/main /app/main

# アプリケーションポートの公開
EXPOSE 8080

# アプリケーションの実行
CMD ["/app/main"]