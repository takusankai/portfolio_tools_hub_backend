#!/bin/bash
# filepath: /Volumes/1TB_SSD/DevelopsDirectory/Repositories/portfolio_tools_hub_backend/build/scripts/gcr_deploy.sh

# 引数の取得（Base64エンコードされた値）
IMAGE_NAME=$1
POSTGRES_PASSWORD_B64=$2
SUPABASE_KEY_B64=$3
SUPABASE_URL_B64=$4

# Base64デコード
echo "Decoding base64 encoded secrets..."
POSTGRES_PASSWORD=$(echo -n "$POSTGRES_PASSWORD_B64" | base64 --decode)
SUPABASE_KEY=$(echo -n "$SUPABASE_KEY_B64" | base64 --decode)
SUPABASE_URL=$(echo -n "$SUPABASE_URL_B64" | base64 --decode)

# デバッグ出力（センシティブ情報は部分表示）
echo "Image name: $IMAGE_NAME"
echo "Postgres password: ${POSTGRES_PASSWORD:0:3}****"
echo "Supabase key: ${SUPABASE_KEY:0:10}****"
echo "Supabase URL: $SUPABASE_URL"

# 環境変数を読み込む
echo "Loading environment variables from build/env/prob.env"
if [ -f "build/env/prob.env" ]; then
    grep -v '^#' build/env/prob.env > /tmp/envs.txt
    
    # 環境変数を処理
    ENV_FLAGS=""
    while IFS= read -r line; do
        if [[ ! -z "$line" ]]; then
            # 以下の場合はスキップ
            # - SUPABASEで始まる場合（引数で上書き）
            # - PASSWORDを含む場合（シークレットで管理）
            # - PORTの場合（Cloud Runの予約変数）
            if [[ "$line" != SUPABASE_* && "$line" != *PASSWORD* && "$line" != PORT=* ]]; then
                ENV_FLAGS="${ENV_FLAGS}${line},"
            fi
        fi
    done < /tmp/envs.txt
    
    # 末尾のカンマを削除
    ENV_FLAGS=${ENV_FLAGS%,}
    
    # デコードした引数で渡された値を追加
    if [ ! -z "$SUPABASE_URL" ]; then
        ENV_FLAGS="${ENV_FLAGS},SUPABASE_URL=${SUPABASE_URL}"
    fi
    
    # デプロイコマンド実行
    echo "Deploying with environment variables and secrets"
    gcloud run deploy portfolio-backend \
        --image=$IMAGE_NAME \
        --region=asia-northeast1 \
        --platform=managed \
        --allow-unauthenticated \
        --set-env-vars="$ENV_FLAGS,POSTGRES_PASSWORD=$POSTGRES_PASSWORD,SUPABASE_KEY=$SUPABASE_KEY" \
        --port=8080
else
    echo "ERROR: build/env/prob.env file not found!"
    exit 1
fi
