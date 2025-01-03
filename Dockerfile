# Go言語の公式イメージを使用
FROM golang:1.23-alpine

# 作業ディレクトリの設定
WORKDIR /app

# Goモジュールのキャッシュを利用
# go.sumはあとで生成されるはず
# COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod tidy

# airをインストール
RUN go install github.com/air-verse/air@latest

# ソースコードをコンテナにコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o /app/cmd/app/main ./cmd/app

# コンテナ起動時に実行するコマンド
CMD ["/app/cmd/app/main"]
