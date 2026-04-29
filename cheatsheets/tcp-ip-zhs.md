---
title: TCP/IP
icon: fa-network-wired
primary: "#0078D4"
lang: bash
locale: zhs
---

## fa-layer-group OSI 与 TCP/IP 模型

```text
OSI 模型                TCP/IP 模型           协议示例
7 应用层                应用层                HTTP, FTP, DNS, SSH, SMTP
6 表示层                应用层                TLS, JPEG, JSON, XML
5 会话层                应用层                NetBIOS, RPC, PPTP
4 传输层                传输层                TCP, UDP, SCTP, QUIC
3 网络层                网际层                IP, ICMP, IGMP, ARP
2 数据链路层            网络接口层            Ethernet, Wi-Fi, PPP, VLAN
1 物理层                网络接口层            光纤, 网线, 无线电波
```

```text
封装（发送）：数据 → 报文段 → 数据包 → 帧 → 比特
解封（接收）：比特 → 帧 → 数据包 → 报文段 → 数据
```

## fa-location-dot IP 地址（v4/v6）

```bash
ip addr show
ip addr show eth0
ip addr add 192.168.1.10/24 dev eth0
ip addr del 192.168.1.10/24 dev eth0
```

```text
IPv4 地址分类：
A: 1.0.0.0     – 126.255.255.255   /8    （大型网络）
B: 128.0.0.0   – 191.255.255.255   /16   （中型网络）
C: 192.0.0.0   – 223.255.255.255   /24   （小型网络）

私有地址范围：
10.0.0.0/8          172.16.0.0/12        192.168.0.0/16
127.0.0.0/8（回环）       169.254.0.0/16（链路本地）
```

```text
IPv6：
2001:0db8:85a3:0000:0000:8a2e:0370:7334
2001:db8:85a3::8a2e:370:7334          # 压缩格式
::1                                    # 回环地址
fe80::                                 # 链路本地
ff02::                                 # 组播
```

```bash
ip -6 addr show
ip -6 route show
ping6 ::1
```

## fa-scissors 子网划分与 CIDR

```text
CIDR        子网掩码              主机数       可用主机数
/8          255.0.0.0            16,777,214
/16         255.255.0.0          65,534
/20         255.255.240.0        4,094
/24         255.255.255.0        254
/25         255.255.255.128      126
/26         255.255.255.192      62
/27         255.255.255.224      30
/28         255.255.255.240      14
/30         255.255.255.252      2
/32         255.255.255.255      1（主机路由）
```

```bash
ipcalc 192.168.1.0/24
ipcalc 10.0.0.0/20
sipcalc 172.16.0.0/12

ip route add 10.10.0.0/16 via 192.168.1.1
ip route add default via 192.168.1.1
ip route show
```

## fa-handshake TCP 三次握手

```text
三次握手（建立连接）：

Client                          Server
  |--- SYN (seq=x) ------------->|     1. 客户端发送 SYN
  |<-- SYN-ACK (seq=y,ack=x+1) -|     2. 服务端发送 SYN-ACK
  |--- ACK (ack=y+1) ----------->|     3. 客户端发送 ACK
  |        连接已建立              |

四次挥手（断开连接）：

Client                          Server
  |--- FIN (seq=u) ------------->|
  |<-- ACK (ack=u+1) -----------|
  |<-- FIN (seq=w) -------------|
  |--- ACK (ack=w+1) ----------->|
  |        连接已关闭              |
```

```bash
tcpdump -i eth0 'tcp[tcpflags] & (tcp-syn|tcp-fin) != 0'
ss -tn state established
ss -tn state time-wait
```

## fa-circle-nodes TCP 状态

