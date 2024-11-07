# go-postgresql

## 1. データベーステーブルの作成
```sql
-- 既存のmicropostsテーブルがある場合は安全に削除します
DROP TABLE IF EXISTS microposts;

-- ブログ投稿を保存するためのテーブルを作成します
-- id: 自動採番される主キー
-- title: 255文字まで保存できる必須の投稿タイトル
CREATE TABLE microposts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);
```

## 2. テストデータの投入
```sql
-- サンプルとして3つのブログ投稿データを追加します
INSERT INTO microposts (title) VALUES
    ('はじめてのブログ投稿'),
    ('PostgreSQLについて'),
    ('今日の天気について');

-- 追加したデータが正しく保存されているか確認します
SELECT * FROM microposts;
```

## 3. アプリケーションの実行
```powershell
 $env:APP_ENV="development"; go run main.go
```

これで、ブログ投稿を保存するためのデータベース環境が整います。
各コマンドは順番に実行する必要があります。
