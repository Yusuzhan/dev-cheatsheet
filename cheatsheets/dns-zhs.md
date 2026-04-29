---
title: DNS
icon: fa-globe
primary: "#FF6B35"
lang: bash
locale: zhs
---

## fa-list DNS 记录类型

```text
A       地址记录                   192.0.2.1
AAAA    IPv6 地址记录              2001:db8::1
CNAME   别名记录                   alias.example.com → example.com
MX      邮件交换                   10 mail.example.com.
NS      名称服务器                 ns1.example.com.
TXT     文本记录                   "v=spf1 include:_spf.example.com ~all"
SOA     起始授权                   主 NS、管理员邮箱、序列号、TTL
SRV     服务记录                   10 60 5060 sip.example.com.
PTR     指针记录（反向 DNS）        1.2.0.192.in-addr.arpa → host.example.com
CAA     证书授权                   0 issue "letsencrypt.org"
NS      委派                       ns1.example.com.
```

## fa-magnifying-glass dig 命令

```bash
dig example.com                    # A 记录
dig example.com A                  # 显式查询 A 记录
dig example.com AAAA               # IPv6
dig example.com MX                 # 邮件服务器
dig example.com NS                 # 名称服务器
dig example.com TXT                # TXT 记录
dig example.com ANY                # 所有记录
dig -x 192.0.2.1                   # 反向解析
dig example.com +short             # 简短输出
dig example.com +trace             # 追踪委派路径
dig @8.8.8.8 example.com           # 指定 DNS 服务器
dig example.com CNAME +noall +answer
```

## fa-search nslookup

```bash
nslookup example.com               # 基本查询
nslookup example.com 8.8.8.8       # 指定 DNS 服务器
nslookup -type=MX example.com      # MX 记录
nslookup -type=NS example.com      # NS 记录
nslookup -type=TXT example.com     # TXT 记录
nslookup -type=SOA example.com     # SOA 记录
nslookup -type=ANY example.com     # 所有记录
nslookup 192.0.2.1                 # 反向解析
nslookup -debug example.com        # 详细输出
```

## fa-route DNS 解析流程

```text
浏览器缓存 → 操作系统缓存 → 本地解析器 → 递归解析器

递归解析过程：
1. 客户端向递归解析器发起请求
2. 解析器查询根服务器 (.) → 返回 TLD 服务器
3. 解析器查询 TLD 服务器 (.com) → 返回权威服务器
4. 解析器查询权威服务器 → 返回最终结果
5. 解析器缓存结果（TTL），返回给客户端

根服务器：a.root-servers.net（13 个逻辑节点，数百个任播）
TLD：    .com, .org, .net, 国家代码 (.cn, .jp)
权威：   ns1.example.com（区域所有者）
```

```bash
dig . NS                           # 根服务器
dig com. NS                        # .com TLD 服务器
dig example.com @a.gtld-servers.net
dig example.com +trace            # 完整委派路径
```

## fa-file-lines 区域文件

```text
$ORIGIN example.com.
$TTL 86400

@       IN  SOA  ns1.example.com. admin.example.com. (
            2025010101  ; 序列号
            3600        ; 刷新
            900         ; 重试
            604800      ; 过期
            86400       ; 最小 TTL
        )

@       IN  NS   ns1.example.com.
@       IN  NS   ns2.example.com.
@       IN  A    192.0.2.10
@       IN  AAAA 2001:db8::10
@       IN  MX   10 mail.example.com.
@       IN  TXT  "v=spf1 ip4:192.0.2.0/24 ~all"

www     IN  CNAME @
api     IN  A    192.0.2.20
_sip._tcp  IN  SRV  10 60 5060 sip.example.com.
```

## fa-clock TTL

```text
TTL（生存时间）：记录可被缓存的秒数

常见 TTL 值：
60       — 1 分钟     （故障切换、测试）
300      — 5 分钟     （频繁变更）
3600     — 1 小时     （常规使用）
86400    — 1 天       （稳定记录）
604800   — 1 周       （静态：NS、根）
```

```bash
dig example.com +noall +answer
# example.com.  300  IN  A  192.0.2.1
#                ^^^ TTL（秒）

dig example.com +dnssec | grep "flags:"
```

## fa-rotate 反向 DNS (PTR)

```text
IPv4:  192.0.2.1 → 1.2.0.192.in-addr.arpa
IPv6:  2001:db8::1 → 1.0.0.0...8.b.d.0.1.0.0.2.ip6.arpa
```

