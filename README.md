# 分散システムデモ (Go, React, GraphQL, gRPC)

このプロジェクトは、モダンな技術スタックを使用したシンプルな分散システムのデモです。
ユーザーとそれに紐づく投稿（Post）を管理する機能を提供します。

## アーキテクチャ概要

システムは以下の4つの主要コンポーネントで構成されています。

1.  **Frontend (React)**:
    -   ユーザーインターフェースを提供します。
    -   Apollo Clientを使用してBFFのGraphQL APIを叩きます。
2.  **BFF (Backend For Frontend - Go)**:
    -   GraphQLサーバーとして機能します (`gqlgen` 使用)。
    -   フロントエンドからのリクエストを受け、バックエンドのgRPCサービスを呼び出してデータを集約・整形します。
3.  **User Service (Go)**:
    -   gRPCサーバーとして機能します。
    -   **Clean Architecture**を採用しています。
    -   ユーザー情報の管理を担当します。
4.  **Post Service (Go)**:
    -   gRPCサーバーとして機能します。
    -   **Clean Architecture**を採用しています。
    -   投稿情報の管理を担当します。

```mermaid
graph LR
    Client[Frontend (React)] -- GraphQL --> BFF[BFF (Go)]
    BFF -- gRPC --> UserSvc[User Service]
    BFF -- gRPC --> PostSvc[Post Service]
    UserSvc --> DB[(PostgreSQL)]
    PostSvc --> DB[(PostgreSQL)]
```

## データ構造

主なデータモデルは以下の通りです。

### User (ユーザー)
-   **ID**: 一意の識別子 (UUID)
-   **Name**: ユーザー名
-   **Email**: メールアドレス

### Post (投稿)
-   **ID**: 一意の識別子 (UUID)
-   **Title**: タイトル
-   **Content**: 本文
-   **UserID**: 投稿者のユーザーID (Userモデルとの紐付け)

## ディレクトリ構造

モノレポ構成を採用しています。

-   `user-service/`: ユーザー管理のgRPCマイクロサービス (Clean Architecture)
    -   `domain/`: エンティティとインターフェース定義
    -   `internal/repository/`: データアクセス (GORM)
    -   `internal/usecase/`: ビジネスロジック
    -   `internal/delivery/`: ハンドラー (gRPC)
-   `post-service/`: 投稿管理のgRPCマイクロサービス (Clean Architecture)
    -   `domain/`: エンティティとインターフェース定義
    -   `internal/repository/`: データアクセス (GORM)
    -   `internal/usecase/`: ビジネスロジック
    -   `internal/delivery/`: ハンドラー (gRPC)
-   `bff/`: GraphQL BFFサーバーのコード
-   `frontend/`: Reactフロントエンドのコード
-   `k8s/`: Kubernetesマニフェストファイル

## 実行方法 (ローカル開発)

### 前提条件
-   Go 1.24+
-   Node.js 18+
-   Docker & Kubernetes (デプロイ確認用)

### 手順 (推奨: Docker Compose)

Docker Composeを使用すると、データベースを含む全サービスを一括で起動できます。

```bash
docker compose up --build
```

-   Frontend: `http://localhost:5173`
-   GraphQL Playground: `http://localhost:8080`

<details>
<summary><strong>手順 (手動実行)</strong></summary>

個別にサービスを立ち上げる場合の手順です。PostgreSQLは別途起動しておく必要があります。

1.  **User Serviceの起動**:
    ```bash
    cd user-service
    go run main.go
    ```
    `localhost:50051` で起動します。

2.  **Post Serviceの起動**:
    別のターミナルで実行してください。
    ```bash
    cd post-service
    go run main.go
    ```
    `localhost:50052` で起動します。

3.  **BFFの起動**:
    別のターミナルで実行してください。
    ```bash
    cd bff
    go run server.go
    ```
    `localhost:8080` で起動します。

4.  **Frontendの起動**:
    別のターミナルで実行してください。
    ```bash
    cd frontend
    npm install
    npm run dev
    ```
    ブラウザで表示されるURL（通常 `http://localhost:5173`）にアクセスしてください。

</details>

## Kubernetesへのデプロイ (ローカル)

1.  **Dockerイメージのビルド**:
    ```bash
    docker build -t user-service:latest -f user-service/Dockerfile .
    docker build -t post-service:latest -f post-service/Dockerfile .
    docker build -t bff:latest -f bff/Dockerfile .
    docker build -t frontend:latest -f frontend/Dockerfile .
    ```

2.  **マニフェストの適用**:
    ```bash
    kubectl apply -f k8s/user-service.yaml
    kubectl apply -f k8s/post-service.yaml
    kubectl apply -f k8s/bff.yaml
    kubectl apply -f k8s/frontend.yaml
    ```

3.  **アクセス**:
    -   Frontend: `http://localhost` (LoadBalancerが機能する場合) またはポートフォワードを行ってください。
