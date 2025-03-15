# 共通ベースイメージ
FROM golang:1.24.0-bookworm AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# 開発環境用イメージ
FROM base AS dev
# 開発ツールのインストール
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    postgresql-client \
    git \
    && rm -rf /var/lib/apt/lists/*

# Airとその他開発ツールのインストール
RUN go install github.com/air-verse/air@latest

# ソースコードをコピー
COPY . .

# エントリーポイントを設定
COPY ./build/scripts/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# 初期化SQLとAir設定をコピー
COPY ./db/init/init.sql /app/init.sql
COPY ./build/.air.toml /app/.air.toml

# 必要なディレクトリを作成
RUN mkdir -p /app/tmp

# エントリーポイント
ENTRYPOINT ["/app/docker_entrypoint.sh"]
CMD ["/app/main"]

# ビルダーイメージ
FROM base AS builder
# ソースコードをコピーしてビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./app

# 本番環境用イメージ
FROM debian:bookworm-slim AS prod
# 必要最小限のツールのみインストール
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# ワーキングディレクトリ
WORKDIR /app

# ビルドしたバイナリと必要ファイルのみをコピー
COPY --from=builder /app/main /app/main
COPY --from=builder /app/db/init/init.sql /app/init.sql
COPY --from=builder /app/build/scripts/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# エントリーポイント
ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["/app/main"]
