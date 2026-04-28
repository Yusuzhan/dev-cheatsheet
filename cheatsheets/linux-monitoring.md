---
title: Linux Monitoring
icon: fa-chart-line
primary: "#FCC624"
lang: bash
---

## fa-microchip CPU & Process (top/htop)

```bash
top                                      # interactive process viewer
htop                                     # better top (needs install)
top -u nginx                             # filter by user
top -p 1234,5678                         # monitor specific PIDs

# top interactive keys
# P     sort by CPU usage
# M     sort by memory usage
# k     kill a process
# z     color highlight
# q     quit

mpstat 1 5                               # CPU stats every 1s, 5 times
uptime                                   # load averages (1min, 5min, 15min)
nproc                                    # number of CPU cores
```

## fa-list Process Management (ps)

```bash
ps aux                                   # all processes, BSD style
ps -ef                                   # all processes, System V style
ps aux --sort=-%mem                      # sort by memory usage (desc)
ps aux --sort=-%cpu                      # sort by CPU usage (desc)
ps -u nginx                              # processes by user
ps -p 1234 -o pid,ppid,cmd,%mem,%cpu     # custom format for a PID
ps aux | grep nginx                      # find process by name
ps -ef --forest                          # process tree view
pgrep -f "nginx"                         # get PID by pattern
pidof nginx                              # get PID by exact name
```

## fa-memory Memory (free)

```bash
free -h                                  # human-readable memory info
free -m                                  # in megabytes
free -s 2                                # refresh every 2 seconds
vmstat 1 5                               # virtual memory stats
```

```bash
#  quick summary
#  total   total installed memory
#  used    used memory
#  free    unused memory
#  shared  shared memory (tmpfs, IPC)
#  buff/cache  kernel buffers + page cache
#  available   memory available for new processes
```

## fa-hard-drive Disk Usage (df/du)

```bash
df -h                                    # disk free, human-readable
df -T                                    # show filesystem type
df -i                                    # show inode usage
du -sh /var/log/*                        # directory sizes, summarized
du -h --max-depth=1 /var                 # one level deep
du -sh /                                 # total disk usage
du -ah /home | sort -rh | head -20       # top 20 largest files/dirs
lsblk                                    # list block devices
fdisk -l                                 # disk partitions
iostat -x 1                              # extended I/O stats
iotop                                    # I/O usage by process (needs root)
```

## fa-network-wired Network (ss/netstat)

```bash
ss -tlnp                                 # listening TCP ports with process
ss -tlnp | grep :80                      # check specific port
ss -s                                    # connection statistics summary
ss -tn                                   # established TCP connections
ss -un                                   # UDP connections
ss -tn state established '( dport = :443 or sport = :443 )'  # filter by port

# netstat (legacy, ss is preferred)
netstat -tlnp                            # listening TCP ports
netstat -anp                             # all connections with PIDs
netstat -rn                              # routing table
```

## fa-door-open Open Files (lsof)

```bash
lsof                                     # all open files
lsof -i :80                              # who is using port 80
lsof -i :80,443                          # multiple ports
lsof -u nginx                            # files opened by user
lsof -p 1234                             # files opened by PID
lsof /var/log/syslog                     # who opened this file
lsof -i tcp                              # all TCP connections
lsof -i udp                              # all UDP connections
lsof +D /var/log                         # files opened under directory
```

## fa-gear System Services (systemctl)

```bash
systemctl list-units --type=service      # list all services
systemctl list-units --type=service --state=running   # running services
systemctl status nginx                   # service status
systemctl start nginx                    # start service
systemctl stop nginx                     # stop service
systemctl restart nginx                  # restart service
systemctl reload nginx                   # reload config without restart
systemctl enable nginx                   # start on boot
systemctl disable nginx                  # don't start on boot
systemctl is-active nginx                # check if running
systemctl is-enabled nginx               # check if enabled on boot
journalctl -u nginx                      # service logs
journalctl -u nginx -f                   # follow service logs
journalctl -u nginx --since "1 hour ago" # recent logs
journalctl --since "2025-04-29 10:00"    # logs since timestamp
```

## fa-fire Network Firewall (iptables/ufw)

```bash
# ufw (Ubuntu default)
ufw status                               # show firewall status
ufw allow 80/tcp                         # allow port 80
ufw allow 443/tcp                        # allow port 443
ufw allow from 10.0.0.0/8               # allow subnet
ufw deny 3306                            # deny port
ufw delete allow 8080                    # remove rule

# iptables
iptables -L -n -v                        # list all rules
iptables -L -n --line-numbers            # with rule numbers
iptables -A INPUT -p tcp --dport 80 -j ACCEPT   # allow port 80
iptables -D INPUT 3                      # delete rule by number
```

## fa-clock Scheduled Tasks

```bash
crontab -l                               # list current user's crontab
crontab -e                               # edit crontab
crontab -l -u nginx                      # list another user's crontab

#  ┌──────── minute (0-59)
#  │ ┌────── hour (0-23)
#  │ │ ┌──── day of month (1-31)
#  │ │ │ ┌── month (1-12)
#  │ │ │ │ ┌ day of week (0-7, 0 and 7 = Sunday)
#  │ │ │ │ │
#  * * * * * command

#  */5 * * * *     /opt/healthcheck.sh        # every 5 minutes
#  0 2 * * *      /opt/backup.sh              # daily at 2 AM
#  0 0 1 * *      /opt/monthly-report.sh      # first day of month
#  0 9 * * 1-5    /opt/weekday-job.sh          # weekdays at 9 AM

systemctl list-timers                    # list systemd timers
timedatectl                              # timezone & NTP status
```

## fa-user User & Login

```bash
who                                      # logged-in users
w                                        # logged-in users with activity
whoami                                   # current username
id                                       # user ID and groups
last                                     # login history
lastb                                    # failed login attempts
lastlog                                  # last login for all users

# user management
useradd -m -s /bin/bash newuser          # create user with home dir
userdel -r olduser                       # delete user and home dir
passwd username                          # change password
groups username                          # show user groups
usermod -aG docker username              # add user to group
```

## fa-lightbulb Quick Health Check

```bash
# one-liner system overview
echo "=== CPU ===" && nproc && uptime
echo "=== Memory ===" && free -h
echo "=== Disk ===" && df -h /
echo "=== Top CPU ===" && ps aux --sort=-%cpu | head -6
echo "=== Top Mem ===" && ps aux --sort=-%mem | head -6
echo "=== Listening ===" && ss -tlnp
echo "=== Failed Services ===" && systemctl --failed

# find largest files
du -ah / 2>/dev/null | sort -rh | head -20

# find deleted files still held open
lsof +L1

# check recent reboots
last reboot | head -5

# kernel messages (hardware issues)
dmesg -T | grep -i error
```
