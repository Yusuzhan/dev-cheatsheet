---
title: Docker
icon: fa-docker
primary: "#2496ED"
lang: bash
locale: zhs
---

## fa-play 容器生命周期

```bash
docker create --name web nginx:latest  # 创建容器但不启动
docker start web                       # 启动已停止的容器
docker stop web                        # 优雅停止 (先 SIGTERM 再 SIGKILL)
docker restart web                     # 停止后重新启动
docker pause web                       # 冻结所有进程 (cgroups freeze)
docker unpause web                     # 解除冻结
docker rm web                          # 删除已停止的容器
docker rm -f web                       # 强制删除运行中的容器
```

## fa-terminal 运行容器

```bash
docker run nginx                       # 前台运行
docker run -d --name web -p 8080:80 nginx  # 后台运行，端口映射
docker run -it ubuntu:22.04 bash       # 交互式终端
docker run --rm alpine echo "hello"    # 退出后自动删除
docker run -e MYSQL_ROOT_PASSWORD=secret mysql  # 设置环境变量
docker run -v /host/data:/container/data nginx   # 绑定挂载卷
docker run --network mynet --name app myimage     # 加入指定网络
docker run --cpus="1.5" --memory="512m" nginx     # 限制资源
```

## fa-list 查看容器

```bash
docker ps                              # 仅显示运行中的容器
docker ps -a                           # 所有容器（含已停止）
docker ps -q                           # 静默模式，仅显示 ID
docker ps --filter "status=running"    # 按状态过滤
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
docker container ls -a                 # docker ps -a 的别名
docker stats                           # 实时资源占用
docker top web                         # 查看容器内进程
```

## fa-box 执行与日志

```bash
docker exec -it web bash               # 进入交互式 shell
docker exec web cat /etc/hosts         # 在容器内执行命令
docker exec -u root web apt-get update # 以指定用户执行
docker logs web                        # 查看全部日志
docker logs -f web                     # 实时跟踪日志
docker logs --tail 100 web             # 最后 100 行
docker logs --since 2h web             # 最近 2 小时的日志
docker cp web:/app/config.yml ./config.yml  # 从容器复制文件
docker cp ./data web:/app/data               # 复制文件到容器
```

## fa-layer-group 镜像管理

```bash
docker images                          # 列出本地镜像
docker pull node:20-alpine             # 从仓库拉取镜像
docker push myrepo/myapp:latest        # 推送到仓库
docker build -t myapp:1.0 .            # 从当前目录 Dockerfile 构建
docker build -f Dockerfile.prod -t myapp:prod .  # 指定 Dockerfile
docker rmi nginx:latest                # 删除镜像
docker image prune                     # 清理悬空镜像
docker image prune -a                  # 清理所有未使用镜像
docker tag myapp:1.0 myrepo/myapp:latest  # 打标签用于推送
```

## fa-hard-drive 数据卷

```bash
docker volume create mydata            # 创建命名卷
docker volume ls                       # 列出所有卷
docker volume inspect mydata           # 查看卷详情（挂载路径等）
docker volume rm mydata                # 删除卷
docker volume prune                    # 清理未使用的卷

docker run -v mydata:/var/lib/mysql mysql       # 命名卷
docker run -v $(pwd):/app node:20-alpine        # 绑定挂载宿主目录
docker run --tmpfs /tmp nginx                    # 内存文件系统
docker run --mount type=bind,src=/host,dst=/container nginx  # 显式挂载
```

## fa-network-wired 网络管理

```bash
docker network create mynet            # 创建自定义 bridge 网络
docker network ls                      # 列出所有网络
docker network inspect mynet           # 网络详情和已连接容器
docker network rm mynet                # 删除网络
docker network prune                   # 清理未使用的网络

docker run --network mynet --name api myimage  # 在指定网络中运行
docker network connect mynet web       # 将运行中容器加入网络
docker network disconnect mynet web    # 从网络断开
```

## fa-copy Docker Compose

```yaml
version: "3.9"
services:
  web:
    build: .                           # 从本地 Dockerfile 构建
    ports:
      - "8080:80"                      # 宿主:容器 端口映射
    volumes:
      - .:/app                         # 绑定挂载，支持热更新
    environment:
      - NODE_ENV=development
    depends_on:
      - db                             # 先启动 db 再启动 web
  db:
    image: postgres:16                 # 使用官方镜像
    volumes:
      - pgdata:/var/lib/postgresql/data  # 持久化数据库数据
    environment:
      POSTGRES_PASSWORD: secret

volumes:
  pgdata:                              # 命名卷声明
```

```bash
docker compose up -d                   # 启动所有服务（后台）
docker compose down                    # 停止并移除容器
docker compose down -v                 # 同时删除数据卷
docker compose logs -f                 # 实时查看所有服务日志
docker compose ps                      # 列出服务容器
docker compose exec web bash           # 进入运行中的服务
docker compose build                   # 重新构建镜像
docker compose pull                    # 拉取最新镜像
docker compose restart web             # 重启单个服务
```

## fa-file-code Dockerfile

```dockerfile
# 多阶段构建：最终镜像更小
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci                             # 安装依赖（比 npm install 更快）
COPY . .
RUN npm run build                      # 编译生产包

FROM node:20-alpine                    # 运行阶段（不含构建工具）
WORKDIR /app
COPY --from=builder /app/dist ./dist   # 仅复制构建产物
COPY --from=builder /app/node_modules ./node_modules
EXPOSE 3000                            # 声明端口（不会自动发布）
HEALTHCHECK --interval=30s CMD wget -qO- http://localhost:3000/health
USER node                              # 以非 root 用户运行
CMD ["node", "dist/main.js"]           # 默认启动命令
```

## fa-broom 清理

```bash
docker system df                       # 查看磁盘占用摘要
docker system prune                    # 清理未使用资源（交互式确认）
docker system prune -a --volumes       # 彻底清理所有未使用资源

docker container prune                 # 清理已停止的容器
docker image prune -a                  # 清理所有未使用镜像
docker volume prune                    # 清理未使用的数据卷
docker network prune                   # 清理未使用的网络

docker rmi $(docker images -q -f dangling=true)   # 删除悬空镜像
docker rm $(docker ps -aq -f status=exited)       # 删除已退出容器
```

## fa-magnifying-glass 检查与调试

```bash
docker inspect web                     # 完整容器配置 (JSON)
docker inspect -f '{{.NetworkSettings.IPAddress}}' web  # 提取特定字段
docker port web                        # 查看端口映射
docker diff web                        # 容器启动后的文件系统变更
docker history nginx:latest            # 查看镜像构建层
docker logs --since 30m -f web 2>&1 | grep ERROR  # 实时错误监控
docker exec -it web sh -c "apt-get update && apt-get install -y curl"  # 安装调试工具
```

## fa-lightbulb 实用技巧

```bash
docker run --init node:20 node app.js  # init 进程，正确处理信号
docker run --read-only --tmpfs /tmp nginx  # 只读文件系统（安全加固）
docker build --no-cache -t myapp .     # 无缓存完整重建
docker build --build-arg VERSION=1.0 -t myapp .  # 传递构建时变量
docker compose up -d --scale worker=3  # 运行 3 个 worker 实例
docker save myapp:1.0 | gzip > myapp.tar.gz  # 导出镜像到文件
docker load < myapp.tar.gz             # 从文件导入镜像
docker system info                     # Docker 引擎信息与配置
```
