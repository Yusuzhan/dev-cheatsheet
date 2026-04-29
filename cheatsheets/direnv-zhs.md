---
title: direnv
icon: fa-right-from-bracket
primary: "#D946EF"
lang: bash
locale: zhs
---

## fa-rotate 安装与钩子

```bash
eval "$(direnv hook bash)"         # bash (~/.bashrc)
eval "$(direnv hook zsh)"          # zsh (~/.zshrc)
direnv hook fish | source          # fish (~/.config/fish/config.fish)
eval "$(direnv hook elvish)"       # elvish

# 验证钩子已激活
echo $DIRENV_DIR
direnv status
```

## fa-file .envrc 基础

```bash
# 在项目根目录创建 .envrc
cat > .envrc << 'EOF'
export PROJECT_NAME="myapp"
export NODE_ENV="development"
export DATABASE_URL="postgres://localhost/myapp_dev"
PATH_add bin
EOF

direnv allow                       # 批准 .envrc
direnv deny                        # 拒绝 .envrc
direnv reload                      # 手动编辑后重新加载
```

## fa-arrow-right-from-bracket 导出变量

```bash
export API_KEY="sk-test-123"
export DEBUG="true"
export LOG_LEVEL="debug"
export APP_PORT=8080

# 展开已有变量
export PATH="$HOME/.local/bin:$PATH"
export PYTHONPATH="$PWD/src:$PYTHONPATH"

# 使用子命令结果
export VERSION=$(git describe --tags --always)
export BRANCH=$(git rev-parse --abbrev-ref HEAD)
```

## fa-route PATH 操作

```bash
PATH_add node_modules/.bin         # 前置到 PATH
PATH_add ./bin                     # 前置 ./bin
PATH_add "$HOME/go/bin"            # 前置 Go bin
PATH_add .venv/bin                 # 前置 Python venv

# MANPATH 操作
MANPATH_add /usr/local/share/man

# 卸载时移除路径
PATH_rm node_modules/.bin
```

## fa-table-cells Layout 函数

```bash
# Python 虚拟环境
layout python

# 指定 Python 版本
layout python3.11

# Node.js 布局
layout node

# Go 布局 (设置 GOPATH)
layout go

# Ruby 布局
layout ruby

# 自定义布局函数
layout_poetry() {
  source $(poetry env info --path)/bin/activate
}
layout_poetry
```

## fa-link source_env / source_up

```bash
# 引入另一个 .envrc (相对路径)
source_env ../shared/.envrc

# 从环境变量引用
source_env "$HOME/.config/direnv/common.envrc"

# 向上遍历目录查找 .envrc
source_up

# 引入 .env 文件
source_env .env

# 条件引入
if [ -f .env.local ]; then
  source_env .env.local
fi
```

## fa-code use (node/python/go)

```bash
# Node.js 版本管理
use node                          # 使用 .nvmrc / .node-version
use node 18                       # 使用指定 node 版本

# Python 版本管理
use python 3.11                   # 使用指定 python 版本
use python                        # 使用 .python-version

# Go 版本
use go 1.21                       # 使用指定 go 版本

# Ruby 版本
use ruby 3.2                      # 使用指定 ruby 版本

# 多工具组合
use node 18
use python 3.11
```

## fa-eye watch_file

```bash
# 监视特定文件变化
watch_file .env
watch_file .nvmrc
watch_file .python-version
watch_file go.mod
watch_file pyproject.toml
watch_file package.json

# 使用 glob 监视
watch_file*.json

# 配置变化时重新加载
watch_file docker-compose.yaml
export COMPOSE_FILE="docker-compose.yaml"
```

## fa-file-lines dotenv 支持

```bash
# 自动加载 .env 文件
dotenv                            # 加载 .env
dotenv .env.local                 # 加载指定文件
dotenv .env.development           # 加载开发环境变量

# 多个 dotenv 文件
dotenv .env .env.local

# 条件 dotenv
if [ -f .env ]; then
  dotenv
fi

# 带回退的 dotenv
dotenv_if_exists .env.local
```

## fa-wrench 自定义标准库函数

```bash
# 在 .envrc 中定义自定义函数
has_command() {
  command -v "$1" &>/dev/null
}

set_java_home() {
  if has_command java; then
    export JAVA_HOME=$(dirname $(dirname $(readlink -f $(which java))))
  fi
}

set_java_home

# Lambda 风格
export_sha256() {
  export "$1"=$(sha256sum "$2" | cut -d' ' -f1)
}
```

## fa-check 允许 / 拒绝

```bash
direnv allow                       # 批准当前 .envrc
direnv allow .envrc                # 批准指定文件
direnv deny                        # 拒绝当前 .envrc
direnv deny .envrc                 # 拒绝指定文件

# 列出已允许/拒绝的
direnv status

# 编辑并重新批准
$EDITOR .envrc && direnv allow

# 从目录外批准
direnv allow /path/to/project/.envrc
```

## fa-broom 卸载与清理

```bash
# 离开目录时自动卸载
# direnv 卸载变量并恢复之前的环境

# 手动卸载
direnv export bash | reverse

# .envrc 中的清理函数
cleanup() {
  rm -f /tmp/$PROJECT_NAME-*
}

# 检查是否在 direnv 上下文中
if [ -n "$DIRENV_DIR" ]; then
  echo "在 direnv 环境中"
fi
```

## fa-layer-group 嵌套 .envrc

```bash
# 父级: ~/projects/.envrc
export COMPANY="acme"
export DEFAULT_REGION="us-east-1"

# 子级: ~/projects/myapp/.envrc
source_up                         # 继承父级设置
export PROJECT="myapp"
PATH_add ./node_modules/.bin

# 覆盖父级值
export REGION="eu-west-1"         # 覆盖 DEFAULT_REGION

# 目录结构
# ~/projects/
#   .envrc          (共享配置)
#   myapp/
#     .envrc        (项目配置, 引入父级)
#     api/
#       .envrc      (API 专用配置)
```

## fa-lightbulb 常用模式

```bash
# 全栈项目 .envrc
cat > .envrc << 'EOF'
export APP_ENV="development"
dotenv .env.development

layout python
PATH_add node_modules/.bin
PATH_add ./scripts

export DATABASE_URL="postgres://localhost:5432/myapp"
export REDIS_URL="redis://localhost:6379"
export API_PORT=3000

watch_file pyproject.toml
watch_file package.json
EOF

# 按环境配置
source_env .envrc.$(hostname)
if [ -f ".envrc.local" ]; then
  source_env .envrc.local
fi

# Docker Compose 项目
export COMPOSE_PROJECT_NAME="myapp"
export COMPOSE_FILE="docker-compose.yaml:docker-compose.dev.yaml"

# 从密钥库读取
export DB_PASSWORD=$(cat /run/secrets/db_password 2>/dev/null || echo "dev")
```
