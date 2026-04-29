---
title: direnv
icon: fa-right-from-bracket
primary: "#D946EF"
lang: bash
---

## fa-rotate Setup & Hook

```bash
eval "$(direnv hook bash)"         # bash (~/.bashrc)
eval "$(direnv hook zsh)"          # zsh (~/.zshrc)
direnv hook fish | source          # fish (~/.config/fish/config.fish)
eval "$(direnv hook elvish)"       # elvish

# Verify hook is active
echo $DIRENV_DIR
direnv status
```

## fa-file .envrc Basics

```bash
# Create .envrc in project root
cat > .envrc << 'EOF'
export PROJECT_NAME="myapp"
export NODE_ENV="development"
export DATABASE_URL="postgres://localhost/myapp_dev"
PATH_add bin
EOF

direnv allow                       # approve the .envrc
direnv deny                        # deny the .envrc
direnv reload                      # reload after manual edit
```

## fa-arrow-right-from-bracket Export Variables

```bash
export API_KEY="sk-test-123"
export DEBUG="true"
export LOG_LEVEL="debug"
export APP_PORT=8080

# Expand existing variables
export PATH="$HOME/.local/bin:$PATH"
export PYTHONPATH="$PWD/src:$PYTHONPATH"

# Use subshell results
export VERSION=$(git describe --tags --always)
export BRANCH=$(git rev-parse --abbrev-ref HEAD)
```

## fa-route PATH Manipulation

```bash
PATH_add node_modules/.bin         # prepend to PATH
PATH_add ./bin                     # prepend ./bin
PATH_add "$HOME/go/bin"            # prepend Go bin
PATH_add .venv/bin                 # prepend Python venv

# MANPATH manipulation
MANPATH_add /usr/local/share/man

# Remove path entries on unload
PATH_rm node_modules/.bin
```

## fa-table-cells Layout Functions

```bash
# Python virtual environment
layout python

# Python with specific version
layout python3.11

# Node.js layout
layout node

# Go layout (sets GOPATH)
layout go

# Ruby layout
layout ruby

# Custom layout in function
layout_poetry() {
  source $(poetry env info --path)/bin/activate
}
layout_poetry
```

## fa-link source_env / source_up

```bash
# Source another .envrc (relative path)
source_env ../shared/.envrc

# Source from environment variable
source_env "$HOME/.config/direnv/common.envrc"

# Walk up directory tree to find .envrc
source_up

# Source a .env file
source_env .env

# Conditional sourcing
if [ -f .env.local ]; then
  source_env .env.local
fi
```

## fa-code use (node/python/go)

```bash
# Node.js version manager
use node                          # use .nvmrc / .node-version
use node 18                       # use specific node version

# Python version manager
use python 3.11                   # use specific python version
use python                        # use .python-version

# Go version
use go 1.21                       # use specific go version

# Ruby version
use ruby 3.2                      # use specific ruby version

# Multiple tools
use node 18
use python 3.11
```

## fa-eye watch_file

```bash
# Watch specific files for changes
watch_file .env
watch_file .nvmrc
watch_file .python-version
watch_file go.mod
watch_file pyproject.toml
watch_file package.json

# Watch with glob
watch_file*.json

# Reload when config changes
watch_file docker-compose.yaml
export COMPOSE_FILE="docker-compose.yaml"
```

## fa-file-lines dotenv Support

```bash
# Load .env file automatically
dotenv                            # loads .env
dotenv .env.local                 # load specific file
dotenv .env.development           # load development env

# Multiple dotenv files
dotenv .env .env.local

# Conditional dotenv
if [ -f .env ]; then
  dotenv
fi

# dotenv with fallback
dotenv_if_exists .env.local
```

## fa-wrench Custom stdlib Functions

```bash
# Define custom functions in .envrc
has_command() {
  command -v "$1" &>/dev/null
}

set_java_home() {
  if has_command java; then
    export JAVA_HOME=$(dirname $(dirname $(readlink -f $(which java))))
  fi
}

set_java_home

# Lambda-style
export_sha256() {
  export "$1"=$(sha256sum "$2" | cut -d' ' -f1)
}
```

## fa-check Allow / Deny

```bash
direnv allow                       # approve current .envrc
direnv allow .envrc                # approve specific file
direnv deny                        # deny current .envrc
direnv deny .envrc                 # deny specific file

# List allowed/denied
direnv status

# Edit and re-allow
$EDITOR .envrc && direnv allow

# Allow from outside directory
direnv allow /path/to/project/.envrc
```

## fa-broom Unload & Cleanup

```bash
# Automatic unload when leaving directory
# direnv unloads vars and restores previous env

# Manual unload
direnv export bash | reverse

# Cleanup function in .envrc
cleanup() {
  rm -f /tmp/$PROJECT_NAME-*
}
# Note: cleanup runs in subshell, use trap for real cleanup

# Check if in direnv context
if [ -n "$DIRENV_DIR" ]; then
  echo "inside direnv"
fi
```

## fa-layer-group Nested .envrc

```bash
# Parent: ~/projects/.envrc
export COMPANY="acme"
export DEFAULT_REGION="us-east-1"

# Child: ~/projects/myapp/.envrc
source_up                         # inherit parent settings
export PROJECT="myapp"
PATH_add ./node_modules/.bin

# Override parent values
export REGION="eu-west-1"         # overrides DEFAULT_REGION

# Directory structure
# ~/projects/
#   .envrc          (shared config)
#   myapp/
#     .envrc        (project config, sources parent)
#     api/
#       .envrc      (api-specific config)
```

## fa-lightbulb Common Patterns

```bash
# Full-stack project .envrc
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

# Per-environment configuration
source_env .envrc.$(hostname)
if [ -f ".envrc.local" ]; then
  source_env .envrc.local
fi

# Docker Compose project
export COMPOSE_PROJECT_NAME="myapp"
export COMPOSE_FILE="docker-compose.yaml:docker-compose.dev.yaml"

# Secrets from vault
export DB_PASSWORD=$(cat /run/secrets/db_password 2>/dev/null || echo "dev")
```
