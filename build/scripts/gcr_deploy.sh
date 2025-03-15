#!/bin/bash
# filepath: /Volumes/1TB_SSD/DevelopsDirectory/Repositories/portfolio_tools_hub_backend/build/scripts/gcr_deploy.sh

# 引数の取得
IMAGE_NAME=$1
POSTGRES_PASSWORD=$2
SUPABASE_KEY=$3
SUPABASE_URL=$4

# 引数の確認
echo "Image name: $IMAGE_NAME"
echo "Supabase URL received: ${SUPABASE_URL:0:10}..."

# 環境変数を読み込む
echo "Loading environment variables from build/env/prob.env"
if [ -f "build/env/prob.env" ]; then
    grep -v '^#' build/env/prob.env > /tmp/envs.txt
    
    # 環境変数を処理
    ENV_FLAGS=""
    while IFS= read -r line; do
        if [[ ! -z "$line" ]]; then
            # SUPABASEで始まる場合はスキップ（引数で上書き）
            if [[ "$line" != SUPABASE_* && "$line" != *PASSWORD* ]]; then
                ENV_FLAGS="${ENV_FLAGS}${line},"
            fi
        fi
    done < /tmp/envs.txt
    
    # 末尾のカンマを削除
    ENV_FLAGS=${ENV_FLAGS%,}
    
    # 引数で渡された値を追加
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
        --set-env-vars="$ENV_FLAGS" \
        --update-secrets="POSTGRES_PASSWORD=POSTGRES_PASSWORD:latest,SUPABASE_KEY=SUPABASE_KEY:latest" \
        --port=8080
else
    echo "ERROR: build/env/prob.env file not found!"
    exit 1
fi
