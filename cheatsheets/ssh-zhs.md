---
title: SSH
icon: fa-key
primary: "#000000"
lang: bash
locale: zhs
---

## fa-terminal 基本连接

```bash
ssh user@host                      # 连接远程主机
ssh -p 2222 user@host              # 自定义端口
ssh user@host "ls -la /tmp"        # 执行命令后退出
ssh -l user host                   # 使用 -l 参数
ssh -v user@host                   # 详细输出（调试）
ssh -vvv user@host                 # 最大详细度
ssh -J jumpuser@jumphost user@target   # 通过跳板机
ssh -o StrictHostKeyChecking=no user@host
ssh -o ConnectTimeout=10 user@host
```

## fa-key 密钥生成 (ssh-keygen)

```bash
ssh-keygen -t ed25519              # Ed25519（推荐）
ssh-keygen -t ed25519 -C "user@host"
ssh-keygen -t rsa -b 4096 -C "user@host"  # RSA 4096 位
ssh-keygen -t ecdsa -b 521         # ECDSA P-521

ssh-keygen -f ~/.ssh/mykey         # 自定义文件名
ssh-keygen -p -f ~/.ssh/id_ed25519 # 修改密码短语
ssh-keygen -y -f ~/.ssh/id_ed25519 # 从私钥输出公钥
ssh-keygen -l -f ~/.ssh/id_ed25519 # 指纹
ssh-keygen -R host                 # 从 known_hosts 移除
ssh-keygen -r host                 # 导出 DNS SSHFP 记录
```

## fa-fingerprint 密钥认证

```bash
ssh-copy-id user@host              # 复制公钥到远程
ssh-copy-id -i ~/.ssh/mykey.pub user@host
ssh-copy-id -p 2222 user@host

# 手动复制
cat ~/.ssh/id_ed25519.pub | ssh user@host "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys"

# 远程端设置权限
chmod 700 ~/.ssh
chmod 600 ~/.ssh/authorized_keys
```

## fa-user-shield ssh-agent

```bash
eval "$(ssh-agent -s)"             # 启动 agent
ssh-add ~/.ssh/id_ed25519          # 添加密钥
ssh-add -l                         # 列出已加载密钥
ssh-add -L                         # 列出公钥
ssh-add -d ~/.ssh/id_ed25519       # 移除密钥
ssh-add -D                         # 移除所有密钥
ssh-add -t 3600 ~/.ssh/id_ed25519  # 添加（有效期 1 小时）
ssh-add -x                         # 锁定 agent
ssh-add -X                         # 解锁 agent
```

## fa-file-lines 配置文件 (~/.ssh/config)

```text
Host dev
    HostName dev.example.com
    User admin
    Port 2222
    IdentityFile ~/.ssh/id_ed25519
    ForwardAgent yes

Host prod-*
    User deploy
    IdentityFile ~/.ssh/prod_key
    StrictHostKeyChecking no

Host jump
    HostName bastion.example.com
    User jumpuser
    IdentityFile ~/.ssh/jump_key

Host *.internal
    ProxyJump jump
    User ubuntu

Host github.com
    User git
    IdentityFile ~/.ssh/github_key
    IdentitiesOnly yes

Host tunnel
    HostName db.example.com
    LocalForward 5432 localhost:5432
    User admin
```

```bash
ssh dev                            # 使用配置
ssh prod-web1                      # 通配符匹配
ssh -G dev                         # 显示生效配置
```

## fa-upload SCP 与 SFTP

```bash
scp file.txt user@host:/tmp/       # 本地 → 远程
scp user@host:/tmp/file.txt ./     # 远程 → 本地
scp -r dir/ user@host:/tmp/        # 递归复制
scp -P 2222 file.txt user@host:/tmp/

sftp user@host
sftp> ls
sftp> cd /var/log
sftp> get syslog ./local_syslog
sftp> put local_file /tmp/
sftp> lls                          # 本地 ls
sftp> lcd ~/Downloads              # 本地 cd
sftp> mkdir remote_dir
sftp> quit
```

## fa-arrow-right 本地隧道

```bash
ssh -L 8080:internal:80 user@bastion
# localhost:8080 → 堡垒机 → internal:80

ssh -L 5432:db.example.com:5432 user@bastion
# localhost:5432 → 堡垒机 → db.example.com:5432

ssh -L 8888:localhost:80 user@host
# localhost:8888 → host → localhost:80（host 上的）

ssh -NL 3306:db:3306 user@bastion  # -N：不执行远程命令
ssh -fNL 8080:internal:80 user@bastion  # -f：后台运行
```

