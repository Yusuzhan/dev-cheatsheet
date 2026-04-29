---
title: strace / ltrace
icon: fa-bug
primary: "#E74C3C"
lang: bash
locale: zhs
---

## fa-terminal 基础 strace

```bash
strace ls
strace -o trace.log ./myapp
strace -e trace=read,write ./myapp
strace -c ./myapp                    # 统计汇总
strace -tt ./myapp                   # 时间戳
strace -T ./myapp                    # 系统调用耗时
```

## fa-filter 按系统调用过滤

```bash
strace -e trace=open,openat ./myapp
strace -e trace=network ./myapp
strace -e trace=file ./myapp
strace -e trace=process ./myapp
strace -e trace=signal ./myapp
strace -e trace=memory ./myapp
strace -e trace=read,write ./myapp
strace -e trace=%net ./myapp
strace -e trace=\!close ./myapp       # 排除 close
```

## fa-folder-open 跟踪文件操作

```bash
strace -e trace=open,openat,read,write,close ./myapp
strace -e trace=stat,lstat,access ./myapp
strace -e trace=mkdir,rmdir,unlink ./myapp
strace -e trace=chmod,chown ./myapp
strace -e trace=rename,link,symlink ./myapp
strace -y ./myapp                    # 显示文件描述符路径
strace -yy ./myapp                   # 显示 socket 的 IP/端口
```

## fa-network-wired 跟踪网络

```bash
strace -e trace=socket,bind,listen,accept,connect ./myapp
strace -e trace=sendto,recvfrom,sendmsg,recvmsg ./myapp
strace -e trace=getsockopt,setsockopt ./myapp
strace -e trace=%network ./myapp
strace -f -e trace=network ./myapp   # 跟踪子进程
strace -yy -e trace=network ./myapp  # 显示地址
```

## fa-bolt 跟踪信号

```bash
strace -e trace=signal ./myapp
strace -e signal=SIGSEGV ./myapp
strace -e signal=SIGINT,SIGTERM ./myapp
strace -e signal=ALL ./myapp
strace -e trace=kill,tkill,tgkill ./myapp
```

## fa-clock 计时与统计

```bash
strace -c ./myapp                    # 计数/时间汇总
strace -c -S time ./myapp            # 按耗时排序
strace -c -S calls ./myapp           # 按调用次数排序
strace -c -S errors ./myapp          # 按错误数排序
strace -T ./myapp                    # 每次系统调用耗时
strace -tt ./myapp                   # 每行时间戳
strace -ttt ./myapp                  # Unix 时间戳
strace -r ./myapp                    # 相对时间戳
strace -w ./myapp                    # 仅汇总
```

## fa-link 附加到进程 (-p)

```bash
strace -p 1234                       # 附加到 PID
strace -p 1234 -p 5678              # 附加到多个 PID
strace -p 1234 -e trace=network     # 附加时过滤
strace -p 1234 -o trace.log         # 附加并记录日志
strace -p 1234 -f                   # 附加 + 跟踪子进程
sudo strace -p $(pgrep -f nginx)    # 按名称附加
strace -p 1234 -s 200               # 增加字符串显示长度
```

## fa-code-branch 跟踪子进程 (-f)

```bash
strace -f ./myapp                    # 跟踪子进程
strace -f -o trace.log ./myapp      # 记录含子进程信息
strace -f -e trace=execve ./myapp   # 跟踪 exec 调用
strace -f -ff -o trace ./myapp      # 每个 PID 单独文件
strace -f -e trace=clone ./myapp    # 跟踪线程创建
```

## fa-file-lines strace 输出格式

```bash
# 典型输出行
# openat(AT_FDCWD, "/etc/hosts", O_RDONLY) = 3
# read(3, "127.0.0.1\tlocalhost\n", 4096) = 20
# close(3)                          = 0

# 错误输出
# openat(AT_FDCWD, "/noexist", O_RDONLY) = -1 ENOENT (No such file or directory)

# 带 -T (系统调用耗时)
# openat(AT_FDCWD, "/etc/hosts", O_RDONLY) = 3 <0.000123>

# 带 -y (文件路径)
# read(3</etc/hosts>, "127.0.0.1\n", 4096) = 14

# 带 -tt (时间戳)
# 14:23:01.123456 openat(AT_FDCWD, "/etc/hosts", O_RDONLY) = 3

# 详细结构体解码
strace -v -e trace=getdents64 ./myapp
strace -e trace=mmap -v ./myapp
```

## fa-magnifying-glass ltrace 基础

```bash
ltrace ./myapp                       # 跟踪库函数调用
ltrace -o ltrace.log ./myapp
ltrace -c ./myapp                    # 统计汇总
ltrace -T ./myapp                    # 每次调用耗时
ltrace -tt ./myapp                   # 时间戳
ltrace -r ./myapp                    # 相对时间戳
ltrace -p 1234                       # 附加到 PID
ltrace -f ./myapp                    # 跟踪子进程
ltrace -n 2 ./myapp                  # 嵌套调用缩进
ltrace -S ./myapp                    # 同时显示系统调用
```

## fa-sliders ltrace 过滤

```bash
ltrace -e malloc ./myapp             # 仅跟踪 malloc
ltrace -e malloc+free ./myapp       # 跟踪 malloc 和 free
ltrace -e '*printf' ./myapp         # glob 模式匹配
ltrace -e '\!strlen' ./myapp        # 排除 strlen
ltrace -e 'printf+scanf' ./myapp   # 跟踪 printf 和 scanf
ltrace -e '*' -S ./myapp            # 所有库调用 + 系统调用
```

## fa-lightbulb 常见用例

```bash
# 查找应用读取哪些配置文件
strace -e trace=openat -f ./myapp 2>&1 | grep '\.conf\|\.yaml\|\.json'

# 查找二进制找不到库的原因
strace -e trace=openat ./myapp 2>&1 | grep 'No such file'

# 检查 DNS 解析
strace -e trace=sendto,recvfrom -yy wget https://example.com

# 发现文件描述符泄漏
strace -e trace=openat,close -c ./myapp

# 监控内存分配
ltrace -e malloc+free+realloc -c ./myapp

# 调试权限拒绝
strace -e trace=openat,access,stat -f ./myapp 2>&1 | grep EACCES
```

## fa-gauge 性能开销

```bash
# strace 增加约 2-10 倍开销 (仅系统调用)
# ltrace 增加约 5-20 倍开销 (库函数调用)

# 使用过滤器减少开销
strace -e trace=openat ./myapp       # 过滤减少开销
strace -c ./myapp                    # 汇总模式更轻量

# 对比耗时
time ./myapp
time strace -e trace=openat ./myapp

# 生产环境使用 perf 安全追踪
perf stat ./myapp
perf record -g ./myapp
perf report
```

## fa-wrench 调试技巧

```bash
# 查找挂起进程
strace -p $PID -e trace=all -T

# 查找慢系统调用 (>10ms)
strace -T ./myapp 2>&1 | grep '<0\.0*[1-9]'

# 跟踪特定线程
strace -p $(pgrep -f "worker" | head -1) -f

# 查找泄漏的文件描述符
strace -e trace=openat,close -o /dev/stdout ./myapp | \
  awk '/openat/{fd=$NF} /close\(.*'"$fd"'\)/{fd=""} END{if(fd)print "leak:"fd}'

# 从跟踪中还原 stdout
strace -s 65535 -e trace=write ./myapp 2>&1 | \
  grep write\(1 | sed 's/.*"\(.*\)".*/\1/'

# 对比两个二进制的差异
diff <(strace -c ./app_v1 2>&1) <(strace -c ./app_v2 2>&1)
```
