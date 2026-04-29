---
title: tmux
icon: fa-window-restore
primary: "#1BB91F"
lang: bash
locale: zhs
---

## fa-terminal 会话

```bash
tmux new -s dev                    # 新建命名会话
tmux new -s project -d             # 新建分离会话
tmux ls                            # 列出会话
tmux attach -t dev                 # 连接会话
tmux attach -t dev -d              # 连接并分离其他客户端
tmux kill-session -t dev           # 终止会话
tmux rename-session -t dev work    # 重命名会话
tmux switch-client -t dev          # 切换到会话
```

## fa-window-maximize 窗口

```bash
tmux new-window -n logs            # 新建命名窗口
tmux new-window -n build -c ~/src  # 在指定目录新建窗口
Ctrl+b c                           # 创建窗口
Ctrl+b n                           # 下一个窗口
Ctrl+b p                           # 上一个窗口
Ctrl+b 0-9                         # 按编号切换窗口
Ctrl+b &                           # 关闭窗口
Ctrl+b ,                           # 重命名窗口
Ctrl+b w                           # 列出窗口
```

## fa-columns 面板

```bash
Ctrl+b "                           # 水平分割
Ctrl+b %                           # 垂直分割
Ctrl+b o                           # 循环切换面板
Ctrl+b ↑↓←→                        # 导航面板
Ctrl+b x                           # 关闭面板
Ctrl+b z                           # 切换缩放
Ctrl+b {                           # 面板上移
Ctrl+b }                           # 面板下移
Ctrl+b !                           # 面板转为窗口
Ctrl+b q                           # 显示面板编号
```

## fa-copy 复制模式

```bash
Ctrl+b [                           # 进入复制模式
Ctrl+b ]                           # 粘贴缓冲区
Ctrl+b #                           # 列出缓冲区
Ctrl+b =                           # 选择缓冲区粘贴

# 复制模式中 (vi 键位):
Space                               # 开始选择
Enter                               # 复制选中内容
q                                   # 退出复制模式
/                                   # 向前搜索
?                                   # 向后搜索
g                                   # 缓冲区顶部
G                                   # 缓冲区底部
```

## fa-file-code 配置 (~/.tmux.conf)

```bash
set -g prefix C-a                  # 前缀键改为 Ctrl+a
set -g base-index 1                # 窗口从 1 开始编号
setw -g pane-base-index 1          # 面板从 1 开始编号
set -g mouse on                    # 启用鼠标
set -g default-terminal "tmux-256color"
set -g history-limit 10000
set -g renumber-windows on
set -g display-time 2000
setw -g mode-keys vi               # vi 键绑定
bind | split-window -h -c "#{pane_current_path}"
bind - split-window -v -c "#{pane_current_path}"
bind r source-file ~/.tmux.conf \; display "已重新加载!"
```

## fa-keyboard 键绑定

```bash
bind-key -T prefix "'" command-prompt -p "window:" "select-window -t '%%'"
bind-key -T prefix m set mouse \; display "mouse: #{mouse}"
bind-key -T copy-mode-vi v send-keys -X begin-selection
bind-key -T copy-mode-vi y send-keys -X copy-selection-and-cancel

# 绑定自定义命令
bind-key -T prefix S run "tmux split-window -l 40 'tmux list-sessions'"
bind-key -T prefix P run "tmux capture-pane -p -S -50 | fzf"

# 列出所有绑定
tmux list-keys
```

## fa-table-columns 布局

```bash
# 主水平布局
select-layout main-horizontal

# 预定义布局
Ctrl+b M-1                         # 均匀水平
Ctrl+b M-2                         # 均匀垂直
Ctrl+b M-3                         # 主水平
Ctrl+b M-4                         # 主垂直
Ctrl+b M-5                         # 平铺

# 自定义布局
tmux split-window -h -p 30
tmux split-window -v -p 50
tmux select-layout main-vertical
```

## fa-file-lines 脚本化 (tmuxinator)

