---
title: Nginx
icon: fa-server
primary: "#009639"
lang: nginx
locale: zhs
---

## fa-server 基本 Server 块

```nginx
server {
    listen 80;
    server_name example.com www.example.com;
    root /var/www/html;
    index index.html index.htm;

    location / {
        try_files $uri $uri/ =404;
    }
}
```

## fa-arrow-right 静态文件与 Location

```nginx
# 精确匹配
location = /favicon.ico {
    log_not_found off;              # 不记录 404 日志
    return 204;                     # 返回无内容
}

# 前缀匹配（优先级：精确 > 优先前缀 > 正则 > 前缀）
location ^~ /static/ {
    expires 30d;                    # 缓存 30 天
    add_header Cache-Control "public, immutable";
}

# 正则匹配（不区分大小写）
location ~* \.(jpg|jpeg|png|gif|ico|css|js)$ {
    expires 7d;
    gzip_static on;                 # 使用预压缩的 .gz 文件
}

# 禁止访问隐藏文件
location ~ /\. {
    deny all;
}
```

## fa-rotate 反向代理

```nginx
server {
    listen 80;
    server_name api.example.com;

    location / {
        proxy_pass http://127.0.0.1:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # 代理 WebSocket
    location /ws/ {
        proxy_pass http://127.0.0.1:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
```

## fa-shield-halved SSL / HTTPS

```nginx
server {
    listen 443 ssl http2;
    server_name example.com;

    ssl_certificate     /etc/ssl/certs/example.com.crt;
    ssl_certificate_key /etc/ssl/private/example.com.key;

    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;

    # HTTP 重定向到 HTTPS（放在单独的 server 块中）
}

server {
    listen 80;
    server_name example.com;
    return 301 https://$host$request_uri;
}
```

## fa-scale-balanced 负载均衡

```nginx
upstream backend {
    # 负载均衡策略
    # least_conn;                    # 最少连接数
    # ip_hash;                       # 按 IP 保持会话
    # hash $request_uri consistent;  # 一致性哈希

    server 10.0.0.1:3000 weight=3;   # weight=3 分配更多流量
    server 10.0.0.2:3000;
    server 10.0.0.3:3000 backup;     # 仅在其他节点不可用时启用
}

server {
    listen 80;
    location / {
        proxy_pass http://backend;
        proxy_next_upstream error timeout http_502;
    }
}
```

## fa-bolt 缓存与压缩

```nginx
# gzip 压缩
gzip on;
gzip_vary on;
gzip_min_length 1024;               # 仅压缩大于 1KB 的响应
gzip_types text/plain text/css application/json application/javascript
           text/xml application/xml application/xml+rss text/javascript;

# 代理缓存
proxy_cache_path /var/cache/nginx levels=1:2 keys_zone=api_cache:10m
                 max_size=1g inactive=60m use_temp_path=off;

server {
    location /api/ {
        proxy_pass http://backend;
        proxy_cache api_cache;
        proxy_cache_valid 200 10m;              # 200 响应缓存 10 分钟
        proxy_cache_valid 404 1m;
        proxy_cache_use_stale error timeout updating;
        add_header X-Cache-Status $upstream_cache_status;
    }
}
```

## fa-ban 限流与安全

```nginx
# 请求限流
limit_req_zone $binary_remote_addr zone=api_limit:10m rate=10r/s;

server {
    location /api/ {
        limit_req zone=api_limit burst=20 nodelay;
        proxy_pass http://backend;
    }
}

# 连接数限制
limit_conn_zone $binary_remote_addr zone=conn_limit:10m;

# 安全响应头
add_header X-Frame-Options "SAMEORIGIN" always;
add_header X-Content-Type-Options "nosniff" always;
add_header X-XSS-Protection "1; mode=block" always;
add_header Referrer-Policy "strict-origin-when-cross-origin" always;
add_header Content-Security-Policy "default-src 'self'" always;

# IP 白名单/黑名单
location /admin/ {
    allow 192.168.1.0/24;
    allow 10.0.0.1;
    deny all;
    proxy_pass http://backend;
}
```

## fa-compass 重写与重定向

```nginx
# 永久重定向 (301)
return 301 https://$host$request_uri;

# 临时重定向 (302)
return 302 /maintenance.html;

# URL 重写
rewrite ^/old-page/(.*)$ /new-page/$1 permanent;
rewrite ^/blog/(\d{4})/(\d{2})/(.*)$ /articles/$3?year=$1&month=$2 last;

# SPA 应用 try_files
location / {
    try_files $uri $uri/ /index.html;
}

# 强制尾部斜杠
rewrite ^([^.]*[^/])$ $1/ permanent;
```

## fa-file-code PHP-FPM

```nginx
server {
    listen 80;
    server_name example.com;
    root /var/www/laravel/public;
    index index.php;

    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location ~ \.php$ {
        fastcgi_pass unix:/run/php/php8.2-fpm.sock;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
        fastcgi_buffer_size 128k;
        fastcgi_buffers 4 256k;
        fastcgi_read_timeout 300;
    }

    location ~ /\.ht {
        deny all;
    }
}
```

## fa-gears 日志与调试

```nginx
# 访问日志格式
log_format main '$remote_addr - $remote_user [$time_local] '
                '"$request" $status $body_bytes_sent '
                '"$http_referer" "$http_user_agent" '
                '$request_time';

access_log /var/log/nginx/access.log main;
error_log  /var/log/nginx/error.log warn;   # debug, info, notice, warn, error, crit

# 条件日志（跳过健康检查）
map $request_uri $loggable {
    ~/health$ 0;
    default 1;
}
access_log /var/log/nginx/access.log main if=$loggable;
```

## fa-terminal 命令行操作

```bash
nginx -t                  # 测试配置文件语法
nginx -T                  # 测试并打印完整配置
nginx -s reload           # 重载配置（不中断服务）
nginx -s stop             # 快速停止
nginx -s quit             # 优雅停止（处理完当前请求）
nginx -c /etc/nginx/nginx.conf  # 指定配置文件
```

## fa-lightbulb 性能优化

```nginx
worker_processes auto;                # 每个 CPU 核心一个 worker
worker_rlimit_nofile 65535;          # 文件描述符上限

events {
    worker_connections 4096;          # 每个 worker 的连接数
    multi_accept on;                  # 一次性接受所有新连接
    use epoll;                        # Linux 使用 epoll
}

http {
    sendfile on;                      # 内核级文件传输
    tcp_nopush on;                    # 将头信息合并为一个包发送
    tcp_nodelay on;                   # 禁用 Nagle 算法
    keepalive_timeout 65;             # 长连接超时时间
    client_max_body_size 20m;         # 最大上传大小
    server_tokens off;                # 隐藏 Nginx 版本号
    client_body_buffer_size 16k;      # 请求体缓冲区
    open_file_cache max=1000 inactive=20s;
    open_file_cache_valid 30s;
}
```
