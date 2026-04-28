---
title: Docker
icon: fa-docker
primary: "#2496ED"
lang: bash
---

## fa-play Container Lifecycle

```bash
docker create --name web nginx:latest  # create container without starting
docker start web                       # start a stopped container
docker stop web                        # gracefully stop (SIGTERM then SIGKILL)
docker restart web                     # stop then start
docker pause web                       # freeze all processes (cgroups freeze)
docker unpause web                     # unfreeze paused container
docker rm web                          # remove stopped container
docker rm -f web                       # force remove running container
```

## fa-terminal Run

```bash
docker run nginx                       # run in foreground
docker run -d --name web -p 8080:80 nginx  # detached mode, map port
docker run -it ubuntu:22.04 bash       # interactive terminal
docker run --rm alpine echo "hello"    # auto-remove after exit
docker run -e MYSQL_ROOT_PASSWORD=secret mysql  # set env variable
docker run -v /host/data:/container/data nginx   # bind mount volume
docker run --network mynet --name app myimage     # attach to network
docker run --cpus="1.5" --memory="512m" nginx     # limit resources
```

## fa-list Containers

```bash
docker ps                              # running containers only
docker ps -a                           # all containers (including stopped)
docker ps -q                           # quiet mode, IDs only
docker ps --filter "status=running"    # filter by status
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
docker container ls -a                 # alias for docker ps -a
docker stats                           # live resource usage stream
docker top web                         # processes inside container
```

## fa-box Exec & Logs

```bash
docker exec -it web bash               # open interactive shell
docker exec web cat /etc/hosts         # run command in container
docker exec -u root web apt-get update # run as specific user
docker logs web                        # all logs
docker logs -f web                     # follow (tail) logs
docker logs --tail 100 web             # last 100 lines
docker logs --since 2h web             # logs from last 2 hours
docker cp web:/app/config.yml ./config.yml  # copy from container
docker cp ./data web:/app/data               # copy to container
```

## fa-layer-group Images

```bash
docker images                          # list local images
docker pull node:20-alpine             # pull from registry
docker push myrepo/myapp:latest        # push to registry
docker build -t myapp:1.0 .            # build from Dockerfile in cwd
docker build -f Dockerfile.prod -t myapp:prod .  # specify Dockerfile
docker rmi nginx:latest                # remove image
docker image prune                     # remove dangling images
docker image prune -a                  # remove all unused images
docker tag myapp:1.0 myrepo/myapp:latest  # tag image for push
```

## fa-hard-drive Volumes

```bash
docker volume create mydata            # create named volume
docker volume ls                       # list volumes
docker volume inspect mydata           # volume details (mount path etc)
docker volume rm mydata                # remove volume
docker volume prune                    # remove unused volumes

docker run -v mydata:/var/lib/mysql mysql       # named volume
docker run -v $(pwd):/app node:20-alpine        # bind mount host dir
docker run --tmpfs /tmp nginx                    # tmpfs (in-memory)
docker run --mount type=bind,src=/host,dst=/container nginx  # explicit mount
```

## fa-network-wired Networks

```bash
docker network create mynet            # create custom bridge network
docker network ls                      # list networks
docker network inspect mynet           # network details & connected containers
docker network rm mynet                # remove network
docker network prune                   # remove unused networks

docker run --network mynet --name api myimage  # run container in network
docker network connect mynet web       # connect running container to network
docker network disconnect mynet web    # disconnect from network
```

## fa-copy Compose

```yaml
version: "3.9"
services:
  web:
    build: .                           # build from local Dockerfile
    ports:
      - "8080:80"                      # host:container port mapping
    volumes:
      - .:/app                         # bind mount for live reload
    environment:
      - NODE_ENV=development
    depends_on:
      - db                             # start db before web
  db:
    image: postgres:16                 # use official image
    volumes:
      - pgdata:/var/lib/postgresql/data  # persist database data
    environment:
      POSTGRES_PASSWORD: secret

volumes:
  pgdata:                              # named volume declaration
```

```bash
docker compose up -d                   # start all services (detached)
docker compose down                    # stop and remove containers
docker compose down -v                 # also remove volumes
docker compose logs -f                 # follow logs from all services
docker compose ps                      # list service containers
docker compose exec web bash           # shell into running service
docker compose build                   # rebuild images
docker compose pull                    # pull latest images
docker compose restart web             # restart single service
```

## fa-file-code Dockerfile

```dockerfile
# Multi-stage build: smaller final image
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci                             # install deps (faster than npm install)
COPY . .
RUN npm run build                      # compile production bundle

FROM node:20-alpine                    # runtime stage (no build tools)
WORKDIR /app
COPY --from=builder /app/dist ./dist   # only copy build output
COPY --from=builder /app/node_modules ./node_modules
EXPOSE 3000                            # document port (does not publish)
HEALTHCHECK --interval=30s CMD wget -qO- http://localhost:3000/health
USER node                              # run as non-root user
CMD ["node", "dist/main.js"]           # default command
```

## fa-broom Cleanup

```bash
docker system df                       # show disk usage summary
docker system prune                    # remove unused data (interactive)
docker system prune -a --volumes       # remove everything unused

docker container prune                 # remove stopped containers
docker image prune -a                  # remove all unused images
docker volume prune                    # remove unused volumes
docker network prune                   # remove unused networks

docker rmi $(docker images -q -f dangling=true)   # remove dangling images
docker rm $(docker ps -aq -f status=exited)       # remove exited containers
```

## fa-magnifying-glass Inspect & Debug

```bash
docker inspect web                     # full container config (JSON)
docker inspect -f '{{.NetworkSettings.IPAddress}}' web  # extract specific field
docker port web                        # show port mappings
docker diff web                        # filesystem changes since container start
docker history nginx:latest            # show image layers
docker logs --since 30m -f web 2>&1 | grep ERROR  # live error monitoring
docker exec -it web sh -c "apt-get update && apt-get install -y curl"  # install tools
```

## fa-lightbulb Useful Tips

```bash
docker run --init node:20 node app.js  # init process for proper signal handling
docker run --read-only --tmpfs /tmp nginx  # read-only filesystem (security)
docker build --no-cache -t myapp .     # force full rebuild, no cache
docker build --build-arg VERSION=1.0 -t myapp .  # pass build-time variable
docker compose up -d --scale worker=3  # run 3 instances of worker service
docker save myapp:1.0 | gzip > myapp.tar.gz  # export image to file
docker load < myapp.tar.gz             # import image from file
docker system info                     # docker daemon info & config
```
