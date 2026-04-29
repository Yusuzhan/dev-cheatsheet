---
title: OAuth 2.0 / OIDC
icon: fa-shield-halved
primary: "#3B82F6"
lang: http
---

## fa-sitemap OAuth Roles & Flow

```text
Roles:
Resource Owner     User who grants access
Client             App requesting access
Authorization Server  Issues tokens
Resource Server    API serving protected data

Grant Types:
authorization_code   Server-side apps (most secure)
implicit             Legacy SPA (deprecated)
client_credentials   Service-to-service
password             Legacy (deprecated)
refresh_token        Obtain new access token
device_code          IoT / constrained devices
```

```text
+----------+     +-----------+     +------------------+
|  Client  |---->| Resource  |---->| Authorization    |
|  (App)   |     | Owner     |     | Server           |
|          |<----| (User)    |<----| (auth provider)  |
+----------+     +-----------+     +------------------+
     |                                     |
     |---------- Token Request ----------->|
     |<--------- Access Token -------------|
     |                                     |
     |---------- API Request ------------->| Resource Server
     |<---------- Protected Data ---------| (separate or same)
```

## fa-arrow-right Authorization Code Flow

```http
# Step 1: Redirect user to authorize
GET /authorize?
  response_type=code
  &client_id=myapp
  &redirect_uri=https://app.example.com/callback
  &scope=read write
  &state=random_csrf_token HTTP/1.1
Host: auth.example.com
```

```http
# Step 2: User approves, callback with code
HTTP/1.1 302 Found
Location: https://app.example.com/callback?
  code=SplxlOBeZQQYbYS6WxSbIA
  &state=random_csrf_token
```

```http
# Step 3: Exchange code for token
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

grant_type=authorization_code
&code=SplxlOBeZQQYbYS6WxSbIA
&redirect_uri=https://app.example.com/callback
```

```http
# Step 4: Token response
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

## fa-lock PKCE Extension

```text
PKCE (Proof Key for Code Exchange) - required for public clients (SPAs, mobile)
```

```bash
code_verifier = random_string_43_to_128_chars
code_challenge = base64url(sha256(code_verifier))
# e.g., "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"
```

```http
# Authorization request with PKCE
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
# Token exchange with code_verifier
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=SplxlOBeZQQYbYS6WxSbIA
&client_id=myapp
&redirect_uri=https://app.example.com/callback
&code_verifier=dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk
```

## fa-gears Client Credentials Flow

```http
# Service-to-service (no user involved)
POST /token HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

grant_type=client_credentials
&scope=read write
```

```http
# Response (no refresh_token for client credentials)
HTTP/1.1 200 OK
Content-Type: application/json

{
  "access_token": "eyJhbGciOiJSUzI1NiJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "scope": "read write"
}
```

## fa-user Resource Owner Password

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
DEPRECATED: Use only for highly trusted first-party clients.
Migrate to authorization_code + PKCE instead.
```

## fa-arrows-rotate Refresh Token

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
Best practices:
- Rotation: issue new refresh_token on each use
- Detection: revoke if old refresh_token is reused (compromise)
- Expiry: set absolute lifetime + idle timeout
- Revocation: provide /revoke endpoint
```

## fa-circle-check Token Introspection

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
# Revocation
POST /revoke HTTP/1.1
Host: auth.example.com
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

token=eyJhbGciOiJSUzI1NiJ9...
&token_type_hint=access_token
```

## fa-id-card OIDC ID Token

```text
OpenID Connect = OAuth 2.0 + Identity Layer

ID Token: JWT containing user identity claims
Verified by checking: signature, issuer, audience, expiry, nonce
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
ID Token (JWT decoded payload):
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

## fa-compass OIDC Discovery

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

## fa-sliders Scopes & Claims

```text
OAuth 2.0 Scopes:
openid           Required for OIDC
profile          name, family_name, given_name, picture
email            email, email_verified
address          address
phone            phone_number, phone_number_verified
offline_access   Request refresh_token

Custom scopes:
read:posts       Read user posts
write:posts      Create/edit posts
admin            Administrative access
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

## fa-barcode JWT Structure

```text
JWT = Header.Payload.Signature

Header (base64url):
{ "alg": "RS256", "typ": "JWT", "kid": "key-2025-01" }

Payload (base64url):
{ "sub": "user123", "iss": "https://auth.example.com", "aud": "myapp",
  "exp": 1735689600, "iat": 1735686000, "scope": "read write" }

Signature:
RS256(base64url(header) + "." + base64url(payload), private_key)
```

```bash
# Decode JWT (no verification)
echo "eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1c2VyMTIzIn0.sig" | cut -d. -f2 | base64 -d 2>/dev/null

# Using jq
echo "eyJ..." | awk -F. '{printf "%s", $2}' | basenc --base64url -d | jq .
```

```text
Registered Claims:
iss  Issuer           aud  Audience
sub  Subject          exp  Expiration
iat  Issued at        nbf  Not before
jti  JWT ID           scope Permissions
nonce  ID token only  azp  Authorized party
```

## fa-building Common Providers (Google/GitHub)

```text
Google:
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
GitHub:
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

## fa-shield Security Best Practices

```text
Authorization:
- Always use state parameter (CSRF protection)
- Use PKCE for all public clients
- Validate redirect_uri exactly (no wildcards)
- Use authorization_code flow (avoid implicit grant)

Tokens:
- Short access_token lifetime (5-15 minutes)
- Rotate refresh tokens on use
- Validate JWT: signature, iss, aud, exp
- Store tokens securely (not localStorage in browser)

Transport:
- Always use HTTPS
- Use TLS 1.2+
- Pin certificates for high-security apps
- HSTS headers on auth endpoints

Client:
- Never expose client_secret in public clients
- Use private_key_jwt for confidential clients
- Register all redirect_uris explicitly
- Use minimal scopes (principle of least privilege)
```

## fa-box Token Storage

```text
Browser (SPA):
✓ In-memory (JavaScript variable)
✓ Service worker cache
✓ httpOnly cookie (if backend-for-frontend)
✗ localStorage (XSS accessible)
✗ sessionStorage (XSS accessible)

Mobile:
✓ Keychain (iOS) / Keystore (Android)
✓ Secure enclave / hardware-backed
✗ AsyncStorage / SharedPreferences

Server:
✓ Encrypted database with access controls
✓ Hash refresh_tokens before storage
✓ Environment variables for client secrets
✗ Plain text in database
✗ Committed to version control
```

```text
Token Lifecycle:
Access Token:     5-15 minutes, used in API calls
Refresh Token:    Days-weeks, stored securely, rotated
ID Token:         Validated once, extract claims, discard
Authorization Code: Single use, expires in ~10 minutes
```
