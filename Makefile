.PHONY: up down build ps logs clean

# Dockerコンテナをバックグラウンドで起動
up:
	docker compose --profile dev up -d

# Dockerコンテナを停止し削除する
down:
	docker compose --profile dev down

# Dockerイメージをビルドまたは再ビルドする
build:
	docker-compose build

# 実行中のコンテナを一覧表示
ps:
	docker-compose ps

# コンテナからの出力を表示
logs:
	docker-compose logs -f

# 未使用のDockerリソースをクリーンアップ
clean:
	docker-compose down --remove-orphans --volumes
	docker system prune -f

# データベースに接続
sample:
	docker exec -it portfolio_tools_hub_db_dev psql -U portfolio_user -d portfolio_db -c "SELECT * FROM users;"