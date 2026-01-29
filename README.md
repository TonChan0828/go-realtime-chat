# go-realtime-chat

Go + WebSocket で動くシンプルなリアルタイムチャットのサンプルです。ルーム単位でメッセージをブロードキャストし、接続・切断イベントも配信します。フロントは最小の HTML/JS で動作します。

## 概要
- HTTP サーバで `/ws` の WebSocket を提供
- ルームごとに Hub を生成し、接続中クライアントへ一斉配信
- 送信データは JSON、受信時にサーバで `type/username/timestamp` を付与
- Ping/Pong と deadline 設定で接続維持
- Context と OS シグナルで graceful shutdown

## 現在の構成と使用方法
構成:
- サーバ: Go + Gorilla WebSocket（`/ws` エンドポイント）
- クライアント: `web/index.html` + `web/app.js`（最小構成の HTML/JS）

使い方:
1) サーバ起動
```bash
go run ./cmd/server/main.go
```
2) ブラウザでアクセス（`username` と `room` を指定）
```
http://localhost:8080/?username=alice&room=general
```
3) 別タブで `username` / `room` を変えて開くと、同一ルーム内でブロードキャストされます。

接続先 WebSocket:
```
ws://localhost:8080/ws?username={name}&room={room}
```

## 機能
- ルーム単位のメッセージブロードキャスト
- 参加/退出イベントの配信（`join` / `leave`）
- メッセージ送受信（JSON）
- 再接続リトライ（指数バックオフ上限あり）
- Ping/Pong と deadline による接続維持

## 起動方法
```bash
go run ./cmd/server/main.go
```

ブラウザで以下にアクセスします。
```
http://localhost:8080/?username=alice&room=general
```

別のタブで `username` と `room` を変えると、同一ルーム内でのブロードキャストが確認できます。

## 使い方
- 画面の入力欄にメッセージを入力して Send
- 受信ログに `join/leave/message/system` が表示

## WebSocket 仕様
接続:
```
ws://localhost:8080/ws?username={name}&room={room}
```

クライアント → サーバ:
```json
{"content":"hello"}
```

サーバ → クライアント:
```json
{
  "type": "message",
  "username": "alice",
  "content": "hello",
  "timestamp": "2025-01-01T12:00:00Z"
}
```

イベント種別:
- `join`: 接続通知
- `leave`: 切断通知
- `message`: チャットメッセージ
- `system`: システム通知（予約枠）

## ディレクトリ構成
```
cmd/server/main.go         エントリポイント、HTTP/静的配信、終了処理
internal/handler           WebSocket ハンドラ
internal/hub               Hub / Client / Room 管理
internal/model             メッセージ構造体
web/                       簡易フロント（index.html, app.js）
docs/requirements.md       要件メモ
```

## アーキテクチャ概要
- `RoomManager` がルーム名ごとに `Hub` を生成
- `Hub` は `register/unregister/broadcast` の各チャネルを持ち、1 goroutine で管理
- `Client` は ReadPump/WritePump で読み書きを分離
- `Hub.Register/Unregister` が join/leave をブロードキャスト

## 制約と注意点
- 認証/認可なし
- 永続化なし（メッセージ履歴は保持しない）
- ローカル利用前提
- `CheckOrigin` は常に `true`（セキュリティ用途では未対応）

## 依存
- Go: `go.mod` では `go 1.25.0`
- WebSocket: `github.com/gorilla/websocket`

## 今後の拡張アイデア
- ルーム一覧 API と管理 UI
- メッセージ履歴の保存
- 認証/認可
- フロントエンドのリッチ化
