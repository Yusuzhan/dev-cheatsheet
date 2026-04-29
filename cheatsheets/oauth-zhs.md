---
title: OAuth 2.0 / OIDC
icon: fa-shield-halved
primary: "#3B82F6"
lang: http
locale: zhs
---

## fa-sitemap OAuth 角色与流程

```text
角色：
Resource Owner（资源所有者）      授权访问的用户
Client（客户端）                  请求访问的应用
Authorization Server（授权服务器） 签发令牌
Resource Server（资源服务器）      提供受保护数据的 API

授权类型：
authorization_code   服务端应用（最安全）
implicit             旧版 SPA（已废弃）
client_credentials   服务间通信
password             旧版（已废弃）
refresh_token        获取新的访问令牌
device_code          IoT / 受限设备
```

```text
+----------+     +-----------+     +------------------+
|  客户端  |---->| 资源      |---->| 授权             |
|  (应用)  |     | 所有者    |     | 服务器           |
|          |<----| (用户)    |<----| (认证提供者)     |
+----------+     +-----------+     +------------------+
     |                                     |
     |---------- 令牌请求 ---------------->|
     |<--------- 访问令牌 ----------------|
     |                                     |
     |---------- API 请求 ---------------->| 资源服务器
     |<---------- 受保护数据 -------------| （独立或相同）
```

## fa-arrow-right 授权码流程

```http
# 步骤 1：重定向用户到授权页
GET /authorize?
  response_type=code
  &client_id=myapp
  &redirect_uri=https://app.example.com/callback
  &scope=read write
  &state=random_csrf_token HTTP/1.1
Host: auth.example.com
```

```http
# 步骤 2：用户同意，回调携带授权码
HTTP/1.1 302 Found
Location: https://app.example.com/callback?
  code=SplxlOBeZQQYbYS6WxSbIA
  &state=random_csrf_token
```

```http
# 步骤 3：用授权码换取令牌
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

grant_type=authorization_code
&code=SplxlOBeZQQYbYS6WxSbIA
&redirect_uri=https://app.example.com/callback
```

```http
# 步骤 4：令牌响应
HTTP/1.1 200 OK
Content-Type: application/json

{
  "access_token": "eyJhbGciOiJSUzI1NiJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "tGzv3JOkF0XG5Qx2TlKWIA",
  "scope": "read write"
}
```

## fa-lock PKCE 扩展

```text
PKCE（代码交换证明密钥）— 公共客户端（SPA、移动端）必须使用
```

```bash
code_verifier = 随机字符串_43到128个字符
code_challenge = base64url(sha256(code_verifier))
# 例如："dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"
```

```http
# 带 PKCE 的授权请求
GET /authorize?
  response_type=code
  &client_id=myapp
  &redirect_uri=https://app.example.com/callback
  &scope=read
  &state=abc123
  &code_challenge=E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM
  &code_challenge_method=S256 HTTP/1.1
```

```http
# 携带 code_verifier 的令牌交换
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=SplxlOBeZQQYbYS6WxSbIA
&client_id=myapp
&redirect_uri=https://app.example.com/callback
&code_verifier=dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk
```

## fa-gears 客户端凭证流程

```http
# 服务间通信（不涉及用户）
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

grant_type=client_credentials
&scope=read write
```

```http
# 响应（客户端凭证无 refresh_token）
HTTP/1.1 200 OK
Content-Type: application/json

{
  "access_token": "eyJhbGciOiJSUzI1NiJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "scope": "read write"
}
```

## fa-user 资源所有者密码

```http
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

grant_type=password
&username=user@example.com
&password=secret123
&scope=read
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "access_token": "eyJhbGciOiJSUzI1NiJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "tGzv3JOkF0XG5Qx2TlKWIA",
  "scope": "read"
}
```

```text
已废弃：仅用于高度信任的第一方客户端。
建议迁移到 authorization_code + PKCE。
```

## fa-arrows-rotate 刷新令牌

```http
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

grant_type=refresh_token
&refresh_token=tGzv3JOkF0XG5Qx2TlKWIA
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "access_token": "new_eyJhbGciOiJSUzI1NiJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "new_tGzv3JOkF0XG5Qx2TlKWIA",
  "scope": "read write"
}
```

```text
最佳实践：
- 轮换：每次使用时签发新的 refresh_token
- 检测：旧 refresh_token 被重用时撤销（泄露检测）
- 过期：设置绝对有效期 + 空闲超时
- 撤销：提供 /revoke 端点
```

## fa-circle-check 令牌内省

```http
POST /introspect HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

token=eyJhbGciOiJSUzI1NiJ9...
&token_type_hint=access_token
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "active": true,
  "scope": "read write",
  "client_id": "myapp",
  "sub": "user123",
  "exp": 1735689600,
  "iat": 1735686000,
  "token_type": "Bearer"
}
```

```http
# 令牌撤销
POST /revoke HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

token=eyJhbGciOiJSUzI1NiJ9...
&token_type_hint=access_token
```

## fa-id-card OIDC ID 令牌

```text
OpenID Connect = OAuth 2.0 + 身份层

ID Token：包含用户身份声明的 JWT
验证项：签名、签发者、受众、过期时间、nonce
```

```http
GET /authorize?
  response_type=code
  &client_id=myapp
  &redirect_uri=https://app.example.com/callback
  &scope=openid profile email
  &state=abc123
  &nonce=random_nonce HTTP/1.1
```

```text
ID Token（JWT 解码后载荷）：
{
  "iss": "https://auth.example.com",
  "sub": "user123",
  "aud": "myapp",
  "exp": 1735689600,
  "iat": 1735686000,
  "nonce": "random_nonce",
  "name": "Alice",
  "email": "alice@example.com",
  "email_verified": true,
  "picture": "https://auth.example.com/avatar.png"
}
```

