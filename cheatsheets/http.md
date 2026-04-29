---
title: HTTP
icon: fa-globe
primary: "#E44D26"
lang: http
---

## fa-arrow-right Request

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

## fa-arrow-left Response

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Content-Length: 85
Cache-Control: max-age=3600
X-Request-Id: abc123

{"id": 1, "name": "Alice", "email": "a@b.com"}
```

## fa-list Methods

```text
GET       Retrieve resource              Safe, Idempotent
POST      Create resource / trigger action
PUT       Replace entire resource         Idempotent
PATCH     Partial update resource
DELETE    Remove resource                 Idempotent
HEAD      Same as GET, body excluded      Safe, Idempotent
OPTIONS   Supported methods & CORS        Safe, Idempotent
TRACE     Echo request back               Safe, Idempotent
```

## fa-circle-half-stroke Status Codes

```text
200 OK                    Request succeeded
201 Created               Resource created (POST/PUT)
204 No Content            Success, no body (DELETE)
301 Moved Permanently     Permanent redirect
302 Found                 Temporary redirect
304 Not Modified          Cache hit
307 Temporary Redirect    Preserve method
308 Permanent Redirect    Preserve method

400 Bad Request           Invalid syntax
401 Unauthorized          Authentication required
403 Forbidden             No permission
404 Not Found             Resource missing
405 Method Not Allowed    Wrong HTTP method
409 Conflict              State conflict
422 Unprocessable Entity  Validation failed
429 Too Many Requests     Rate limited

500 Internal Server Error
502 Bad Gateway           Upstream error
503 Service Unavailable   Overloaded / maintenance
504 Gateway Timeout       Upstream timeout
```

## fa-file-lines Headers

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

## fa-cookie-bite Cookies

```http
Set-Cookie: session=abc123; Path=/; HttpOnly; Secure; SameSite=Strict
Set-Cookie: lang=en; Max-Age=86400; Domain=.example.com
Set-Cookie: theme=dark; Expires=Wed, 09 Jun 2026 10:18:14 GMT

Cookie: session=abc123; lang=en; theme=dark
```

```text
HttpOnly     No JS access (XSS protection)
Secure       HTTPS only
SameSite     Strict | Lax | None
Domain       Cookie scope
Path         URL prefix scope
Max-Age      Seconds until expiry
Expires      Absolute expiry date
```

## fa-shield-halved CORS

```http
# Preflight request (OPTIONS)
OPTIONS /api/data HTTP/1.1
Origin: https://app.example.com
Access-Control-Request-Method: POST
Access-Control-Request-Headers: Content-Type, Authorization

# Preflight response
HTTP/1.1 204 No Content
Access-Control-Allow-Origin: https://app.example.com
Access-Control-Allow-Methods: GET, POST, PUT, DELETE
Access-Control-Allow-Headers: Content-Type, Authorization
Access-Control-Max-Age: 86400
Access-Control-Allow-Credentials: true

# Simple response
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: X-Total-Count
```

## fa-lock Authentication

```http
# Basic Auth
Authorization: Basic dXNlcjpwYXNz

# Bearer Token
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0

# API Key in header
X-API-Key: sk_live_abc123def456

# API Key in query
GET /api/data?api_key=sk_live_abc123def456

# Digest Auth
Authorization: Digest username="user", realm="api", nonce="abc", uri="/api", response="xyz"
```

## fa-cloud Caching

```http
Cache-Control: no-store
Cache-Control: no-cache
Cache-Control: max-age=3600, public
Cache-Control: max-age=0, must-revalidate

ETag: "33a64df551425fcc55e4d42a148795d9f25f89d4"
ETag: W/"0815"

If-None-Match: "33a64df551425fcc55e4d42a148795d9f25f89d4"
If-Modified-Since: Wed, 21 Oct 2025 07:28:00 GMT

Last-Modified: Wed, 21 Oct 2025 07:28:00 GMT
Vary: Accept-Encoding
Age: 1200
```

## fa-arrow-down-wide-short Content Negotiation

```http
Accept: text/html, application/json;q=0.9, */*;q=0.1
Accept-Language: en-US,en;q=0.9,zh-CN;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Charset: utf-8

# Server responds with chosen variant
Content-Type: text/html; charset=utf-8
Content-Language: en-US
Content-Encoding: br
Vary: Accept-Encoding, Accept-Language
```

## fa-link Connection

```http
# Keep-Alive
Connection: keep-alive
Keep-Alive: timeout=5, max=100

# Upgrade to WebSocket
GET /ws HTTP/1.1
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
Sec-WebSocket-Version: 13

# HTTP/2 pseudo headers
:method GET
:path /api/users
:scheme https
:authority example.com
```

## fa-arrows-rotate Redirect

```http
HTTP/1.1 301 Moved Permanently
Location: https://new.example.com/page

HTTP/1.1 302 Found
Location: /login

HTTP/1.1 307 Temporary Redirect
Location: https://other.example.com/api
# POST stays POST (unlike 302)

HTTP/1.1 308 Permanent Redirect
Location: https://new.example.com/api
# POST stays POST (unlike 301)
```

## fa-file-upload Multipart

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
# Chunked transfer
Transfer-Encoding: chunked

1a
abcdefghijklmnopqrstuvwxyz
0
```

## fa-bolt HTTP/2 & HTTP/3

```text
HTTP/2 Features:
  - Binary framing (not text-based)
  - Multiplexed streams over one connection
  - Header compression (HPACK)
  - Server push
  - Stream priorities & flow control

HTTP/3 Features:
  - QUIC transport (UDP-based)
  - 0-RTT connection setup
  - No head-of-line blocking
  - Connection migration (IP change)
  - Built-in TLS 1.3
```

```http
# HTTP/2 server push (deprecated but notable)
Push-Promise: /style.css
```