## fa-arrow-left 远程隧道

```bash
ssh -R 9090:localhost:80 user@remote
# remote:9090 → 隧道 → localhost:80（你的机器）

ssh -R 2222:localhost:22 user@remote
# remote:2222 → 隧道 → 你的 SSH 服务

ssh -NR 9090:localhost:8080 user@remote  # 后台远程转发
ssh -R 0.0.0.0:9090:localhost:80 user@remote  # 绑定所有接口
```

```text
# 远程服务器需在 sshd_config 中设置 GatewayPorts yes
# 以允许任何主机连接到 remote:9090
```

## fa-shield-halved 代理跳转 / 堡垒机

```bash
ssh -J jumpuser@bastion user@target
ssh -J j1@host1,j2@host2 user@target   # 多跳

# 在 ~/.ssh/config 中配置
Host target
    HostName 10.0.0.50
    ProxyJump bastion

Host bastion
    HostName bastion.example.com
    User jumpuser

ssh target                          # 自动通过堡垒机
```

```bash
ssh -o ProxyCommand="ssh -W %h:%p bastion" user@target
# 等效于 ProxyJump（旧语法）
```

## fa-network-wired 端口转发

```bash
ssh -L local_port:dest_host:dest_port user@ssh_server   # 本地
ssh -R remote_port:local_host:local_port user@ssh_server # 远程
ssh -D 1080 user@ssh_server                              # SOCKS 代理

curl --socks5 localhost:1080 http://internal-site.local
ssh -ND 1080 user@bastion             # 动态（SOCKS）后台运行

ssh -L 8888:localhost:8888 -L 3000:localhost:3000 user@host  # 多个转发
```

## fa-desktop X11 转发

```bash
ssh -X user@host                   # X11 转发
ssh -Y user@host                   # 信任 X11（安全性较低）
ssh -X user@host firefox           # 运行远程 GUI 程序

ssh -XC user@host                  # 启用压缩
```

```text
# 远程端：/etc/ssh/sshd_config 中 X11Forwarding yes
# 本地端：X 服务器必须运行
# Linux：原生 X11
# macOS：XQuartz
# Windows：Xming, VcXsrv, WSLg
```

## fa-shield SSH 安全加固

```text
# /etc/ssh/sshd_config
Port 2222
PermitRootLogin no
PasswordAuthentication no
PubkeyAuthentication yes
MaxAuthTries 3
ClientAliveInterval 300
ClientAliveCountMax 2
AllowUsers admin deploy
AllowGroups ssh-users
X11Forwarding no
PermitEmptyPasswords no
HostKey /etc/ssh/ssh_host_ed25519_key
KexAlgorithms curve25519-sha256
Ciphers chacha20-poly1305@openssh.com
MACs hmac-sha2-512-etm@openssh.com
```

```bash
sshd -t                            # 验证配置
systemctl reload sshd              # 应用更改
ssh-audit host                     # 审计 SSH 配置
```

## fa-globe SSHFP DNS 记录

```bash
ssh-keygen -r host.example.com     # 生成 SSHFP 记录
ssh-keygen -r host.example.com -D sha256  # SHA-256 指纹
```

```text
; 添加到 DNS 区域
host.example.com. IN SSHFP 1 1 4A3C...
host.example.com. IN SSHFP 1 2 B2D8...
host.example.com. IN SSHFP 4 1 7F2A...
host.example.com. IN SSHFP 4 2 C4E1...
; 算法：1=RSA, 2=DSA, 3=ECDSA, 4=Ed25519
; 指纹：1=SHA-1, 2=SHA-256
```

```bash
ssh -o VerifyHostKeyDNS=yes user@host
```

## fa-wrench 故障排查

```bash
ssh -vvv user@host                 # 详细调试输出
ssh -T user@host                   # 测试连接（无 shell）
ssh -o BatchMode=yes user@host     # 密钥认证失败则报错

# 检查 SSH 服务
systemctl status sshd
ss -tlnp | grep :22

# 连接问题
nc -zv host 22                     # 测试端口
telnet host 22                     # 检查 SSH 横幅

# 密钥问题
ssh-add -l                         # 列出 agent 密钥
ssh -i ~/.ssh/key user@host        # 指定密钥
ssh-keygen -l -f ~/.ssh/known_hosts
ssh-keygen -R host                 # 移除过期的主机密钥

# Agent 转发
ssh -A user@host                   # 启用 agent 转发
ssh user@host "SSH_AUTH_SOCK=$SSH_AUTH_SOCK ssh thirdhost"
```
