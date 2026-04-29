---
title: strace / ltrace
icon: fa-bug
primary: "#E74C3C"
lang: bash
---

## fa-terminal Basic strace

```bash
strace ls
strace -o trace.log ./myapp
strace -e trace=read,write ./myapp
strace -c ./myapp                    # summary with counts
strace -tt ./myapp                   # timestamps
strace -T ./myapp                    # time spent in syscalls
```

## fa-filter Filter by Syscall

```bash
strace -e trace=open,openat ./myapp
strace -e trace=network ./myapp
strace -e trace=file ./myapp
strace -e trace=process ./myapp
strace -e trace=signal ./myapp
strace -e trace=memory ./myapp
strace -e trace=read,write ./myapp
strace -e trace=%net ./myapp
strace -e trace=\!close ./myapp       # exclude close
```

## fa-folder-open Trace File Operations

```bash
strace -e trace=open,openat,read,write,close ./myapp
strace -e trace=stat,lstat,access ./myapp
strace -e trace=mkdir,rmdir,unlink ./myapp
strace -e trace=chmod,chown ./myapp
strace -e trace=rename,link,symlink ./myapp
strace -y ./myapp                    # show file descriptor paths
strace -yy ./myapp                   # show ip/port for socket fds
```

## fa-network-wired Trace Network

```bash
strace -e trace=socket,bind,listen,accept,connect ./myapp
strace -e trace=sendto,recvfrom,sendmsg,recvmsg ./myapp
strace -e trace=getsockopt,setsockopt ./myapp
strace -e trace=%network ./myapp
strace -f -e trace=network ./myapp   # follow forks
strace -yy -e trace=network ./myapp  # show addresses
```

## fa-bolt Trace Signals

```bash
strace -e trace=signal ./myapp
strace -e signal=SIGSEGV ./myapp
strace -e signal=SIGINT,SIGTERM ./myapp
strace -e signal=ALL ./myapp
strace -e trace=kill,tkill,tgkill ./myapp
```

## fa-clock Timing & Statistics

```bash
strace -c ./myapp                    # count/time summary
strace -c -S time ./myapp            # sort by time
strace -c -S calls ./myapp           # sort by call count
strace -c -S errors ./myapp          # sort by errors
strace -T ./myapp                    # time per syscall
strace -tt ./myapp                   # timestamp per line
strace -ttt ./myapp                  # epoch timestamp
strace -r ./myapp                    # relative timestamp
strace -w ./myapp                    # summary only
```

## fa-link Attach to Process (-p)

```bash
strace -p 1234                       # attach to PID
strace -p 1234 -p 5678              # attach to multiple PIDs
strace -p 1234 -e trace=network     # filter while attached
strace -p 1234 -o trace.log         # attach and log
strace -p 1234 -f                   # attach + follow forks
sudo strace -p $(pgrep -f nginx)    # attach by name
strace -p 1234 -s 200               # increase string size
```

## fa-code-branch Follow Forks (-f)

```bash
strace -f ./myapp                    # trace children
strace -f -o trace.log ./myapp      # log with fork info
strace -f -e trace=execve ./myapp   # trace exec calls
strace -f -ff -o trace ./myapp      # separate file per PID
strace -f -e trace=clone ./myapp    # trace thread creation
```

## fa-file-lines strace Output Format

```bash
# Typical output line
# openat(AT_FDCWD, "/etc/hosts", O_RDONLY) = 3
# read(3, "127.0.0.1\tlocalhost\n", 4096) = 20
# close(3)                          = 0

# Error output
# openat(AT_FDCWD, "/noexist", O_RDONLY) = -1 ENOENT (No such file or directory)

# With -T (time in syscall)
# openat(AT_FDCWD, "/etc/hosts", O_RDONLY) = 3 <0.000123>

# With -y (file paths)
# read(3</etc/hosts>, "127.0.0.1\n", 4096) = 14

# With -tt (timestamps)
# 14:23:01.123456 openat(AT_FDCWD, "/etc/hosts", O_RDONLY) = 3

# Verbose struct decoding
strace -v -e trace=getdents64 ./myapp
strace -e trace=mmap -v ./myapp
```

## fa-magnifying-glass ltrace Basics

```bash
ltrace ./myapp                       # trace library calls
ltrace -o ltrace.log ./myapp
ltrace -c ./myapp                    # summary counts
ltrace -T ./myapp                    # time per call
ltrace -tt ./myapp                   # timestamps
ltrace -r ./myapp                    # relative timestamps
ltrace -p 1234                       # attach to PID
ltrace -f ./myapp                    # follow forks
ltrace -n 2 ./myapp                  # indent nested calls
ltrace -S ./myapp                    # also show syscalls
```

## fa-sliders ltrace Filter

```bash
ltrace -e malloc ./myapp             # trace malloc only
ltrace -e malloc+free ./myapp       # trace malloc and free
ltrace -e '*printf' ./myapp         # glob pattern
ltrace -e '\!strlen' ./myapp        # exclude strlen
ltrace -e 'printf+scanf' ./myapp   # trace printf and scanf
ltrace -e '*' -S ./myapp            # all library + syscalls
```

## fa-lightbulb Common Use Cases

```bash
# Find which config files an app reads
strace -e trace=openat -f ./myapp 2>&1 | grep '\.conf\|\.yaml\|\.json'

# Find why a binary can't find a library
strace -e trace=openat ./myapp 2>&1 | grep 'No such file'

# Check DNS resolution
strace -e trace=sendto,recvfrom -yy wget https://example.com

# Find file descriptor leaks
strace -e trace=openat,close -c ./myapp

# Monitor memory allocation
ltrace -e malloc+free+realloc -c ./myapp

# Debug permission denied
strace -e trace=openat,access,stat -f ./myapp 2>&1 | grep EACCES
```

## fa-gauge Performance Overhead

```bash
# strace adds ~2-10x overhead (syscalls only)
# ltrace adds ~5-20x overhead (library calls)

# Minimize overhead with filters
strace -e trace=openat ./myapp       # filter reduces overhead
strace -c ./myapp                    # summary mode is lighter

# Compare timing
time ./myapp
time strace -e trace=openat ./myapp

# Use perf for production-safe tracing
perf stat ./myapp
perf record -g ./myapp
perf report
```

## fa-wrench Debugging Recipes

```bash
# Find hanging process
strace -p $PID -e trace=all -T

# Find slow syscalls (>10ms)
strace -T ./myapp 2>&1 | grep '<0\.0*[1-9]'

# Trace specific thread
strace -p $(pgrep -f "worker" | head -1) -f

# Find leaked file descriptors
strace -e trace=openat,close -o /dev/stdout ./myapp | \
  awk '/openat/{fd=$NF} /close\(.*'"$fd"'\)/{fd=""} END{if(fd)print "leak:"fd}'

# Reconstruct stdout from trace
strace -s 65535 -e trace=write ./myapp 2>&1 | \
  grep write\(1 | sed 's/.*"\(.*\)".*/\1/'

# Compare two binary behaviors
diff <(strace -c ./app_v1 2>&1) <(strace -c ./app_v2 2>&1)
```
