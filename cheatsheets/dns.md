---
title: DNS
icon: fa-globe
primary: "#FF6B35"
lang: bash
---

## fa-list DNS Record Types

```text
A       Address                    192.0.2.1
AAAA    IPv6 Address               2001:db8::1
CNAME   Canonical Name             alias.example.com → example.com
MX      Mail Exchange              10 mail.example.com.
NS      Name Server                ns1.example.com.
TXT     Text                       "v=spf1 include:_spf.example.com ~all"
SOA     Start of Authority         primary ns, admin email, serial, TTLs
SRV     Service                    10 60 5060 sip.example.com.
PTR     Pointer (reverse DNS)      1.2.0.192.in-addr.arpa → host.example.com
CAA     Certification Authority    0 issue "letsencrypt.org"
NS      Delegation                 ns1.example.com.
```

## fa-magnifying-glass dig Command

```bash
dig example.com                    # A record
dig example.com A                  # explicit A record
dig example.com AAAA               # IPv6
dig example.com MX                 # mail servers
dig example.com NS                 # name servers
dig example.com TXT                # TXT records
dig example.com ANY                # all records
dig -x 192.0.2.1                   # reverse lookup
dig example.com +short             # short answer
dig example.com +trace             # trace delegation path
dig @8.8.8.8 example.com           # query specific server
dig example.com CNAME +noall +answer
```

## fa-search nslookup

```bash
nslookup example.com               # basic lookup
nslookup example.com 8.8.8.8       # use specific DNS server
nslookup -type=MX example.com      # MX records
nslookup -type=NS example.com      # NS records
nslookup -type=TXT example.com     # TXT records
nslookup -type=SOA example.com     # SOA record
nslookup -type=ANY example.com     # all records
nslookup 192.0.2.1                 # reverse lookup
nslookup -debug example.com        # verbose output
```

## fa-route DNS Resolution Flow

```text
Browser Cache → OS Cache → Resolver (Stub) → Recursive Resolver

Recursive Resolution:
1. Client asks recursive resolver
2. Resolver queries Root server (.) → TLD referral
3. Resolver queries TLD server (.com) → authoritative referral
4. Resolver queries Authoritative server → final answer
5. Resolver caches answer (TTL), returns to client

Root:     a.root-servers.net  (13 logical, hundreds anycast)
TLD:      .com, .org, .net, country codes (.cn, .jp)
Auth:     ns1.example.com     (zone owner)
```

```bash
dig . NS                           # root servers
dig com. NS                        # .com TLD servers
dig example.com @a.gtld-servers.net
dig example.com +trace            # full delegation path
```

## fa-file-lines Zone Files

```text
$ORIGIN example.com.
$TTL 86400

@       IN  SOA  ns1.example.com. admin.example.com. (
            2025010101  ; serial
            3600        ; refresh
            900         ; retry
            604800      ; expire
            86400       ; minimum TTL
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
TTL (Time To Live): seconds a record may be cached

Common TTLs:
60       — 1 minute    (failover, testing)
300      — 5 minutes   (frequent changes)
3600     — 1 hour      (typical)
86400    — 1 day       (stable records)
604800   — 1 week      (static: NS, root)
```

```bash
dig example.com +noall +answer
# example.com.  300  IN  A  192.0.2.1
#                ^^^ TTL in seconds

dig example.com +dnssec | grep "flags:"
# check if TTL is respected by resolvers
```

## fa-rotate Reverse DNS (PTR)

```text
IPv4:  192.0.2.1 → 1.2.0.192.in-addr.arpa
IPv6:  2001:db8::1 → 1.0.0.0...8.b.d.0.1.0.0.2.ip6.arpa
```

```bash
dig -x 192.0.2.1                   # reverse lookup
dig -x 2001:db8::1                 # IPv6 reverse
nslookup 192.0.2.1
host 192.0.2.1

dig 1.2.0.192.in-addr.arpa PTR
```