```text
CLOSED → SYN_SENT → ESTABLISHED → FIN_WAIT_1 → FIN_WAIT_2 → TIME_WAIT → CLOSED
LISTEN → SYN_RCVD → ESTABLISHED → CLOSE_WAIT → LAST_ACK → CLOSED

关键状态：
LISTEN        服务端等待连接
SYN_SENT      客户端已发 SYN，等待 SYN-ACK
SYN_RCVD      服务端收到 SYN，已发 SYN-ACK
ESTABLISHED   连接建立，可传输数据
FIN_WAIT_1    已发 FIN，等待 ACK
FIN_WAIT_2    收到 FIN 的 ACK，等待对端 FIN
CLOSE_WAIT    收到对端 FIN，等待本端关闭
LAST_ACK      CLOSE_WAIT 后发 FIN，等待 ACK
TIME_WAIT     主动关闭，等待 2*MSL 后清理
```

```bash
ss -tna                           # 所有 TCP 状态
ss -tn state time-wait            # TIME_WAIT 套接字
ss -tn state close-wait           # CLOSE_WAIT 套接字
cat /proc/sys/net/ipv4/tcp_fin_timeout
cat /proc/sys/net/ipv4/ip_local_port_range
```

## fa-bolt UDP

```text
TCP vs UDP：
TCP：面向连接、可靠、有序、流量控制、拥塞控制
UDP：无连接、不保证可靠、不保证顺序、低开销、低延迟

适用 UDP：DNS, DHCP, TFTP, SNMP, NTP, VoIP, 视频流, 游戏
适用 TCP：HTTP, SSH, FTP, SMTP, 数据库连接
```

```bash
nc -u -l 12345                    # UDP 监听
nc -u 192.168.1.10 12345          # UDP 发送
echo "hello" | nc -u -w1 10.0.0.1 514

nmap -sU -p 53,67,68,123 10.0.0.1    # UDP 端口扫描
```

## fa-plug 常用端口

```text
FTP 数据      20/TCP     FTP 控制      21/TCP
SSH           22/TCP     Telnet        23/TCP
SMTP          25/TCP     DNS           53/TCP/UDP
DHCP 服务端   67/UDP     DHCP 客户端   68/UDP
TFTP          69/UDP     HTTP          80/TCP
POP3          110/TCP    NTP           123/UDP
NetBIOS       137-139    IMAP          143/TCP
SNMP          161/UDP    HTTPS         443/TCP
Syslog        514/UDP    SMTPS         465/TCP
LDAPS         636/TCP    iSCSI         860/TCP
MySQL         3306/TCP   RDP           3389/TCP
PostgreSQL    5432/TCP   Redis         6379/TCP
HTTP 备用     8080/TCP   HTTPS 备用    8443/TCP
```

## fa-satellite-dish ping / traceroute

```bash
ping -c 4 8.8.8.8                 # 发送 4 次
ping -i 0.5 10.0.0.1              # 间隔 0.5 秒
ping -W 2 10.0.0.1                # 超时 2 秒
ping -s 1400 -M do 10.0.0.1      # MTU 探测（不分片）

traceroute 8.8.8.8                # 基于 UDP
traceroute -T -p 443 example.com  # 基于 TCP
mtr -rwzc 50 example.com          # 持续 traceroute
tracepath 8.8.8.8                 # 路径 MTU 发现

ping6 -c 4 2001:4860:4860::8888
traceroute6 2001:4860:4860::8888
```

## fa-table-list netstat / ss

```bash
ss -tlnp                           # TCP 监听端口及进程
ss -tnp                            # TCP 已建立连接
ss -tunlp                          # TCP + UDP 监听
ss -s                              # 套接字统计
ss -tn state established '( dport = :443 or sport = :443 )'
ss -tn dst 192.168.1.0/24         # 按目标地址过滤

netstat -tlnp                      # 传统方式：TCP 监听
netstat -tulnp                     # 传统方式：TCP + UDP 监听
netstat -tn                        # 传统方式：TCP 已建立
netstat -s                         # 传统方式：统计信息
netstat -rn                        # 传统方式：路由表
```

## fa-magnifying-glass tcpdump

