---
title: JWT
icon: fa-id-badge
primary: "#D63AFF"
lang: json
locale: zhs
---

## fa-layer-group JWT 结构（Header/Payload/Signature）

```
header.payload.signature
```

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

```json
{
  "sub": "1234567890",
  "name": "Alice",
  "iat": 1516239022,
  "exp": 1516242622
}
```

```bash
echo -n '{"alg":"HS256","typ":"JWT"}' | base64url
echo -n '{"sub":"1234567890","name":"Alice","iat":1516239022}' | base64url
```

## fa-shield-halved HMAC（HS256）签名

```bash
echo -n "header.payload" | openssl dgst -sha256 -hmac "your-secret-key" -binary | base64url
```

```javascript
const jwt = require("jsonwebtoken");
const token = jwt.sign({ sub: "123", name: "Alice" }, "secret", {
  algorithm: "HS256",
  expiresIn: "1h",
});
const decoded = jwt.verify(token, "secret");
```

```python
import jwt
token = jwt.encode({"sub": "123", "name": "Alice"}, "secret", algorithm="HS256")
decoded = jwt.decode(token, "secret", algorithms=["HS256"])
```

## fa-lock RSA（RS256）签名

```bash
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem
```

```javascript
const fs = require("fs");
const jwt = require("jsonwebtoken");
const priv = fs.readFileSync("private.pem");
const pub = fs.readFileSync("public.pem");
const token = jwt.sign({ sub: "123" }, priv, { algorithm: "RS256" });
const decoded = jwt.verify(token, pub, { algorithms: ["RS256"] });
```

```python
import jwt
token = jwt.encode({"sub": "123"}, open("private.pem").read(), algorithm="RS256")
decoded = jwt.decode(token, open("public.pem").read(), algorithms=["RS256"])
```

## fa-signature ECDSA（ES256）签名

```bash
openssl ecparam -genkey -name prime256v1 -noout -out ec-private.pem
openssl ec -in ec-private.pem -pubout -out ec-public.pem
```

```javascript
const jwt = require("jsonwebtoken");
const token = jwt.sign({ sub: "123" }, ecPrivate, { algorithm: "ES256" });
const decoded = jwt.verify(token, ecPublic, { algorithms: ["ES256"] });
```

```python
import jwt
token = jwt.encode({"sub": "123"}, open("ec-private.pem").read(), algorithm="ES256")
decoded = jwt.decode(token, open("ec-public.pem").read(), algorithms=["ES256"])
```

## fa-clipboard-list 声明（iss/sub/aud/exp）

```json
{
  "iss": "https://auth.example.com",
  "sub": "user-123",
  "aud": "https://api.example.com",
  "exp": 1735689600,
  "iat": 1735603200,
  "nbf": 1735603200,
  "jti": "unique-token-id",
  "scope": "read write",
  "roles": ["admin", "editor"]
}
```

```
iss  = Issuer（签发者）
sub  = Subject（主题，用户 ID）
aud  = Audience（受众，接收方）
exp  = Expiration（过期时间，Unix 时间戳）
iat  = Issued At（签发时间）
nbf  = Not Before（生效时间）
jti  = JWT ID（唯一标识）
```

## fa-circle-check Token 验证

```python
import jwt, time

def validate_token(token, secret, issuer=None, audience=None):
    try:
        decoded = jwt.decode(
            token,
            secret,
            algorithms=["HS256"],
            issuer=issuer,
            audience=audience,
        )
        return decoded
    except jwt.ExpiredSignatureError:
        raise Exception("Token expired")
    except jwt.InvalidTokenError as e:
        raise Exception(f"Invalid token: {e}")
```

```go
token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
    if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected method: %v", t.Header["alg"])
    }
    return []byte("secret"), nil
})
if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    fmt.Println(claims["sub"])
}
```

## fa-rotate 刷新令牌

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiJ9...",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

```json
{
  "sub": "user-123",
  "type": "refresh",
  "exp": 1735862400,
  "jti": "refresh-abc-123"
}
```

```json
POST /auth/refresh
{
  "grant_type": "refresh_token",
  "refresh_token": "eyJhbGciOiJIUzI1NiJ9..."
}
```

## fa-terminal jwt-cli / jwt.io

```bash
jwt decode eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjMifQ.xxx
jwt encode --secret mysecret --sub 123 --exp +1h
jwt decode --json eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjMifQ.xxx

cut -d. -f1 <<< "$TOKEN" | base64 -d 2>/dev/null
cut -d. -f2 <<< "$TOKEN" | base64 -d 2>/dev/null
```

在线解码：https://jwt.io

## fa-golang Go jwt-go

```go
import "github.com/golang-jwt/jwt/v5"

claims := jwt.MapClaims{
    "sub": "user-123",
    "exp": time.Now().Add(time.Hour).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, _ := token.SignedString([]byte("secret"))

parsed, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
    return []byte("secret"), nil
})
```

## fa-node-js Node jsonwebtoken

```javascript
const jwt = require("jsonwebtoken");

const token = jwt.sign(
  { sub: "user-123", role: "admin" },
  process.env.JWT_SECRET,
  { algorithm: "HS256", expiresIn: "1h", issuer: "myapp" }
);

try {
  const decoded = jwt.verify(token, process.env.JWT_SECRET, {
    issuer: "myapp",
    algorithms: ["HS256"],
  });
} catch (err) {
  console.error(err.message);
}

jwt.decode(token, { complete: true });
```

## fa-python Python PyJWT

```python
import jwt, datetime

payload = {
    "sub": "user-123",
    "exp": datetime.datetime.now(datetime.timezone.utc) + datetime.timedelta(hours=1),
    "iat": datetime.datetime.now(datetime.timezone.utc),
}
token = jwt.encode(payload, "secret", algorithm="HS256")

decoded = jwt.decode(token, "secret", algorithms=["HS256"])

header = jwt.get_unverified_header(token)
```

## fa-shield-heart 安全注意事项

```
始终验证算法 —— 防止 "alg: none" 攻击
不要在 payload 中存储敏感数据 —— 仅 base64 编码，非加密
使用强密钥 —— HS256 至少 256 位
设置短过期时间 —— 访问令牌 15 分钟，刷新令牌 7 天
验证所有声明 —— iss、aud、exp、nbf
使用 HTTPS —— 传输中的令牌必须保护
定期轮换密钥 —— 定期更换签名密钥
优先使用 RS256/ES256 —— 非对称密钥可分离验证和签发
```

## fa-repeat 常见模式

```json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9...

Cookie: access_token=eyJhbGciOiJIUzI1NiJ9...; HttpOnly; Secure; SameSite=Strict
```

```nginx
location /api/ {
    auth_request /auth/verify;
}
location = /auth/verify {
    internal;
    proxy_pass http://auth:8080/verify;
    proxy_set_header X-Token $http_authorization;
}
```

```javascript
function authMiddleware(req, res, next) {
  const auth = req.headers.authorization?.split(" ")[1];
  if (!auth) return res.status(401).json({ error: "No token" });
  try {
    req.user = jwt.verify(auth, process.env.JWT_SECRET);
    next();
  } catch {
    res.status(401).json({ error: "Invalid token" });
  }
}
```
