---
title: Docker
icon: fa-docker
primary: "#2496ED"
lang: bash
---

## fa-play Container Lifecycle

```bash
docker create --name web nginx:latest
docker start web
docker stop web
docker restart web
docker pause web
docker unpause web
docker rm web
docker rm -f web
```

## fa-terminal Run

```bash
docker run nginx
docker run -d --name web -p 8080:80 nginx
docker run -it ubuntu:22.04 bash
docker run --rm alpine echo "hello"
docker run -e MYSQL_ROOT_PASSWORD=secret mysql
docker run -v /host/data:/container/data nginx
docker run --network mynet --name app myimage
docker run --cpus="1.5" --memory="512m" nginx
```

## fa-list Containers

```bash
docker ps
docker ps -a
docker ps -q
docker ps --filter "status=running"
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
docker container ls -a
docker stats
docker top web
```

## fa-box Exec & Logs

```bash
docker exec -it web bash
docker exec web cat /etc/hosts
docker exec -u root web apt-get update
docker logs web
docker logs -f web
docker logs --tail 100 web
docker logs --since 2h web
docker cp web:/app/config.yml ./config.yml
docker cp ./data web:/app/data
```

## fa-layer-group Images

```bash
docker images
docker pull node:20-alpine
docker push myrepo/myapp:latest
docker build -t myapp:1.0 .
docker build -f Dockerfile.prod -t myapp:prod .
docker rmi nginx:latest
docker image prune
docker image prune -a
docker tag myapp:1.0 myrepo/myapp:latest
```

## fa-hard-drive Volumes

```bash
docker volume create mydata
docker volume ls
docker volume inspect mydata
docker volume rm mydata
docker volume prune

docker run -v mydata:/var/lib/mysql mysql
docker run -v $(pwd):/app node:20-alpine
docker run --tmpfs /tmp nginx
docker run --mount type=bind,src=/host,dst=/container nginx
```

## fa-network-wired Networks

```bash
docker network create mynet
docker network ls
docker network inspect mynet
docker network rm mynet
docker network prune

docker run --network mynet --name api myimage
docker network connect mynet web
docker network disconnect mynet web
```

## fa-copy Compose

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
docker compose down -v
docker compose logs -f
docker compose ps
docker compose exec web bash
docker compose build
docker compose pull
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

## fa-broom Cleanup

```bash
docker system df
docker system prune
docker system prune -a --volumes

docker container prune
docker image prune -a
docker volume prune
docker network prune

docker rmi $(docker images -q -f dangling=true)
docker rm $(docker ps -aq -f status=exited)
```

## fa-magnifying-glass Inspect & Debug

```bash
docker inspect web
docker inspect -f '{{.NetworkSettings.IPAddress}}' web
docker port web
docker diff web
docker history nginx:latest
docker logs --since 30m -f web 2>&1 | grep ERROR
docker exec -it web sh -c "apt-get update && apt-get install -y curl"
```

## fa-lightbulb Useful Tips

```bash
docker run --init node:20 node app.js
docker run --read-only --tmpfs /tmp nginx
docker build --no-cache -t myapp .
docker build --build-arg VERSION=1.0 -t myapp .
docker compose up -d --scale worker=3
docker save myapp:1.0 | gzip > myapp.tar.gz
docker load < myapp.tar.gz
docker system info
```
