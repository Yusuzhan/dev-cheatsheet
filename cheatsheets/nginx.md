---
title: Nginx
icon: fa-server
primary: "#009639"
lang: nginx
---

## fa-server Basic Server Block

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

## fa-arrow-right Static Files & Location

```nginx
# exact match
location = /favicon.ico {
    log_not_found off;              # don't log 404
    return 204;                     # return no content
}

# prefix match (priority: exact > preferential prefix > regex > prefix)
location ^~ /static/ {
    expires 30d;                    # cache for 30 days
    add_header Cache-Control "public, immutable";
}

# regex match (case-sensitive)
location ~* \.(jpg|jpeg|png|gif|ico|css|js)$ {
    expires 7d;
    gzip_static on;                 # serve pre-compressed .gz
}

# deny access to hidden files
location ~ /\. {
    deny all;
}
```

## fa-rotate Reverse Proxy

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

    # proxy WebSocket
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

    # HTTP to HTTPS redirect
    # (place in a separate server block)
}

server {
    listen 80;
    server_name example.com;
    return 301 https://$host$request_uri;
}
```

## fa-scale-balanced Load Balancing

```nginx
upstream backend {
    # load balancing methods
    # least_conn;                    # least connections
    # ip_hash;                       # session persistence by IP
    # hash $request_uri consistent;  # consistent hashing

    server 10.0.0.1:3000 weight=3;   # weight=3 gets more traffic
    server 10.0.0.2:3000;
    server 10.0.0.3:3000 backup;     # only used when others are down
}

server {
    listen 80;
    location / {
        proxy_pass http://backend;
        proxy_next_upstream error timeout http_502;
    }
}
```

## fa-bolt Caching & Compression

```nginx
# gzip compression
gzip on;
gzip_vary on;
gzip_min_length 1024;               # only compress responses > 1KB
gzip_types text/plain text/css application/json application/javascript
           text/xml application/xml application/xml+rss text/javascript;

# proxy cache
proxy_cache_path /var/cache/nginx levels=1:2 keys_zone=api_cache:10m
                 max_size=1g inactive=60m use_temp_path=off;

server {
    location /api/ {
        proxy_pass http://backend;
        proxy_cache api_cache;
        proxy_cache_valid 200 10m;              # cache 200 responses for 10min
        proxy_cache_valid 404 1m;
        proxy_cache_use_stale error timeout updating;
        add_header X-Cache-Status $upstream_cache_status;
    }
}
```

## fa-ban Rate Limiting & Security

```nginx
# rate limiting
limit_req_zone $binary_remote_addr zone=api_limit:10m rate=10r/s;

server {
    location /api/ {
        limit_req zone=api_limit burst=20 nodelay;
        proxy_pass http://backend;
    }
}

# connection limiting
limit_conn_zone $binary_remote_addr zone=conn_limit:10m;

# security headers
add_header X-Frame-Options "SAMEORIGIN" always;
add_header X-Content-Type-Options "nosniff" always;
add_header X-XSS-Protection "1; mode=block" always;
add_header Referrer-Policy "strict-origin-when-cross-origin" always;
add_header Content-Security-Policy "default-src 'self'" always;

# allow/deny by IP
location /admin/ {
    allow 192.168.1.0/24;
    allow 10.0.0.1;
    deny all;
    proxy_pass http://backend;
}
```

## fa-compass Rewrites & Redirects

```nginx
# permanent redirect (301)
return 301 https://$host$request_uri;

# temporary redirect (302)
return 302 /maintenance.html;

# rewrite URL
rewrite ^/old-page/(.*)$ /new-page/$1 permanent;
rewrite ^/blog/(\d{4})/(\d{2})/(.*)$ /articles/$3?year=$1&month=$2 last;

# try_files for SPA
location / {
    try_files $uri $uri/ /index.html;
}

# force trailing slash
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

## fa-gears Logging & Debug

```nginx
# access log format
log_format main '$remote_addr - $remote_user [$time_local] '
                '"$request" $status $body_bytes_sent '
                '"$http_referer" "$http_user_agent" '
                '$request_time';

access_log /var/log/nginx/access.log main;
error_log  /var/log/nginx/error.log warn;   # debug, info, notice, warn, error, crit

# conditional logging (skip health checks)
map $request_uri $loggable {
    ~/health$ 0;
    default 1;
}
access_log /var/log/nginx/access.log main if=$loggable;
```

## fa-terminal Command Line

```bash
nginx -t                  # test configuration syntax
nginx -T                  # test and print full config
nginx -s reload           # reload config without downtime
nginx -s stop             # fast shutdown
nginx -s quit             # graceful shutdown (finish requests)
nginx -c /etc/nginx/nginx.conf  # use custom config file
```

## fa-lightbulb Performance Tips

```nginx
worker_processes auto;                # one worker per CPU core
worker_rlimit_nofile 65535;          # file descriptor limit

events {
    worker_connections 4096;          # connections per worker
    multi_accept on;                  # accept all new connections at once
    use epoll;                        # Linux: use epoll
}

http {
    sendfile on;                      # kernel-level file transfer
    tcp_nopush on;                    # send headers in one packet
    tcp_nodelay on;                   # disable Nagle's algorithm
    keepalive_timeout 65;             # keep-alive timeout
    client_max_body_size 20m;         # max upload size
    server_tokens off;                # hide Nginx version
    client_body_buffer_size 16k;      # request body buffer
    open_file_cache max=1000 inactive=20s;
    open_file_cache_valid 30s;
}
```
