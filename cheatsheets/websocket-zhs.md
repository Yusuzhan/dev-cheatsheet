---
title: WebSocket
icon: fa-plug
primary: "#6E00FF"
lang: javascript
locale: zhs
---

## fa-globe Browser Client API

```javascript
const ws = new WebSocket("ws://localhost:8080/ws");

ws.onopen = () => {
  console.log("connected");
  ws.send(JSON.stringify({ type: "hello" }));
};

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log("received:", data);
};

ws.onerror = (err) => console.error("error:", err);

ws.onclose = (event) => {
  console.log("closed:", event.code, event.reason);
};

ws.close(1000, "done");
```

## fa-server Server (Node.js ws)

```javascript
import { WebSocketServer } from "ws";

const wss = new WebSocketServer({ port: 8080 });

wss.on("connection", (ws, req) => {
  const ip = req.headers["x-forwarded-for"] || req.socket.remoteAddress;

  ws.on("message", (data) => {
    const msg = JSON.parse(data);
    wss.clients.forEach((client) => {
      if (client !== ws && client.readyState === WebSocket.OPEN) {
        client.send(JSON.stringify(msg));
      }
    });
  });

  ws.on("close", () => console.log("client disconnected"));
});

console.log("WebSocket server on :8080");
```

## fa-code Server (Go gorilla/websocket)

```go
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWS(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    for {
        mt, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        conn.WriteMessage(mt, msg)
    }
}

http.Handle("/ws", http.HandlerFunc(handleWS))
http.ListenAndServe(":8080", nil)
```

## fa-circle-nodes Connection Lifecycle

```javascript
const ws = new WebSocket("ws://localhost:8080");

ws.addEventListener("open", () => {
  console.log("readyState:", ws.readyState); // 1 = OPEN
});

ws.addEventListener("close", (e) => {
  console.log("code:", e.code);     // 1000 = normal
  console.log("reason:", e.reason); // custom reason
  console.log("wasClean:", e.wasClean);
});

// ws.readyState: 0=CONNECTING, 1=OPEN, 2=CLOSING, 3=CLOSED
```

## fa-paper-plane Sending & Receiving

```javascript
// 文本消息
ws.send("hello");
ws.send(JSON.stringify({ action: "chat", text: "hi" }));

// 二进制消息
const buffer = new Uint8Array([1, 2, 3, 4]);
ws.send(buffer);

const blob = new Blob(["binary data"], { type: "application/octet-stream" });
ws.send(blob);

// 接收消息
ws.binaryType = "arraybuffer"; // 或 "blob"
ws.onmessage = (e) => {
  if (typeof e.data === "string") {
    const msg = JSON.parse(e.data);
  } else {
    const arr = new Uint8Array(e.data);
  }
};
```

## fa-heart-pulse Ping/Pong & Heartbeat

```javascript
// 浏览器：无原生 ping API，使用应用层心跳
let heartbeatInterval;

ws.onopen = () => {
  heartbeatInterval = setInterval(() => {
    if (ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "ping" }));
    }
  }, 30000);
};

ws.onclose = () => clearInterval(heartbeatInterval);

// Node.js 服务端：原生 ping/pong
wss.on("connection", (ws) => {
  ws.isAlive = true;
  ws.on("pong", () => { ws.isAlive = true; });
});

const interval = setInterval(() => {
  wss.clients.forEach((ws) => {
    if (!ws.isAlive) return ws.terminate();
    ws.isAlive = false;
    ws.ping();
  });
}, 30000);

wss.on("close", () => clearInterval(interval));
```

## fa-rotate Reconnection Strategy

```javascript
function createReconnectingWebSocket(url) {
  let ws;
  let retries = 0;
  const maxRetries = 10;
  const maxDelay = 30000;

  function connect() {
    ws = new WebSocket(url);

    ws.onopen = () => {
      retries = 0;
      console.log("connected");
    };

    ws.onmessage = (e) => onMessage(JSON.parse(e.data));

    ws.onclose = () => {
      if (retries < maxRetries) {
        const delay = Math.min(1000 * Math.pow(2, retries), maxDelay);
        retries++;
        setTimeout(connect, delay);
      }
    };

    ws.onerror = () => {};
  }

  function send(data) {
    if (ws?.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify(data));
    }
  }

  connect();
  return { send, close: () => ws?.close() };
}
```

## fa-users Rooms / Channels

