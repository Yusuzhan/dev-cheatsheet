---
title: TCP/IP
icon: fa-network-wired
primary: "#0078D4"
lang: bash
---

## fa-layer-group OSI vs TCP/IP Model

```text
OSI Model              TCP/IP Model          Protocol Examples
7 Application          Application           HTTP, FTP, DNS, SSH, SMTP
6 Presentation         Application           TLS, JPEG, JSON, XML
5 Session              Application           NetBIOS, RPC, PPTP
4 Transport            Transport             TCP, UDP, SCTP, QUIC
3 Network              Internet              IP, ICMP, IGMP, ARP
2 Data Link            Link Access           Ethernet, Wi-Fi, PPP, VLAN
1 Physical             Link Access           Fiber, Cat6, Radio
```

```text
Encapsulation (send):  Data → Segment → Packet → Frame → Bits
Decapsulation (recv):  Bits → Frame → Packet → Segment → Data
```

## fa-location-dot IP Addressing (v4/v6)

```bash
ip addr show
ip addr show eth0
ip addr add 192.168.1.10/24 dev eth0
ip addr del 192.168.1.10/24 dev eth0
```

```text
IPv4 Classes:
A: 1.0.0.0     – 126.255.255.255   /8    (large networks)
B: 128.0.0.0   – 191.255.255.255   /16   (medium networks)
C: 192.0.0.0   – 223.255.255.255   /24   (small networks)

Private Ranges:
10.0.0.0/8          172.16.0.0/12        192.168.0.0/16
127.0.0.0/8 (loopback)    169.254.0.0/16 (link-local)
```

```text
IPv6:
2001:0db8:85a3:0000:0000:8a2e:0370:7334
2001:db8:85a3::8a2e:370:7334          # compressed
::1                                    # loopback
fe80::                                 # link-local
ff02::                                 # multicast
```

```bash
ip -6 addr show
ip -6 route show
ping6 ::1
```

## fa-scissors Subnetting & CIDR

```text
CIDR        Subnet Mask          Hosts        Usable
/8          255.0.0.0            16,777,214
/16         255.255.0.0          65,534
/20         255.255.240.0        4,094
/24         255.255.255.0        254
/25         255.255.255.128      126
/26         255.255.255.192      62
/27         255.255.255.224      30
/28         255.255.255.240      14
/30         255.255.255.252      2
/32         255.255.255.255      1 (host route)
```

```bash
ipcalc 192.168.1.0/24
ipcalc 10.0.0.0/20
sipcalc 172.16.0.0/12

ip route add 10.10.0.0/16 via 192.168.1.1
ip route add default via 192.168.1.1
ip route show
```

## fa-handshake TCP Handshake

```text
3-Way Handshake (Connection):

Client                          Server
  |--- SYN (seq=x) ------------->|     1. Client sends SYN
  |<-- SYN-ACK (seq=y,ack=x+1) -|     2. Server sends SYN-ACK
  |--- ACK (ack=y+1) ----------->|     3. Client sends ACK
  |        ESTABLISHED            |

4-Way Handshake (Termination):

Client                          Server
  |--- FIN (seq=u) ------------->|
  |<-- ACK (ack=u+1) -----------|
  |<-- FIN (seq=w) -------------|
  |--- ACK (ack=w+1) ----------->|
  |        CLOSED                |
```

```bash
tcpdump -i eth0 'tcp[tcpflags] & (tcp-syn|tcp-fin) != 0'
ss -tn state established
ss -tn state time-wait
```

## fa-circle-nodes TCP States

```text
CLOSED → SYN_SENT → ESTABLISHED → FIN_WAIT_1 → FIN_WAIT_2 → TIME_WAIT → CLOSED
LISTEN → SYN_RCVD → ESTABLISHED → CLOSE_WAIT → LAST_ACK → CLOSED

Key States:
LISTEN        Server waiting for connections
SYN_SENT      Client sent SYN, waiting for SYN-ACK
SYN_RCVD      Server received SYN, sent SYN-ACK
ESTABLISHED   Connection active, data transfer
FIN_WAIT_1    Sent FIN, waiting for ACK
FIN_WAIT_2    Received ACK for FIN, waiting for remote FIN
CLOSE_WAIT    Received FIN, waiting for local close
LAST_ACK      Sent FIN after CLOSE_WAIT, waiting for ACK
TIME_WAIT     Active close, waiting 2*MSL before cleanup
```

```bash
ss -tna                           # all TCP states
ss -tn state time-wait            # TIME_WAIT sockets
ss -tn state close-wait           # CLOSE_WAIT sockets
cat /proc/sys/net/ipv4/tcp_fin_timeout
cat /proc/sys/net/ipv4/ip_local_port_range
```

## fa-bolt UDP

```text
TCP vs UDP:
TCP: connection-oriented, reliable, ordered, flow control, congestion control
UDP: connectionless, no guarantee, no order, low overhead, low latency

Use UDP for: DNS, DHCP, TFTP, SNMP, NTP, VoIP, video streaming, gaming
Use TCP for: HTTP, SSH, FTP, SMTP, database connections
```

```bash
nc -u -l 12345                    # UDP listener
nc -u 192.168.1.10 12345          # UDP sender
echo "hello" | nc -u -w1 10.0.0.1 514

nmap -sU -p 53,67,68,123 10.0.0.1    # UDP port scan
```

## fa-plug Common Ports

