## 環境構築方法
### 1. クローンの生成
```shell
    git clone https://github.com/omusobadon/Go_TaskScheduler.git
```
### 2. 環境変数ファイルの作成
    GoogleDrive の共有フォルダにある環境変数ファイル .env を Go_TaskScheduler/ へコピー

### 3. Prisma-Client-Goのインストール
```shell
    go get github.com/steebchen/prisma-client-go
```

### 4. /Go_APIServer内で以下のコマンドを実行してDBを同期（DB操作用のパッケージが生成される）
```shell
    go run github.com/steebchen/prisma-client-go db push
```

## ファイル一覧
- db/           prisma-client-goが作成したフォルダ。DB操作用のパッケージ等
- GetTime       時刻同期処理
- scheduler     スケジュールを管理して指定時刻にtaskを実行
- schema        prismaの設定ファイル。DBのURLやテーブルの定義など
- task          schedulerによって呼び出されるタスク処理