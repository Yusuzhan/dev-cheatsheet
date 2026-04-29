---
title: HTTP 协议
icon: fa-globe
primary: "#E44D26"
lang: http
locale: zhs
---

## fa-arrow-right 请求

```http
GET /api/users?page=1&limit=10 HTTP/1.1
Host: example.com
Accept: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9
User-Agent: MyApp/1.0
Connection: keep-alive
```

```http
POST /api/users HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 42

{"name": "Alice", "email": "a@b.com"}
```

## fa-arrow-left 响应

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Content-Length: 85
Cache-Control: max-age=3600
X-Request-Id: abc123

{"id": 1, "name": "Alice", "email": "a@b.com"}
```

## fa-list 请求方法

```text
GET       获取资源                         安全、幂等
POST      创建资源 / 触发操作
PUT       替换整个资源                      幂等
PATCH     部分更新资源
DELETE    删除资源                          幂等
HEAD      同 GET，不返回响应体               安全、幂等
OPTIONS   获取支持的方法与 CORS 信息          安全、幂等
TRACE     回显请求                          安全、幂等
```

## fa-circle-half-stroke 状态码

```text
200 OK                    请求成功
201 Created               资源已创建 (POST/PUT)
204 No Content            成功，无响应体 (DELETE)
301 Moved Permanently     永久重定向
302 Found                 临时重定向
304 Not Modified          缓存未修改
307 Temporary Redirect    临时重定向，保持方法不变
308 Permanent Redirect    永久重定向，保持方法不变

400 Bad Request           请求格式错误
401 Unauthorized          未认证
403 Forbidden             无权限
404 Not Found             资源不存在
405 Method Not Allowed    方法不允许
409 Conflict              状态冲突
422 Unprocessable Entity  验证失败
429 Too Many Requests     请求频率超限

500 Internal Server Error 服务器内部错误
502 Bad Gateway           上游网关错误
503 Service Unavailable   服务不可用 / 维护中
504 Gateway Timeout       上游网关超时
```

## fa-file-lines 常用请求头

```http
Content-Type: application/json; charset=utf-8
Content-Type: text/html; charset=utf-8
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary
Content-Type: application/x-www-form-urlencoded
Content-Type: text/event-stream

Accept: application/json, text/html, */*
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9,zh-CN;q=0.8

Authorization: Basic base64(user:pass)
Authorization: Bearer <token>

Cache-Control: no-cache
Cache-Control: max-age=3600, public
Cache-Control: no-store, must-revalidate

Connection: keep-alive
Connection: close
```

## fa-cookie-bite Cookie

```http
Set-Cookie: session=abc123; Path=/; HttpOnly; Secure; SameSite=Strict
Set-Cookie: lang=en; Max-Age=86400; Domain=.example.com
Set-Cookie: theme=dark; Expires=Wed, 09 Jun 2026 10:18:14 GMT

Cookie: session=abc123; lang=en; theme=dark
```

```text
HttpOnly     禁止 JS 访问（防 XSS）
Secure       仅 HTTPS 传输
SameSite     Strict | Lax | None
Domain       Cookie 作用域
Path         URL 前缀作用域
Max-Age      有效期（秒）
Expires      绝对过期时间
```

## fa-shield-halved 跨域 CORS

```http
# 预检请求 (OPTIONS)
OPTIONS /api/data HTTP/1.1
Origin: https://app.example.com
Access-Control-Request-Method: POST
Access-Control-Request-Headers: Content-Type, Authorization

# 预检响应
HTTP/1.1 204 No Content
Access-Control-Allow-Origin: https://app.example.com
Access-Control-Allow-Methods: GET, POST, PUT, DELETE
Access-Control-Allow-Headers: Content-Type, Authorization
Access-Control-Max-Age: 86400
Access-Control-Allow-Credentials: true

# 简单响应
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: X-Total-Count
```

## fa-lock 认证方式

```http
# Basic 基础认证
Authorization: Basic dXNlcjpwYXNz

# Bearer Token
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0

# API Key 请求头
X-API-Key: sk_live_abc123def456

# API Key 查询参数
GET /api/data?api_key=sk_live_abc123def456

# Digest 摘要认证
Authorization: Digest username="user", realm="api", nonce="abc", uri="/api", response="xyz"
```

## fa-cloud 缓存

```http
Cache-Control: no-store                # 不缓存
Cache-Control: no-cache                # 每次验证
Cache-Control: max-age=3600, public    # 缓存1小时
Cache-Control: max-age=0, must-revalidate  # 必须验证

ETag: "33a64df551425fcc55e4d42a148795d9f25f89d4"  # 强 ETag
ETag: W/"0815"                                        # 弱 ETag

If-None-Match: "33a64df551425fcc55e4d42a148795d9f25f89d4"
If-Modified-Since: Wed, 21 Oct 2025 07:28:00 GMT

Last-Modified: Wed, 21 Oct 2025 07:28:00 GMT
Vary: Accept-Encoding
Age: 1200                              # 已缓存时间（秒）
```

## fa-arrow-down-wide-short 内容协商

```http
Accept: text/html, application/json;q=0.9, */*;q=0.1
Accept-Language: en-US,en;q=0.9,zh-CN;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Charset: utf-8

# 服务端返回选定的变体
Content-Type: text/html; charset=utf-8
Content-Language: en-US
Content-Encoding: br
Vary: Accept-Encoding, Accept-Language
```

## fa-link 连接管理

```http
# Keep-Alive 长连接
Connection: keep-alive
Keep-Alive: timeout=5, max=100

# 升级到 WebSocket
GET /ws HTTP/1.1
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
Sec-WebSocket-Version: 13

# HTTP/2 伪头
:method GET
:path /api/users
:scheme https
:authority example.com
```

## fa-arrows-rotate 重定向

```http
HTTP/1.1 301 Moved Permanently
Location: https://new.example.com/page

HTTP/1.1 302 Found
Location: /login

HTTP/1.1 307 Temporary Redirect
Location: https://other.example.com/api
# POST 仍为 POST（不同于 302）

HTTP/1.1 308 Permanent Redirect
Location: https://new.example.com/api
# POST 仍为 POST（不同于 301）
```

## fa-file-upload Multipart 与传输

```http
POST /upload HTTP/1.1
Content-Type: multipart/form-data; boundary=----Boundary

------Boundary
Content-Disposition: form-data; name="title"

My Document
------Boundary
Content-Disposition: form-data; name="file"; filename="doc.pdf"
Content-Type: application/pdf

<binary data>
------Boundary--
```

```http
# 分块传输
Transfer-Encoding: chunked

1a
abcdefghijklmnopqrstuvwxyz
0
```

## fa-bolt HTTP/2 与 HTTP/3

```text
HTTP/2 特性：
  - 二进制帧（非文本协议）
  - 单连接多路复用
  - 头部压缩 (HPACK)
  - 服务端推送
  - 流优先级与流量控制

HTTP/3 特性：
  - 基于 QUIC 传输（UDP）
  - 0-RTT 连接建立
  - 无队头阻塞
  - 连接迁移（IP 变化不断开）
  - 内置 TLS 1.3
```

```http
# HTTP/2 服务端推送（已废弃但值得了解）
Push-Promise: /style.css
```