```text
FTP Data      20/TCP     FTP Control   21/TCP
SSH           22/TCP     Telnet        23/TCP
SMTP          25/TCP     DNS           53/TCP/UDP
DHCP Server   67/UDP     DHCP Client   68/UDP
TFTP          69/UDP     HTTP          80/TCP
POP3          110/TCP    NTP           123/UDP
NetBIOS       137-139    IMAP          143/TCP
SNMP          161/UDP    HTTPS         443/TCP
Syslog        514/UDP    SMTPS         465/TCP
LDAPS         636/TCP    iSCSI         860/TCP
MySQL         3306/TCP   RDP           3389/TCP
PostgreSQL    5432/TCP   Redis         6379/TCP
HTTP Alt      8080/TCP   HTTPS Alt     8443/TCP
```

## fa-satellite-dish ping / traceroute

```bash
ping -c 4 8.8.8.8                 # 4 pings
ping -i 0.5 10.0.0.1              # interval 0.5s
ping -W 2 10.0.0.1                # timeout 2s
ping -s 1400 -M do 10.0.0.1      # MTU probe with DF bit

traceroute 8.8.8.8                # UDP-based
traceroute -T -p 443 example.com  # TCP-based
mtr -rwzc 50 example.com          # continuous traceroute
tracepath 8.8.8.8                 # PMTU discovery

ping6 -c 4 2001:4860:4860::8888
traceroute6 2001:4860:4860::8888
```

## fa-table-list netstat / ss

```bash
ss -tlnp                           # TCP listening ports with process
ss -tnp                            # TCP established connections
ss -tunlp                          # TCP + UDP listening
ss -s                              # socket summary
ss -tn state established '( dport = :443 or sport = :443 )'
ss -tn dst 192.168.1.0/24         # filter by destination

netstat -tlnp                      # legacy: TCP listening
netstat -tulnp                     # legacy: TCP + UDP listening
netstat -tn                        # legacy: TCP established
netstat -s                         # legacy: statistics
netstat -rn                        # legacy: routing table
```

## fa-magnifying-glass tcpdump

```bash
tcpdump -i eth0                    # all traffic on eth0
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
nmap -sT 10.0.0.1                  # TCP connect scan
nmap -sS 10.0.0.1                  # TCP SYN (stealth) scan
nmap -sU 10.0.0.1                  # UDP scan
nmap -sV -sT 10.0.0.1              # version detection
nmap -O 10.0.0.1                   # OS detection
nmap -sn 192.168.1.0/24            # host discovery (ping sweep)
nmap -p 22,80,443 10.0.0.1         # specific ports
nmap -p- 10.0.0.1                  # all 65535 ports
nmap -A 10.0.0.1                   # aggressive (OS + version + script)
nmap -sT -sV --script vuln 10.0.0.1
```

## fa-shield-halved iptables

```bash
iptables -L -n -v                              # list all rules
iptables -A INPUT -p tcp --dport 22 -j ACCEPT  # allow SSH
iptables -A INPUT -p tcp --dport 80 -j ACCEPT  # allow HTTP
iptables -A INPUT -p tcp --dport 443 -j ACCEPT # allow HTTPS
iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
iptables -A INPUT -i lo -j ACCEPT              # allow loopback
iptables -P INPUT DROP                          # default deny
iptables -A INPUT -s 10.0.0.0/8 -j ACCEPT     # allow from subnet

iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 8080

iptables -D INPUT 3                            # delete rule #3
iptables -F                                    # flush all rules
iptables-save > /etc/iptables/rules.v4
```

## fa-right-left NAT & PAT

```text
NAT Types:
Static NAT     1:1 mapping    192.168.1.10 ↔ 203.0.113.10
Dynamic NAT    Pool mapping   Internal hosts share public IP pool
PAT (SNAT)     N:1 mapping    Many internals → 1 public IP + ports

SNAT (Source NAT):     Change source address (outbound)
DNAT (Destination NAT): Change destination (port forwarding)
MASQUERADE:             SNAT for dynamic IPs
```

```bash
ip addr show                     # view addresses
iptables -t nat -L -n -v        # view NAT rules

iptables -t nat -A POSTROUTING -s 192.168.1.0/24 -o eth0 -j MASQUERADE
iptables -t nat -A PREROUTING -p tcp -d 203.0.113.10 --dport 80 -j DNAT --to-destination 192.168.1.20:8080

cat /proc/sys/net/ipv4/ip_forward   # check forwarding
sysctl -w net.ipv4.ip_forward=1     # enable forwarding
```

## fa-wrench Troubleshooting

```bash
ip link show                       # interface status
ethtool eth0                       # link speed, duplex, errors
ethtool -S eth0                    # interface statistics
mii-tool eth0                      # link status (legacy)

arp -an                            # ARP table
ip neigh show                      # neighbor table
arping -I eth0 192.168.1.1         # ARP ping

nslookup example.com               # DNS resolution
dig example.com                    # DNS query
host example.com                   # DNS lookup

curl -v telnet://10.0.0.1:22      # test TCP connectivity
nc -zv 10.0.0.1 443               # port connectivity test
nc -zv 10.0.0.1 22 80 443         # multi-port test

pathmtu 10.0.0.1                  # path MTU discovery
ip route get 8.8.8.8              # route lookup

sysctl net.ipv4.tcp_keepalive_time
sysctl net.ipv4.tcp_keepalive_intvl
sysctl net.ipv4.tcp_keepalive_probes
```