## fa-compass OIDC 发现

```http
GET /.well-known/openid-configuration HTTP/1.1
Host: auth.example.com
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "issuer": "https://auth.example.com",
  "authorization_endpoint": "https://auth.example.com/authorize",
  "token_endpoint": "https://auth.example.com/token",
  "userinfo_endpoint": "https://auth.example.com/userinfo",
  "jwks_uri": "https://auth.example.com/.well-known/jwks.json",
  "revocation_endpoint": "https://auth.example.com/revoke",
  "introspection_endpoint": "https://auth.example.com/introspect",
  "end_session_endpoint": "https://auth.example.com/logout",
  "scopes_supported": ["openid", "profile", "email"],
  "response_types_supported": ["code", "id_token"],
  "token_endpoint_auth_methods_supported": ["client_secret_basic", "private_key_jwt"]
}
```

```http
GET /.well-known/jwks.json HTTP/1.1
Host: auth.example.com
```

## fa-sliders 范围与声明

```text
OAuth 2.0 范围：
openid           OIDC 必需
profile          name, family_name, given_name, picture
email            email, email_verified
address          address
phone            phone_number, phone_number_verified
offline_access   请求 refresh_token

自定义范围：
read:posts       读取用户文章
write:posts      创建/编辑文章
admin            管理权限
```

```http
GET /userinfo HTTP/1.1
Host: auth.example.com
Authorization: Bearer eyJhbGciOiJSUzI1NiJ9...
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "sub": "user123",
  "name": "Alice",
  "given_name": "Alice",
  "family_name": "Smith",
  "email": "alice@example.com",
  "email_verified": true,
  "picture": "https://auth.example.com/avatar.png",
  "updated_at": 1735686000
}
```

## fa-barcode JWT 结构

```text
JWT = Header.Payload.Signature

Header（base64url）：
{ "alg": "RS256", "typ": "JWT", "kid": "key-2025-01" }

Payload（base64url）：
{ "sub": "user123", "iss": "https://auth.example.com", "aud": "myapp",
  "exp": 1735689600, "iat": 1735686000, "scope": "read write" }

Signature：
RS256(base64url(header) + "." + base64url(payload), private_key)
```

```bash
# 解码 JWT（不验证）
echo "eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1c2VyMTIzIn0.sig" | cut -d. -f2 | base64 -d 2>/dev/null

# 使用 jq
echo "eyJ..." | awk -F. '{printf "%s", $2}' | basenc --base64url -d | jq .
```

```text
标准声明：
iss  签发者          aud  受众
sub  主体            exp  过期时间
iat  签发时间        nbf  生效时间
jti  JWT ID          scope 权限
nonce  仅 ID Token   azp  授权方
```

## fa-building 常见提供者 (Google/GitHub)

```text
Google：
Auth:     https://accounts.google.com/o/oauth2/v2/auth
Token:    https://oauth2.googleapis.com/token
UserInfo: https://www.googleapis.com/oauth2/v3/userinfo
Scope:    openid profile email
Discovery: https://accounts.google.com/.well-known/openid-configuration
```

```http
GET https://accounts.google.com/o/oauth2/v2/auth?
  response_type=code
  &client_id=xxx.apps.googleusercontent.com
  &redirect_uri=https://app.example.com/callback
  &scope=openid profile email
  &state=abc123
  &nonce=random_nonce
```

```text
GitHub：
Auth:     https://github.com/login/oauth/authorize
Token:    https://github.com/login/oauth/access_token
UserInfo: https://api.github.com/user
Scope:    user:email read:org repo
```

```http
GET https://github.com/login/oauth/authorize?
  client_id=Ov23liabc123
  &redirect_uri=https://app.example.com/callback
  &scope=user:email
  &state=abc123
```

## fa-shield 安全最佳实践

```text
授权：
- 始终使用 state 参数（CSRF 防护）
- 所有公共客户端使用 PKCE
- 精确验证 redirect_uri（禁止通配符）
- 使用 authorization_code 流程（避免 implicit）

令牌：
- 短有效期访问令牌（5-15 分钟）
- 每次使用轮换 refresh_token
- 验证 JWT：签名、iss、aud、exp
- 安全存储令牌（浏览器中不用 localStorage）

传输：
- 始终使用 HTTPS
- 使用 TLS 1.2+
- 高安全应用使用证书固定
- 认证端点启用 HSTS

客户端：
- 公共客户端不得暴露 client_secret
- 机密客户端使用 private_key_jwt
- 显式注册所有 redirect_uri
- 使用最小范围（最小权限原则）
```

## fa-box 令牌存储

```text
浏览器（SPA）：
✓ 内存中（JavaScript 变量）
✓ Service Worker 缓存
✓ httpOnly Cookie（如有 BFF 后端）
✗ localStorage（XSS 可访问）
✗ sessionStorage（XSS 可访问）

移动端：
✓ Keychain (iOS) / Keystore (Android)
✓ 安全 enclave / 硬件安全模块
✗ AsyncStorage / SharedPreferences

服务端：
✓ 加密数据库 + 访问控制
✓ 存储 refresh_token 前先哈希
✓ 环境变量存储 client_secret
✗ 数据库明文存储
✗ 提交到版本控制
```

```text
令牌生命周期：
Access Token：     5-15 分钟，用于 API 调用
Refresh Token：    数天至数周，安全存储，轮换使用
ID Token：         验证一次，提取声明后丢弃
Authorization Code：单次使用，约 10 分钟过期
```
