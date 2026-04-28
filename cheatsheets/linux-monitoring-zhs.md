---
title: Linux 监控
icon: fa-chart-line
primary: "#FCC624"
lang: bash
locale: zhs
---

## fa-microchip CPU 与进程 (top/htop)

```bash
top                                      # 交互式进程查看器
htop                                     # 增强版 top（需安装）
top -u nginx                             # 按用户过滤
top -p 1234,5678                         # 监控指定 PID

# top 交互快捷键
# P     按 CPU 使用率排序
# M     按内存使用率排序
# k     杀死进程
# z     彩色高亮
# q     退出

mpstat 1 5                               # CPU 统计，每秒1次，共5次
uptime                                   # 系统负载（1分钟、5分钟、15分钟）
nproc                                    # CPU 核心数
```

## fa-list 进程管理 (ps)

```bash
ps aux                                   # 所有进程，BSD 风格
ps -ef                                   # 所有进程，System V 风格
ps aux --sort=-%mem                      # 按内存使用降序排列
ps aux --sort=-%cpu                      # 按 CPU 使用降序排列
ps -u nginx                              # 按用户查看进程
ps -p 1234 -o pid,ppid,cmd,%mem,%cpu     # 自定义格式显示指定 PID
ps aux | grep nginx                      # 按名称查找进程
ps -ef --forest                          # 进程树视图
pgrep -f "nginx"                         # 按模式获取 PID
pidof nginx                              # 按精确名称获取 PID
```

## fa-memory 内存 (free)

```bash
free -h                                  # 人类可读的内存信息
free -m                                  # 以 MB 为单位
free -s 2                                # 每 2 秒刷新一次
vmstat 1 5                               # 虚拟内存统计
```

```bash
#  各字段含义
#  total    总安装内存
#  used     已使用内存
#  free     未使用内存
#  shared   共享内存（tmpfs、IPC）
#  buff/cache  内核缓冲 + 页缓存
#  available   可用于新进程的内存
```

## fa-hard-drive 磁盘使用 (df/du)

```bash
df -h                                    # 磁盘空间，人类可读
df -T                                    # 显示文件系统类型
df -i                                    # 显示 inode 使用情况
du -sh /var/log/*                        # 各目录大小汇总
du -h --max-depth=1 /var                 # 只显示一层深度
du -sh /                                 # 总磁盘使用量
du -ah /home | sort -rh | head -20       # 最大的 20 个文件/目录
lsblk                                    # 列出块设备
fdisk -l                                 # 磁盘分区信息
iostat -x 1                              # 详细 I/O 统计
iotop                                    # 按进程显示 I/O 使用（需 root）
```

## fa-network-wired 网络 (ss/netstat)

```bash
ss -tlnp                                 # 监听中的 TCP 端口及进程
ss -tlnp | grep :80                      # 检查指定端口
ss -s                                    # 连接统计摘要
ss -tn                                   # 已建立的 TCP 连接
ss -un                                   # UDP 连接
ss -tn state established '( dport = :443 or sport = :443 )'  # 按端口过滤

# netstat（旧版，推荐使用 ss）
netstat -tlnp                            # 监听中的 TCP 端口
netstat -anp                             # 所有连接及 PID
netstat -rn                              # 路由表
```

## fa-door-open 打开文件 (lsof)

```bash
lsof                                     # 所有打开的文件
lsof -i :80                              # 谁在使用 80 端口
lsof -i :80,443                          # 多个端口
lsof -u nginx                            # 用户打开的文件
lsof -p 1234                             # 进程打开的文件
lsof /var/log/syslog                     # 谁打开了这个文件
lsof -i tcp                              # 所有 TCP 连接
lsof -i udp                              # 所有 UDP 连接
lsof +D /var/log                         # 目录下被打开的文件
```

## fa-gear 系统服务 (systemctl)

```bash
systemctl list-units --type=service      # 列出所有服务
systemctl list-units --type=service --state=running   # 运行中的服务
systemctl status nginx                   # 服务状态
systemctl start nginx                    # 启动服务
systemctl stop nginx                     # 停止服务
systemctl restart nginx                  # 重启服务
systemctl reload nginx                   # 重载配置（不重启）
systemctl enable nginx                   # 开机自启
systemctl disable nginx                  # 取消开机自启
systemctl is-active nginx                # 检查是否运行中
systemctl is-enabled nginx               # 检查是否开机自启
journalctl -u nginx                      # 服务日志
journalctl -u nginx -f                   # 实时跟踪服务日志
journalctl -u nginx --since "1 hour ago" # 最近 1 小时日志
journalctl --since "2025-04-29 10:00"    # 指定时间之后的日志
```

## fa-fire 网络防火墙 (iptables/ufw)

```bash
# ufw（Ubuntu 默认）
ufw status                               # 查看防火墙状态
ufw allow 80/tcp                         # 允许 80 端口
ufw allow 443/tcp                        # 允许 443 端口
ufw allow from 10.0.0.0/8               # 允许网段
ufw deny 3306                            # 拒绝端口
ufw delete allow 8080                    # 删除规则

# iptables
iptables -L -n -v                        # 列出所有规则
iptables -L -n --line-numbers            # 显示规则编号
iptables -A INPUT -p tcp --dport 80 -j ACCEPT   # 允许 80 端口
iptables -D INPUT 3                      # 按编号删除规则
```

## fa-clock 定时任务

```bash
crontab -l                               # 列出当前用户的定时任务
crontab -e                               # 编辑定时任务
crontab -l -u nginx                      # 列出其他用户的定时任务

#  ┌──────── 分钟 (0-59)
#  │ ┌────── 小时 (0-23)
#  │ │ ┌──── 日 (1-31)
#  │ │ │ ┌── 月 (1-12)
#  │ │ │ │ ┌ 星期 (0-7, 0和7=周日)
#  │ │ │ │ │
#  * * * * * 命令

#  */5 * * * *     /opt/healthcheck.sh        # 每 5 分钟
#  0 2 * * *      /opt/backup.sh              # 每天凌晨 2 点
#  0 0 1 * *      /opt/monthly-report.sh      # 每月 1 号
#  0 9 * * 1-5    /opt/weekday-job.sh          # 工作日早 9 点

systemctl list-timers                    # 列出 systemd 定时器
timedatectl                              # 时区和 NTP 状态
```

## fa-user 用户与登录

```bash
who                                      # 已登录用户
w                                        # 已登录用户及活动
whoami                                   # 当前用户名
id                                       # 用户 ID 和所属组
last                                     # 登录历史
lastb                                    # 失败的登录尝试
lastlog                                  # 所有用户最近登录时间

# 用户管理
useradd -m -s /bin/bash newuser          # 创建用户（含主目录）
userdel -r olduser                       # 删除用户及主目录
passwd username                          # 修改密码
groups username                          # 查看用户所属组
usermod -aG docker username              # 将用户加入组
```

## fa-lightbulb 快速健康检查

```bash
# 一键系统概览
echo "=== CPU ===" && nproc && uptime
echo "=== 内存 ===" && free -h
echo "=== 磁盘 ===" && df -h /
echo "=== CPU 前5 ===" && ps aux --sort=-%cpu | head -6
echo "=== 内存前5 ===" && ps aux --sort=-%mem | head -6
echo "=== 监听端口 ===" && ss -tlnp
echo "=== 失败服务 ===" && systemctl --failed

# 查找最大的文件
du -ah / 2>/dev/null | sort -rh | head -20

# 查找已删除但仍被占用的文件
lsof +L1

# 查看最近重启记录
last reboot | head -5

# 内核消息（硬件问题）
dmesg -T | grep -i error
```
