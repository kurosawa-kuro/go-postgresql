# go-postgresql

-- テーブルが既に存在する場合は削除
DROP TABLE IF EXISTS microposts;

-- テーブルの作成
CREATE TABLE microposts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

-- シードデータの挿入
INSERT INTO microposts (title) VALUES
    ('はじめてのブログ投稿'),
    ('PostgreSQLについて'),
    ('今日の天気について');

-- データの確認
SELECT * FROM microposts;