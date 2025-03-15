#!/bin/bash
# filepath: /Volumes/1TB_SSD/DevelopsDirectory/Repositories/portfolio_tools_hub_backend/deploy.sh

# 環境変数を読み込む
echo "Loading environment variables from build/env/prob.env"
grep -v '^#' build/env/prob.env > /tmp/envs.txt

# 環境変数を処理
ENV_FLAGS=""
while IFS= read -r line; do
    if [[ ! -z "$line" ]]; then
        ENV_FLAGS="${ENV_FLAGS}${line},"
    fi
done < /tmp/envs.txt
ENV_FLAGS=${ENV_FLAGS%,}  # 末尾のカンマを削除

# デプロイコマンド実行
echo "Deploying with environment variables: $ENV_FLAGS"
gcloud run deploy portfolio-backend \
    --image=$1 \
    --region=asia-northeast1 \
    --platform=managed \
    --allow-unauthenticated \
    --set-env-vars="$ENV_FLAGS" \
    --update-secrets="POSTGRES_PASSWORD=POSTGRES_PASSWORD:latest,SUPABASE_KEY=SUPABASE_KEY:latest" \
    --port=8080
