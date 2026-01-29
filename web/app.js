const log = document.getElementById("log");
const input = document.getElementById("input");
const sendBtn = document.getElementById("sendBtn");

// URL クエリから username / room を取得
const params = new URLSearchParams(location.search);
const username = params.get("username") || "test";
const room = params.get("room") || "general";

let ws = null;
let reconnectAttempts = 0;
let reconnectTimer = null;
let connected = false;

// log
function logLine(text) {
  log.textContent += text + "\n";
}

// 再接続待機時間
function getReconnectDelay() {
  const base = 1000;
  const max = 30000;
  const delay = base * Math.pow(2, reconnectAttempts);
  return Math.min(delay, max);
}

// WebSocket 接続
function connect() {
  const ws = new WebSocket(
    `ws://localhost:8080/ws?username=${encodeURIComponent(username)}&room=${encodeURIComponent(room)}`,
  );

  // 接続確認
  ws.onopen = () => {
    log.textContent += `[connected] room=${room}\n`;
  };

  // 受信処理（1 回だけ定義）
  ws.onmessage = (event) => {
    const msg = JSON.parse(event.data);

    switch (msg.type) {
      case "join":
        log.textContent += `🟢 ${msg.username} joined\n`;
        break;
      case "leave":
        log.textContent += `🔴 ${msg.username} left\n`;
        break;
      case "message":
        log.textContent += `${msg.username}: ${msg.content}\n`;
        break;
      case "system":
        log.textContent += `[system] ${msg.content}\n`;
        break;
    }
  };

  ws.onclose = () => {
    connected = false;
    logLine("[disconnected]");
    scheduleReconnect();
  };
  ws.onerror = () => {
    ws.close();
  };
}

// 再接続スケジュール
function scheduleReconnect() {
  if (reconnectTimer) return;

  reconnectAttempts++;
  const delay = getReconnectDelay();

  logLine(`[reconnect in ${delay / 1000}s]`);

  reconnectTimer = setTimeout(() => {
    reconnectTimer = null;
    connect();
  }, delay);
}

// 送信処理（JSON）
sendBtn.onclick = () => {
  if (!connected || !ws) return;
  if (!input.value) return;

  ws.send(
    JSON.stringify({
      content: input.value,
    }),
  );

  input.value = "";
};

connect();
