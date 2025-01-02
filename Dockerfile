# Go 言語の公式イメージを使用
FROM golang:1.22-alpine

# 作業ディレクトリの設定
WORKDIR /app

# Go モジュールのキャッシュを利用
COPY go.mod go.sum ./
RUN go mod tidy

# ソースコードをコンテナにコピー
COPY . .

# Go アプリケーションをビルド
RUN go build -o /app/cmd/app/main ./cmd/app

# コンテナ起動時に実行するコマンド
CMD ["/app/cmd/app/main"]