```javascript
const rooms = new Map();

function joinRoom(roomId, ws) {
  if (!rooms.has(roomId)) rooms.set(roomId, new Set());
  rooms.get(roomId).add(ws);
  ws.roomId = roomId;
}

function leaveRoom(ws) {
  const room = rooms.get(ws.roomId);
  if (room) {
    room.delete(ws);
    if (room.size === 0) rooms.delete(ws.roomId);
  }
}

function broadcastToRoom(roomId, message, excludeWs) {
  const room = rooms.get(roomId);
  if (!room) return;
  const data = JSON.stringify(message);
  room.forEach((ws) => {
    if (ws !== excludeWs && ws.readyState === WebSocket.OPEN) {
      ws.send(data);
    }
  });
}
```

## fa-file-zipper Binary Data

```javascript
// 以二进制方式发送文件
ws.binaryType = "arraybuffer";

async function sendFile(ws, file) {
  const buffer = await file.arrayBuffer();
  ws.send(buffer);
}

// 通过 WebSocket 使用 Protocol Buffers / MessagePack
import { encode, decode } from "msgpack-lite";

ws.send(encode({ type: "data", payload: [1, 2, 3] }));

ws.onmessage = (e) => {
  const decoded = decode(new Uint8Array(e.data));
};
```

## fa-lock Authentication

```javascript
// 客户端：连接后发送 token
ws.onopen = () => {
  ws.send(JSON.stringify({ type: "auth", token: localStorage.getItem("token") }));
};

// 服务端：连接时验证
wss.on("connection", (ws, req) => {
  const params = new URL(req.url, `http://${req.headers.host}`).searchParams;
  const token = params.get("token");

  try {
    ws.user = jwt.verify(token, SECRET);
  } catch {
    ws.close(4001, "unauthorized");
  }
});

// 或在 HTTP upgrade 前认证
const server = http.createServer(app);
server.on("upgrade", (req, socket, head) => {
  const token = new URL(req.url, `http://${req.headers.host}`).searchParams.get("token");
  if (!isValid(token)) {
    socket.write("HTTP/1.1 401\r\n\r\n");
    socket.destroy();
    return;
  }
  wss.handleUpgrade(req, socket, head, (ws) => wss.emit("connection", ws, req));
});
```

## fa-triangle-exclamation Error Handling

```javascript
ws.onerror = (event) => {
  console.error("WebSocket error:", event.message);
};

ws.onclose = (event) => {
  switch (event.code) {
    case 1000: console.log("正常关闭"); break;
    case 1001: console.log("终端离开"); break;
    case 1006: console.log("异常关闭 (无 close 帧)"); break;
    case 1008: console.log("策略违规"); break;
    case 1011: console.log("服务器内部错误"); break;
    case 4001: console.log("自定义: 未授权"); break;
  }
};

// 服务端：带状态码关闭
ws.close(1008, "invalid message format");
ws.close(4001, "authentication expired");
```

## fa-network-wired Proxy & Load Balancer

```nginx
# Nginx WebSocket 反向代理
location /ws {
    proxy_pass http://backend;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    proxy_set_header Host $host;
    proxy_read_timeout 86400s;
    proxy_send_timeout 86400s;
}
```

```javascript
// HAProxy
// frontend ws_front
//   bind *:80
//   default_backend ws_back
//   timeout client 86400s
//
// backend ws_back
//   server s1 127.0.0.1:8080
//   timeout server 86400s
//   timeout connect 5s
```

## fa-bolt Socket.IO

```javascript
import { Server } from "socket.io";

const io = new Server(httpServer, {
  cors: { origin: "*" },
});

io.on("connection", (socket) => {
  console.log("user connected:", socket.id);

  socket.join("room-123");
  io.to("room-123").emit("message", { text: "hello room" });

  socket.on("chat", (msg) => {
    io.emit("chat", { from: socket.id, ...msg });
  });

  socket.on("disconnect", (reason) => {
    console.log("disconnected:", reason);
  });
});

// 客户端
import { io } from "socket.io-client";
const socket = io("http://localhost:3000");
socket.emit("chat", { text: "hi" });
socket.on("chat", (msg) => console.log(msg));
```

## fa-shield-halved Security Considerations

```javascript
// 限流
const messageCounts = new Map();
wss.on("connection", (ws) => {
  const messages = [];
  ws.on("message", (data) => {
    const now = Date.now();
    messages.push(now);
    while (messages.length && messages[0] < now - 60000) messages.shift();
    if (messages.length > 100) {
      return ws.close(1008, "rate limit exceeded");
    }
    handleMessage(ws, data);
  });
});

// 输入校验
ws.on("message", (data) => {
  let msg;
  try {
    msg = JSON.parse(data);
  } catch {
    return ws.close(1008, "invalid JSON");
  }
  if (msg.type !== "chat" && msg.type !== "ping") {
    return ws.close(1008, "unknown message type");
  }
});

// Origin 检查
const wss = new WebSocketServer({
  port: 8080,
  verifyClient: (info) => {
    return info.origin === "https://example.com";
  },
});
```