```bash
dig -x 192.0.2.1                   # 反向解析
dig -x 2001:db8::1                 # IPv6 反向解析
nslookup 192.0.2.1
host 192.0.2.1

dig 1.2.0.192.in-addr.arpa PTR
```

```text
; 反向区域文件：2.0.192.in-addr.arpa
$ORIGIN 2.0.192.in-addr.arpa.
1       IN  PTR  web.example.com.
10      IN  PTR  mail.example.com.
```

## fa-link CNAME 链

```text
www.example.com  CNAME  web.example.com
web.example.com  CNAME  lb.example.com
lb.example.com   A      192.0.2.10

规则：
- CNAME 不能指向同名的另一个 CNAME
- CNAME 不能与同一名称的其他记录共存
- 域名根 (@) 不能使用 CNAME（使用 ALIAS/ANAME）
- 最大链深度：取决于解析器（通常 8-10）
```

```bash
dig www.example.com CNAME
dig www.example.com +trace        # 追踪完整链
dig www.example.com A +noall +answer
```

## fa-database DNS 缓存

```bash
resolvectl status                  # systemd-resolved 状态
resolvectl query example.com       # 查询（含缓存信息）
resolvectl flush-caches            # 清空本地缓存
resolvectl statistics              # 缓存命中统计

rndc flush                         # 清空 BIND 缓存
rndc stats                         # BIND 统计

systemd-resolve --statistics       # 旧版 systemd
ipconfig /flushdns                 # Windows
dscacheutil -flushcache            # macOS
```

```text
缓存层级：
1. 浏览器 DNS 缓存       chrome://net-internals/#dns
2. 操作系统解析器         /etc/resolv.conf → systemd-resolved
3. 递归解析器             ISP 或公共（8.8.8.8, 1.1.1.1）
4. 权威 TTL              由区域所有者设置
```

## fa-envelope MX 记录与邮件

```text
example.com.  IN  MX  10 mail1.example.com.
example.com.  IN  MX  20 mail2.example.com.
example.com.  IN  MX  30 backup.mail.provider.com.

优先级：数字越小优先级越高
```

```bash
dig example.com MX +short
nslookup -type=MX example.com
host -t MX example.com

dig _submission._tcp.example.com SRV
```

## fa-file-shield TXT 与 SPF/DKIM/DMARC

```text
SPF（发件人策略框架）：
"v=spf1 ip4:192.0.2.0/24 include:_spf.google.com ~all"
  +all    通过（危险）
  ~all    软拒绝
  -all    硬拒绝
  ?all    中立
```

```bash
dig example.com TXT +short
dig _dmarc.example.com TXT +short
dig selector._domainkey.example.com TXT +short
```

```text
DMARC：
"v=DMARC1; p=reject; rua=mailto:dmarc@example.com; pct=100"

DKIM：
"v=DKIM1; k=rsa; p=MIGfMA0GCSqGSIb3DQEBA..."
```

## fa-lock DNSSEC

```text
DNSSEC 信任链：
根 → TLD (.com) → 权威 (example.com)

记录类型：
DNSKEY    区域公钥
RRSIG     记录集签名
DS        委派签名者（父 → 子链接）
NSEC/NSEC3  不存在证明
```

```bash
dig example.com A +dnssec
dig example.com DNSKEY
dig example.com DS @a.gtld-servers.net
dnsviz.net example.com              # 在线可视化
delv example.com                    # 验证解析工具

dig +dnssec @8.8.8.8 example.com A | grep RRSIG
```

## fa-arrows-rotate 动态 DNS

```bash
nsupdate -l                         # 本地动态更新
nsupdate
> server ns1.example.com
> zone example.com
> update add test.example.com. 300 A 192.0.2.50
> send
> quit

nsupdate
> server ns1.example.com
> zone example.com
> update delete test.example.com. A
> send

ddclient -daemon=300                # DDNS 客户端（动态 IP）
```

## fa-wrench 故障排查

```bash
dig example.com +trace             # 追踪委派路径
dig example.com @8.8.8.8           # 查询 Google DNS
dig example.com @1.1.1.1           # 查询 Cloudflare DNS
dig example.com +dnssec            # 检查 DNSSEC
dig example.com +nodnssec +comments

host -v example.com                 # 详细输出
whois example.com                   # 域名注册信息

systemd-resolve --status            # 解析器配置
cat /etc/resolv.conf                # nameserver 配置

tcpdump -i any -nn port 53          # 抓取 DNS 流量
tcpdump -i any -nn port 53 -w dns.pcap

named-checkzone example.com zonefile.db    # 验证区域文件
named-checkconf /etc/named.conf            # 验证配置文件
```
