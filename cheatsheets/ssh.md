---
title: SSH
icon: fa-key
primary: "#000000"
lang: bash
---

## fa-terminal Basic Connection

```bash
ssh user@host                      # connect to remote
ssh -p 2222 user@host              # custom port
ssh user@host "ls -la /tmp"        # run command and exit
ssh -l user host                   # with -l flag
ssh -v user@host                   # verbose (debug)
ssh -vvv user@host                 # maximum verbosity
ssh -J jumpuser@jumphost user@target   # via jump host
ssh -o StrictHostKeyChecking=no user@host
ssh -o ConnectTimeout=10 user@host
```

## fa-key Key Generation (ssh-keygen)

```bash
ssh-keygen -t ed25519              # Ed25519 (recommended)
ssh-keygen -t ed25519 -C "user@host"
ssh-keygen -t rsa -b 4096 -C "user@host"  # RSA 4096-bit
ssh-keygen -t ecdsa -b 521         # ECDSA P-521

ssh-keygen -f ~/.ssh/mykey         # custom filename
ssh-keygen -p -f ~/.ssh/id_ed25519 # change passphrase
ssh-keygen -y -f ~/.ssh/id_ed25519 # print public key from private
ssh-keygen -l -f ~/.ssh/id_ed25519 # fingerprint
ssh-keygen -R host                 # remove host from known_hosts
ssh-keygen -r host                 # export DNS SSHFP records
```

## fa-fingerprint Key-based Auth

```bash
ssh-copy-id user@host              # copy key to remote
ssh-copy-id -i ~/.ssh/mykey.pub user@host
ssh-copy-id -p 2222 user@host

# manual copy
cat ~/.ssh/id_ed25519.pub | ssh user@host "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys"

# remote side
chmod 700 ~/.ssh
chmod 600 ~/.ssh/authorized_keys
```

## fa-user-shield ssh-agent

```bash
eval "$(ssh-agent -s)"             # start agent
ssh-add ~/.ssh/id_ed25519          # add key to agent
ssh-add -l                         # list loaded keys
ssh-add -L                         # list public keys
ssh-add -d ~/.ssh/id_ed25519       # remove key
ssh-add -D                         # remove all keys
ssh-add -t 3600 ~/.ssh/id_ed25519  # add with 1hr lifetime
ssh-add -x                         # lock agent
ssh-add -X                         # unlock agent
```

## fa-file-lines Config File (~/.ssh/config)

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
ssh dev                            # uses config
ssh prod-web1                      # wildcard match
ssh -G dev                         # show effective config
```

## fa-upload SCP & SFTP

```bash
scp file.txt user@host:/tmp/       # local → remote
scp user@host:/tmp/file.txt ./     # remote → local
scp -r dir/ user@host:/tmp/        # recursive copy
scp -P 2222 file.txt user@host:/tmp/

sftp user@host
sftp> ls
sftp> cd /var/log
sftp> get syslog ./local_syslog
sftp> put local_file /tmp/
sftp> lls                          # local ls
sftp> lcd ~/Downloads              # local cd
sftp> mkdir remote_dir
sftp> quit
```

## fa-arrow-right SSH Tunneling (Local)

```bash
ssh -L 8080:internal:80 user@bastion
# localhost:8080 → bastion → internal:80

ssh -L 5432:db.example.com:5432 user@bastion
# localhost:5432 → bastion → db.example.com:5432

ssh -L 8888:localhost:80 user@host
# localhost:8888 → host → localhost:80 (on host)

ssh -NL 3306:db:3306 user@bastion  # -N: no remote command
ssh -fNL 8080:internal:80 user@bastion  # -f: background
```

## fa-arrow-left SSH Tunneling (Remote)

```bash
ssh -R 9090:localhost:80 user@remote
# remote:9090 → tunnel → localhost:80 (your machine)

ssh -R 2222:localhost:22 user@remote
# remote:2222 → tunnel → your SSH server

ssh -NR 9090:localhost:8080 user@remote  # background remote forward
ssh -R 0.0.0.0:9090:localhost:80 user@remote  # bind all interfaces
```

```text
# Remote server: GatewayPorts yes in sshd_config
# to allow connections from any host to remote:9090
```

## fa-shield-halved ProxyJump / Bastion

```bash
ssh -J jumpuser@bastion user@target
ssh -J j1@host1,j2@host2 user@target   # multi-hop

# in ~/.ssh/config
Host target
    HostName 10.0.0.50
    ProxyJump bastion

Host bastion
    HostName bastion.example.com
    User jumpuser

ssh target                          # auto-through bastion
```

```bash
ssh -o ProxyCommand="ssh -W %h:%p bastion" user@target
# equivalent to ProxyJump (legacy syntax)
```

## fa-network-wired Port Forwarding

```bash
ssh -L local_port:dest_host:dest_port user@ssh_server   # local
ssh -R remote_port:local_host:local_port user@ssh_server # remote
ssh -D 1080 user@ssh_server                              # SOCKS proxy

curl --socks5 localhost:1080 http://internal-site.local
ssh -ND 1080 user@bastion             # dynamic (SOCKS) in background

ssh -L 8888:localhost:8888 -L 3000:localhost:3000 user@host  # multiple forwards
```

## fa-desktop X11 Forwarding

```bash
ssh -X user@host                   # X11 forwarding
ssh -Y user@host                   # trusted X11 (less secure)
ssh -X user@host firefox           # run remote GUI app

ssh -XC user@host                  # with compression
```

```text
# Remote: X11Forwarding yes in /etc/ssh/sshd_config
# Local: X server must be running
# Linux: native X11
# macOS: XQuartz
# Windows: Xming, VcXsrv, WSLg
```

## fa-shield SSH Hardening

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
sshd -t                            # validate config
systemctl reload sshd              # apply changes
ssh-audit host                     # audit SSH config
```

## fa-globe SSHFP DNS Records

```bash
ssh-keygen -r host.example.com     # generate SSHFP records
ssh-keygen -r host.example.com -D sha256  # SHA-256 fingerprints
```

```text
; Add to DNS zone
host.example.com. IN SSHFP 1 1 4A3C...
host.example.com. IN SSHFP 1 2 B2D8...
host.example.com. IN SSHFP 4 1 7F2A...
host.example.com. IN SSHFP 4 2 C4E1...
; Algorithm: 1=RSA, 2=DSA, 3=ECDSA, 4=Ed25519
; Fingerprint: 1=SHA-1, 2=SHA-256
```

```bash
ssh -o VerifyHostKeyDNS=yes user@host
```

## fa-wrench Troubleshooting

```bash
ssh -vvv user@host                 # verbose debug output
ssh -T user@host                   # test connection (no shell)
ssh -o BatchMode=yes user@host     # fail if key auth not possible

# check SSH service
systemctl status sshd
ss -tlnp | grep :22

# connection issues
nc -zv host 22                     # test port
telnet host 22                     # check SSH banner

# key issues
ssh-add -l                         # list agent keys
ssh -i ~/.ssh/key user@host        # specify identity
ssh-keygen -l -f ~/.ssh/known_hosts
ssh-keygen -R host                 # remove stale host key

# agent forwarding
ssh -A user@host                   # enable agent forwarding
ssh user@host "SSH_AUTH_SOCK=$SSH_AUTH_SOCK ssh thirdhost"
```