```text
; Reverse zone file: 2.0.192.in-addr.arpa
$ORIGIN 2.0.192.in-addr.arpa.
1       IN  PTR  web.example.com.
10      IN  PTR  mail.example.com.
```

## fa-link CNAME Chaining

```text
www.example.com  CNAME  web.example.com
web.example.com  CNAME  lb.example.com
lb.example.com   A      192.0.2.10

Rules:
- CNAME cannot point to another CNAME at the same name
- CNAME cannot coexist with other data at the same name
- Apex (@) cannot be CNAME (use ALIAS/ANAME at provider)
- Max chain depth: resolver-dependent (typically 8-10)
```

```bash
dig www.example.com CNAME
dig www.example.com +trace        # follow full chain
dig www.example.com A +noall +answer
```

## fa-database DNS Caching

```bash
resolvectl status                  # systemd-resolved status
resolvectl query example.com       # query with cache info
resolvectl flush-caches            # flush local cache
resolvectl statistics              # cache hit/miss stats

rndc flush                         # flush BIND cache
rndc stats                         # BIND statistics

systemd-resolve --statistics       # older systemd
ipconfig /flushdns                 # Windows
dscacheutil -flushcache            # macOS
```

```text
Cache Layers:
1. Browser DNS cache      chrome://net-internals/#dns
2. OS stub resolver       /etc/resolv.conf → systemd-resolved
3. Recursive resolver     ISP or public (8.8.8.8, 1.1.1.1)
4. Authoritative TTL      set by zone owner
```

## fa-envelope MX Records & Email

```text
example.com.  IN  MX  10 mail1.example.com.
example.com.  IN  MX  20 mail2.example.com.
example.com.  IN  MX  30 backup.mail.provider.com.

Priority: lower number = higher priority
```

```bash
dig example.com MX +short
nslookup -type=MX example.com
host -t MX example.com

dig _submission._tcp.example.com SRV
```

## fa-file-shield TXT & SPF/DKIM/DMARC

```text
SPF (Sender Policy Framework):
"v=spf1 ip4:192.0.2.0/24 include:_spf.google.com ~all"
  +all    pass (dangerous)
  ~all    softfail
  -all    hardfail
  ?all    neutral
```

```bash
dig example.com TXT +short
dig _dmarc.example.com TXT +short
dig selector._domainkey.example.com TXT +short
```

```text
DMARC:
"v=DMARC1; p=reject; rua=mailto:dmarc@example.com; pct=100"

DKIM:
"v=DKIM1; k=rsa; p=MIGfMA0GCSqGSIb3DQEBA..."
```

## fa-lock DNSSEC

```text
DNSSEC Chain of Trust:
Root → TLD (.com) → Authoritative (example.com)

Record Types:
DNSKEY    Public key for the zone
RRSIG     Signature on record sets
DS        Delegation Signer (parent → child link)
NSEC/NSEC3  Proof of non-existence
```

```bash
dig example.com A +dnssec
dig example.com DNSKEY
dig example.com DS @a.gtld-servers.net
dnsviz.net example.com              # online visualization
delv example.com                    # validating resolver tool

dig +dnssec @8.8.8.8 example.com A | grep RRSIG
```

## fa-arrows-rotate Dynamic DNS

```bash
nsupdate -l                         # local dynamic update
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

ddclient -daemon=300                # DDNS client for dynamic IPs
```

## fa-wrench Troubleshooting

```bash
dig example.com +trace             # trace delegation path
dig example.com @8.8.8.8           # query Google DNS
dig example.com @1.1.1.1           # query Cloudflare DNS
dig example.com +dnssec            # check DNSSEC
dig example.com +nodnssec +comments

host -v example.com                 # verbose output
whois example.com                   # domain registration info

systemd-resolve --status            # resolver config
cat /etc/resolv.conf                # nameserver config

tcpdump -i any -nn port 53          # capture DNS traffic
tcpdump -i any -nn port 53 -w dns.pcap

named-checkzone example.com zonefile.db    # validate zone
named-checkconf /etc/named.conf            # validate config
```
