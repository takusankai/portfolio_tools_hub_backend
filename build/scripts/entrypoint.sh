#!/bin/bash
set -e
echo "Starting application..."
echo "Current PATH: ${PATH}"

# IPv6接続状況確認
echo "Checking IPv6 connectivity..."
ip -6 route || echo "IPv6 routing not available"

# マイグレーション実行の制御
if [ "${RUN_MIGRATIONS}" = "true" ]; then
    echo "Running database migrations..."

    # 環境変数に基づいてDBホスト決定
    if [ "${USE_LOCAL_DB}" = "true" ]; then
        # ローカルDBへ接続
        echo "Using local database at: ${DB_HOST}"
        echo "DB_PORT: ${DB_PORT}"
        echo "POSTGRES_USER: ${POSTGRES_USER}"
        echo "POSTGRES_DB: ${POSTGRES_DB}"

        # マイグレーション実行（初期化SQLを使用）
        PGPASSWORD="${POSTGRES_PASSWORD}" psql \
            -h "${DB_HOST}" \
            -U "${POSTGRES_USER}" \
            -d "${POSTGRES_DB}" \
            -c "SELECT 1" \
            -f /app/init.sql
    else
        # Supabase DBへ接続
        echo "Using Supabase database at: ${DB_HOST}"
        echo "DB_PORT: ${DB_PORT}"
        echo "POSTGRES_USER: ${POSTGRES_USER}"
        echo "POSTGRES_DB: ${POSTGRES_DB}"

        # IPv6を使用して接続
        if PGPASSWORD="${POSTGRES_PASSWORD}" psql \
            -h "${DB_HOST}" \
            -U "${POSTGRES_USER}" \
            -d "${POSTGRES_DB}" \
            --set=sslmode=require \
            -c "SELECT 1"; \
            then
            echo "Connection successful via URI"
            
            # マイグレーション実行
            echo "Starting migrations..."
            PGPASSWORD="${POSTGRES_PASSWORD}" psql \
                -h "${DB_HOST}" \
                -U "${POSTGRES_USER}" \
                -d "${POSTGRES_DB}" \
                --set=sslmode=require \
                -f /app/init.sql
            echo "Migrations applied to Supabase"
        else
            echo "ERROR: Failed to connect to Supabase database via URI"
        fi
    fi
    echo "Database migrations completed"
fi

# ホットリロードの制御
if [ "${RUN_HOT_RELOAD}" = "true" ]; then
    echo "Starting with hot-reload using Air..."

    # Airの実行前に必要なディレクトリを作成
    mkdir -p /app/tmp

    # ホットリロード開始
    cd /app && air -c /app/.air.toml
    echo "Hot-reload terminated"
else
    echo "Starting without hot-reload..."
    # 通常モードで実行（引数はDockerfileのCMDから渡される）
    exec "$@"
fi