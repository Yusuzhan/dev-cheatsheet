---
title: curl & wget
icon: fa-download
primary: "#6C5CE7"
lang: bash
locale: zhs
---

## fa-globe 基本请求

```bash
curl https://example.com              # 获取 URL 内容，输出到终端
curl -o page.html https://example.com # 保存到指定文件
curl -O https://example.com/file.zip  # 以远程文件名保存
curl -s https://example.com           # 静默模式（不显示进度）
curl -i https://api.example.com       # 包含响应头
curl -I https://example.com           # HEAD 请求（仅获取头信息）

wget https://example.com/file.zip     # 下载文件
wget -q https://example.com/file.zip  # 安静模式
wget -O output.zip https://example.com/file.zip  # 自定义文件名保存
wget -P /tmp/downloads https://example.com/file.zip  # 保存到指定目录
```

## fa-paper-plane HTTP 方法

```bash
curl -X POST https://api.example.com/users     # POST 请求
curl -X PUT https://api.example.com/users/1     # PUT 请求
curl -X DELETE https://api.example.com/users/1  # DELETE 请求
curl -X PATCH https://api.example.com/users/1   # PATCH 请求

# POST 表单数据
curl -d "name=Alice&age=25" https://api.example.com/users

# POST JSON 数据
curl -d '{"name":"Alice","age":25}' \
     -H "Content-Type: application/json" \
     https://api.example.com/users
```

## fa-key 请求头与认证

```bash
curl -H "Authorization: Bearer TOKEN" https://api.example.com  # Bearer 令牌
curl -H "X-Custom-Header: value" https://api.example.com       # 自定义头
curl -H "Accept: application/json" https://api.example.com      # 接受 JSON
curl -u user:pass https://api.example.com                       # 基本认证
curl -u user https://api.example.com                            # 交互式输入密码

wget --header="Authorization: Bearer TOKEN" https://api.example.com
wget --user=user --password=pass https://api.example.com        # 基本认证
wget --header="Content-Type: application/json" https://api.example.com
```

## fa-file-arrow-up 文件上传

```bash
curl -F "file=@photo.jpg" https://api.example.com/upload         # 上传文件
curl -F "file=@photo.jpg;type=image/jpeg" https://api.example.com  # 指定 MIME 类型
curl -F "file=@/path/to/file" -F "name=Alice" https://api.example.com  # 文件 + 字段

# 多部分表单
curl -F "document=@report.pdf" \
     -F "title=Annual Report" \
     -F "category=finance" \
     https://api.example.com/upload
```

## fa-cookie Cookie 处理

```bash
curl -c cookies.txt https://example.com/login     # 保存 cookie 到文件
curl -b cookies.txt https://example.com/dashboard  # 发送已保存的 cookie
curl -b "session=abc123" https://example.com       # 直接发送 cookie
curl -c cookies.txt -b cookies.txt https://example.com  # 同时保存和发送

wget --save-cookies cookies.txt https://example.com/login
wget --load-cookies cookies.txt https://example.com/dashboard
```

## fa-arrows-rotate 重定向与重试

```bash
curl -L https://example.com               # 跟随重定向
curl -L --max-redirs 5 https://example.com # 最大重定向次数
curl --retry 3 https://example.com         # 失败重试 3 次
curl --retry-delay 2 https://example.com   # 重试间隔 2 秒
curl --connect-timeout 10 https://example.com  # 连接超时（秒）
curl --max-time 30 https://example.com     # 总超时（秒）

wget --max-redirect=5 https://example.com  # 最大重定向次数
wget --tries=3 https://example.com         # 重试 3 次
wget --waitretry=5 https://example.com     # 重试间隔 5 秒
wget --timeout=30 https://example.com      # 超时（秒）
```

## fa-download 下载与断点续传

```bash
curl -C - -O https://example.com/large-file.zip  # 断点续传
curl -# -O https://example.com/file.zip           # 显示进度条

wget -c https://example.com/large-file.zip        # 断点续传
wget -r -np -k https://example.com/docs/          # 递归下载网站
wget --mirror https://example.com/                # 完整镜像网站
wget --limit-rate=1m https://example.com/file.zip # 限制下载速度
wget -i urls.txt                                  # 从文件读取 URL 列表下载
```

## fa-shield SSL 与代理

```bash
curl -k https://example.com                # 跳过 SSL 验证
curl --cacert ca.pem https://example.com   # 使用自定义 CA 证书
curl --cert client.pem --key key.pem https://example.com  # 客户端证书
curl -x http://proxy:8080 https://example.com             # HTTP 代理
curl -x socks5://proxy:1080 https://example.com          # SOCKS5 代理

wget --no-check-certificate https://example.com  # 跳过 SSL 验证
wget --ca-certificate=ca.pem https://example.com
wget -e use_proxy=yes http_proxy=http://proxy:8080 https://example.com
```

## fa-file-code 响应处理

```bash
curl -w "\nHTTP Code: %{http_code}\nTime: %{time_total}s\n" \
     -o /dev/null -s https://example.com       # 显示状态码和耗时

curl -D headers.txt https://example.com         # 导出响应头到文件
curl -o response.json -w "%{http_code}" https://api.example.com  # 保存响应体 + 状态码

curl -s https://api.example.com | jq '.'        # 格式化 JSON（需 jq）
curl -s https://api.example.com | python3 -m json.tool  # 格式化 JSON

wget -S https://example.com                     # 显示服务器响应头
wget --server-response https://example.com      # 显示响应头
```

## fa-lightbulb 实用技巧

```bash
# 检查 URL 是否可访问
curl -s -o /dev/null -w "%{http_code}" https://example.com

# 带认证和断点续传下载
curl -L -C - -u user:pass -o file.zip https://example.com/file.zip

# POST JSON 并格式化响应
curl -s -X POST -H "Content-Type: application/json" \
     -d '{"key":"value"}' https://api.example.com | jq '.'

# 调试 HTTP 请求
curl -v https://example.com                    # 详细输出
curl --trace-ascii debug.log https://example.com  # 完整追踪写入文件

# 批量下载多个文件
curl -O https://example.com/a.txt -O https://example.com/b.txt

# 使用 wget 镜像网站
wget --mirror --convert-links --adjust-extension --page-requisites \
     --no-parent https://example.com/docs/
```
