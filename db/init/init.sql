CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email) VALUES
('テストユーザー1', 'user1@example.com'),
('テストユーザー2', 'user2@example.com'),
('テストユーザー3', 'user3@example.com');

CREATE TABLE IF NOT EXISTS portfolios (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO portfolios (user_id, name, description) VALUES
(1, 'ポートフォリオ1', 'テストポートフォリオの説明文1'),
(1, 'ポートフォリオ2', 'テストポートフォリオの説明文2'),
(2, 'ポートフォリオ3', 'テストポートフォリオの説明文3');