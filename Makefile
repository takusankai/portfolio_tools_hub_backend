.PHONY: up up-dev down-up up-dev-migrate down-up-prob up-prob up-prob-migrate down down-dev down-prob build ps logs clean db-test db-reset api-test

# デフォルト環境起動（エイリアス）
up:
	$(MAKE) up-dev

# 開発環境: 再起動(エイリアス)
down-up:
	$(MAKE) down
	$(MAKE) up

# 開発環境: マイグレーションなし + ホットリロード
up-dev:
	docker compose --profile dev up -d

# 開発環境: マイグレーション実行 + ホットリロード
up-dev-migrate:
	RUN_MIGRATIONS=true docker compose --profile dev up -d

# 本番環境: 再起動(エイリアス)
down-up-prob:
	$(MAKE) down-prob
	$(MAKE) up-prob

# 本番環境: マイグレーションなし + 通常実行
up-prob:
	docker compose --profile prob up -d

# 本番環境: マイグレーション実行 + 通常実行
up-prob-migrate:
	RUN_MIGRATIONS=true docker compose --profile prob up -d

# デフォルト環境停止
down:
	docker compose --profile dev --profile prob down

# 開発環境: コンテナを停止する
down-dev:
	docker compose --profile dev down

# 本番環境: コンテナを停止する
down-prob:
	docker compose --profile prob down

# Dockerイメージをビルドする
build:
	docker compose build

# 実行中のコンテナを一覧表示
ps:
	docker compose ps

# コンテナからの出力を表示
logs:
	docker compose logs -f

# 未使用のDockerリソースをクリーンアップ
clean:
	docker compose down --remove-orphans --volumes
	docker system prune -f

# 開発環境のローカルデータベースに接続テスト
db-test:
	docker compose --profile dev exec db-dev psql -U postgres_user -d postgres_db -c "SELECT * FROM users;"

# 開発環境のローカルデータベースを down up する
db-reset:
	docker compose --profile dev down -v
	docker compose --profile dev up -d db-dev

# 開発環境のAPIにシンプルなアクセステスト
api-test:
	curl -X GET http://localhost:8080/