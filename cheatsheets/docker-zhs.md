---
title: Docker
icon: fa-docker
primary: "#2496ED"
lang: bash
locale: zhs
---

## fa-play 容器生命周期

```bash
docker create --name web nginx:latest
docker start web
docker stop web
docker restart web
docker pause web
docker unpause web
docker rm web
docker rm -f web                  # 强制删除运行中的容器
```

## fa-terminal 运行容器

```bash
docker run nginx
docker run -d --name web -p 8080:80 nginx
docker run -it ubuntu:22.04 bash   # 交互式进入
docker run --rm alpine echo "hello"  # 运行后自动删除
docker run -e MYSQL_ROOT_PASSWORD=secret mysql
docker run -v /host/data:/container/data nginx
docker run --network mynet --name app myimage
docker run --cpus="1.5" --memory="512m" nginx
```

## fa-list 查看容器

```bash
docker ps                         # 运行中的容器
docker ps -a                      # 所有容器
docker ps -q                      # 仅显示 ID
docker ps --filter "status=running"
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
docker stats                      # 实时资源占用
docker top web                    # 容器内进程
```

## fa-box 执行与日志

```bash
docker exec -it web bash          # 进入容器
docker exec web cat /etc/hosts
docker exec -u root web apt-get update
docker logs web
docker logs -f web                # 实时跟踪
docker logs --tail 100 web
docker logs --since 2h web
docker cp web:/app/config.yml ./config.yml
docker cp ./data web:/app/data
```

## fa-layer-group 镜像管理

```bash
docker images
docker pull node:20-alpine
docker push myrepo/myapp:latest
docker build -t myapp:1.0 .
docker build -f Dockerfile.prod -t myapp:prod .
docker rmi nginx:latest
docker image prune                # 清理悬空镜像
docker image prune -a             # 清理所有未使用镜像
docker tag myapp:1.0 myrepo/myapp:latest
```

## fa-hard-drive 数据卷

```bash
docker volume create mydata
docker volume ls
docker volume inspect mydata
docker volume rm mydata
docker volume prune

docker run -v mydata:/var/lib/mysql mysql    # 命名卷
docker run -v $(pwd):/app node:20-alpine     # 绑定挂载
docker run --tmpfs /tmp nginx                # 内存卷
docker run --mount type=bind,src=/host,dst=/container nginx
```

## fa-network-wired 网络管理

```bash
docker network create mynet
docker network ls
docker network inspect mynet
docker network rm mynet
docker network prune

docker run --network mynet --name api myimage
docker network connect mynet web      # 将运行中容器加入网络
docker network disconnect mynet web
```

## fa-copy Docker Compose

```yaml
version: "3.9"
services:
  web:
    build: .
    ports:
      - "8080:80"
    volumes:
      - .:/app
    environment:
      - NODE_ENV=development
    depends_on:
      - db
  db:
    image: postgres:16
    volumes:
      - pgdata:/var/lib/postgresql/data
    environment:
      POSTGRES_PASSWORD: secret

volumes:
  pgdata:
```

```bash
docker compose up -d
docker compose down
docker compose down -v              # 同时删除数据卷
docker compose logs -f
docker compose ps
docker compose exec web bash
docker compose build
docker compose restart web
```

## fa-file-code Dockerfile

```dockerfile
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:20-alpine
WORKDIR /app
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
EXPOSE 3000
HEALTHCHECK --interval=30s CMD wget -qO- http://localhost:3000/health
USER node
CMD ["node", "dist/main.js"]
```

## fa-broom 清理

```bash
docker system df                   # 查看磁盘占用
docker system prune                # 清理未使用资源
docker system prune -a --volumes   # 彻底清理

docker container prune             # 清理已停止容器
docker image prune -a
docker volume prune
docker network prune

docker rmi $(docker images -q -f dangling=true)
docker rm $(docker ps -aq -f status=exited)
```

## fa-magnifying-glass 检查与调试

```bash
docker inspect web
docker inspect -f '{{.NetworkSettings.IPAddress}}' web
docker port web
docker diff web                    # 查看文件变更
docker history nginx:latest        # 查看镜像层
docker logs --since 30m -f web 2>&1 | grep ERROR
docker exec -it web sh -c "apt-get update && apt-get install -y curl"
```

## fa-lightbulb 实用技巧

```bash
docker run --init node:20 node app.js       # PID 1 信号处理
docker run --read-only --tmpfs /tmp nginx    # 只读文件系统
docker build --no-cache -t myapp .           # 无缓存构建
docker build --build-arg VERSION=1.0 -t myapp .
docker compose up -d --scale worker=3        # 扩展服务
docker save myapp:1.0 | gzip > myapp.tar.gz  # 导出镜像
docker load < myapp.tar.gz                   # 导入镜像
docker system info
```