```bash
# tmuxinator 项目配置 (~/.tmuxinator/dev.yml)
cat << 'EOF'
name: dev
root: ~/projects/myapp
windows:
  - editor: vim
  - server: npm run dev
  - logs: tail -f logs/development.log
  - database: pgcli mydb
  - git: tig
EOF

tmuxinator start dev               # 启动项目
tmuxinator stop dev                # 停止项目
tmuxinator list                    # 列出项目
tmuxinator edit dev                # 编辑配置

# 纯 tmux 脚本
tmux new-session -d -s work -c ~/src
tmux send-keys -t work 'vim' C-m
tmux split-window -v -t work
tmux send-keys -t work 'npm test' C-m
tmux attach -t work
```

## fa-rotate 会话恢复 (Resurrect / Continuum)

```bash
# tmux-resurrect (前缀 + Ctrl+s / Ctrl+r)
set -g @plugin 'tmux-plugins/tmux-resurrect'
set -g @resurrect-capture-pane-contents 'on'
set -g @resurrect-processes ':all:'

# tmux-continuum (自动保存/恢复)
set -g @plugin 'tmux-plugins/tmux-continuum'
set -g @continuum-save-interval '15'
set -g @continuum-restore 'on'

# 手动恢复
Ctrl+b Ctrl+s                      # 保存会话
Ctrl+b Ctrl+r                      # 恢复会话

# 恢复 neovim 会话
set -g @resurrect-strategy-nvim 'session'
```

## fa-computer-mouse 鼠标支持

```bash
set -g mouse on                    # 启用所有鼠标功能

# 鼠标设置分解
set -g mouse-select-pane on
set -g mouse-resize-pane on
set -g mouse-select-window on

# 鼠标滚轮在复制模式中滚动
bind -T root WheelUpPane \
  if-shell -F "#{alternate_on}" \
    "send-keys -M" \
    "if-shell -F '#{pane_in_mode}' 'send-keys -M' 'select-pane -t= ; copy-mode -e ; send-keys -M'"

# 切换鼠标开/关
bind m set -g mouse \; display "mouse: #{mouse}"
```

## fa-bars 状态栏

```bash
set -g status-position top
set -g status-interval 5
set -g status-style "bg=#1a1b26,fg=#a9b1d6"

set -g status-left "#[fg=#7aa2f7]#S #[fg=#565f89]| "
set -g status-left-length 40
set -g status-right " %H:%M %Y-%m-%d"
set -g status-right-length 50

set -g window-status-format " #I:#W "
set -g window-status-current-format "#[fg=#1a1b26,bg=#7aa2f7] #I:#W "
set -g pane-border-style "fg=#3b4261"
set -g pane-active-border-style "fg=#7aa2f7"
set -g message-style "bg=#1a1b26,fg=#7aa2f7"
```

## fa-network-wired 远程会话

```bash
# SSH 连接到远程 tmux
ssh -t user@host "tmux attach -t dev || tmux new -s dev"

# 持久 SSH tmux 会话
ssh -t production 'tmux new -A -s main'

# SSH 登录自动连接 (添加到 .bashrc)
if [ -z "$TMUX" ] && [ -n "$SSH_TTY" ]; then
  tmux attach-session -t ssh || tmux new-session -s ssh
  exit
fi

# Mosh + tmux 处理不稳定连接
mosh user@host -- tmux attach -d
```

## fa-layer-group 面板同步

```bash
Ctrl+b :                          # 进入命令模式
:setw synchronize-panes on        # 启用同步
:setw synchronize-panes off       # 禁用同步

# 绑定切换
bind S setw synchronize-panes

# 向所有面板广播命令
# 适用于在多台服务器运行相同命令
# 1. 打开多个 SSH 面板
# 2. 启用同步
# 3. 输入一次, 在所有面板执行
```

## fa-plug 实用插件

```bash
# TPM (Tmux 插件管理器)
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'tmux-plugins/tmux-sensible'
set -g @plugin 'tmux-plugins/tmux-pain-control'
set -g @plugin 'tmux-plugins/tmux-yank'
set -g @plugin 'tmux-plugins/tmux-open'
set -g @plugin 'tmux-plugins/tmux-copycat'
set -g @plugin 'tmux-plugins/tmux-resurrect'
set -g @plugin 'tmux-plugins/tmux-continuum'
set -g @plugin 'tmux-plugins/tmux-sessionist'

# 安装 TPM 和插件
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm
~/.tmux/plugins/tpm/bin/install_plugins

# TPM 快捷键
Ctrl+b I                           # 安装插件
Ctrl+b U                           # 更新插件
Ctrl+b M-u                         # 卸载插件
```