```bash
tcpdump -i eth0                    # eth0 上所有流量
tcpdump -i eth0 -nn host 10.0.0.1
tcpdump -i eth0 -nn port 80
tcpdump -i eth0 -nn src net 192.168.0.0/16
tcpdump -i eth0 -nn 'tcp[tcpflags] & tcp-syn != 0'
tcpdump -i eth0 -nn 'port 53 and udp'
tcpdump -i any -nn -w capture.pcap
tcpdump -r capture.pcap -nn
tcpdump -i eth0 -nn -c 100 -vv
tcpdump -i eth0 -nn -XX port 443
```

## fa-crosshairs nmap

```bash
nmap -sT 10.0.0.1                  # TCP 全连接扫描
nmap -sS 10.0.0.1                  # TCP SYN（隐蔽）扫描
nmap -sU 10.0.0.1                  # UDP 扫描
nmap -sV -sT 10.0.0.1              # 服务版本检测
nmap -O 10.0.0.1                   # 操作系统检测
nmap -sn 192.168.1.0/24            # 主机发现（Ping 扫描）
nmap -p 22,80,443 10.0.0.1         # 指定端口
nmap -p- 10.0.0.1                  # 全端口 65535
nmap -A 10.0.0.1                   # 综合扫描（OS + 版本 + 脚本）
nmap -sT -sV --script vuln 10.0.0.1
```

## fa-shield-halved iptables

```bash
iptables -L -n -v                              # 列出所有规则
iptables -A INPUT -p tcp --dport 22 -j ACCEPT  # 允许 SSH
iptables -A INPUT -p tcp --dport 80 -j ACCEPT  # 允许 HTTP
iptables -A INPUT -p tcp --dport 443 -j ACCEPT # 允许 HTTPS
iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
iptables -A INPUT -i lo -j ACCEPT              # 允许回环
iptables -P INPUT DROP                          # 默认拒绝
iptables -A INPUT -s 10.0.0.0/8 -j ACCEPT     # 允许子网

iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 8080

iptables -D INPUT 3                            # 删除第 3 条规则
iptables -F                                    # 清空所有规则
iptables-save > /etc/iptables/rules.v4
```

## fa-right-left NAT 与 PAT

```text
NAT 类型：
静态 NAT     1:1 映射    192.168.1.10 ↔ 203.0.113.10
动态 NAT     地址池映射   内网主机共享公网 IP 池
PAT (SNAT)   N:1 映射    多个内网 IP → 1 个公网 IP + 端口

SNAT（源地址转换）：修改源地址（出站）
DNAT（目标地址转换）：修改目标地址（端口转发）
MASQUERADE：         动态 IP 的 SNAT
```

```bash
ip addr show                     # 查看地址
iptables -t nat -L -n -v        # 查看 NAT 规则

iptables -t nat -A POSTROUTING -s 192.168.1.0/24 -o eth0 -j MASQUERADE
iptables -t nat -A PREROUTING -p tcp -d 203.0.113.10 --dport 80 -j DNAT --to-destination 192.168.1.20:8080

cat /proc/sys/net/ipv4/ip_forward   # 检查转发状态
sysctl -w net.ipv4.ip_forward=1     # 开启转发
```

## fa-wrench 故障排查

```bash
ip link show                       # 接口状态
ethtool eth0                       # 链路速率、双工、错误
ethtool -S eth0                    # 接口统计
mii-tool eth0                      # 链路状态（旧工具）

arp -an                            # ARP 表
ip neigh show                      # 邻居表
arping -I eth0 192.168.1.1         # ARP ping

nslookup example.com               # DNS 解析
dig example.com                    # DNS 查询
host example.com                   # DNS 查找

curl -v telnet://10.0.0.1:22      # 测试 TCP 连通性
nc -zv 10.0.0.1 443               # 端口连通测试
nc -zv 10.0.0.1 22 80 443         # 多端口测试

pathmtu 10.0.0.1                  # 路径 MTU 发现
ip route get 8.8.8.8              # 路由查找

sysctl net.ipv4.tcp_keepalive_time
sysctl net.ipv4.tcp_keepalive_intvl
sysctl net.ipv4.tcp_keepalive_probes
```
