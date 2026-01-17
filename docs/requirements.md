# Requirements – Go Realtime Chat

## 1. Purpose
本プロジェクトは、Go言語におけるリアルタイム通信の実装を通じて、
以下の技術要素を実践的に学習することを目的とする。

- net/http を用いたHTTPサーバ構築
- WebSocketによる双方向通信
- goroutine / channel を用いた並行処理設計
- Contextを用いたライフサイクル管理
- 責務分離を意識したディレクトリ設計

---

## 2. Scope

### In Scope（対象）
- メモリ上で動作する簡易チャットアプリ
- 複数クライアントによるリアルタイムメッセージ配信
- ローカル環境での動作を前提

### Out of Scope（対象外）
- ユーザー認証・認可
- データベース永続化
- 水平スケーリング
- セキュリティ対策（HTTPS等）

---

## 3. Functional Requirements

### 3.1 Connection
- クライアントはWebSocketでサーバに接続できる
- 接続時にユーザー名を指定できる

### 3.2 Messaging
- クライアントが送信したメッセージは、
  同一ルーム内の全クライアントに即時配信される
- メッセージはJSON形式でやり取りされる

### 3.3 System Events
- ユーザーの接続・切断を通知できる
- サーバからのシステムメッセージを送信できる

---

## 4. Non-Functional Requirements

| 項目 | 内容 |
|---|---|
| 言語 | Go 1.22以上 |
| 通信方式 | WebSocket |
| 同時接続数 | 数十クライアント |
| 実行環境 | ローカル |
| 永続化 | なし |

---

## 5. Architecture Overview

- HTTPサーバは `/ws` エンドポイントを提供する
- 各クライアントは Hub に登録される
- Hub は goroutine 上で動作し、以下を管理する
  - Client の登録・解除
  - メッセージのブロードキャスト

---

## 6. Future Extensions
- ルーム分割
- gRPC streaming による再実装
- フロントエンドのReact化
- メッセージ履歴管理
