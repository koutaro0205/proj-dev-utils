# バックエンド用のテンプレートリポジトリ

## ディレクトリ構成（app以下）

```
app/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── graph/
│   │   ├── generated.go
│   │   ├── models_gen.go
│   │   ├── resolver.go
│   │   └── schema.graphqls
│   ├── handlers/
│   │   └── handler.go
│   ├── middleware/
│   │   └── middleware.go
│   ├── models/
│   │   └── models.go
│   ├── repository/
│   │   └── repository.go
│   └── services/
│       └── service.go
├── pkg/
│   └── utils/
│       └── utils.go
├── go.mod
└── go.sum
```

- cmd/server: アプリケーションのエントリーポイント。main.goがここに配置されます。
- internal/gqlgen: gqlgenによって生成されたコードやスキーマファイルを配置します。
- internal/handlers: HTTPリクエストを処理するハンドラーを配置します。
- internal/middleware: ミドルウェアを配置します。
- internal/models: アプリケーションのデータモデルを配置します。
- internal/repository: データベース操作を行うリポジトリを配置します。
- internal/services: ビジネスロジックを含むサービスを配置します。
- pkg/utils: ユーティリティ関数を配置します。
