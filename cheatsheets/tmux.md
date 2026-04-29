---
title: tmux
icon: fa-window-restore
primary: "#1BB91F"
lang: bash
---

## fa-terminal Sessions

```bash
tmux new -s dev                    # new named session
tmux new -s project -d             # new detached session
tmux ls                            # list sessions
tmux attach -t dev                 # attach to session
tmux attach -t dev -d              # attach and detach others
tmux kill-session -t dev           # kill session
tmux rename-session -t dev work    # rename session
tmux switch-client -t dev          # switch to session
```

## fa-window-maximize Windows

```bash
tmux new-window -n logs            # new named window
tmux new-window -n build -c ~/src  # new window in directory
Ctrl+b c                           # create window
Ctrl+b n                           # next window
Ctrl+b p                           # previous window
Ctrl+b 0-9                         # window by number
Ctrl+b &                           # kill window
Ctrl+b ,                           # rename window
Ctrl+b w                           # list windows
```

## fa-columns Panes

```bash
Ctrl+b "                           # split horizontal
Ctrl+b %                           # split vertical
Ctrl+b o                           # cycle panes
Ctrl+b ↑↓←→                        # navigate panes
Ctrl+b x                           # kill pane
Ctrl+b z                           # toggle zoom
Ctrl+b {                           # swap pane up
Ctrl+b }                           # swap pane down
Ctrl+b !                           # pane to window
Ctrl+b q                           # show pane numbers
```

## fa-copy Copy Mode

```bash
Ctrl+b [                           # enter copy mode
Ctrl+b ]                           # paste buffer
Ctrl+b #                           # list buffers
Ctrl+b =                           # choose buffer to paste

# In copy mode (vi keys):
Space                               # start selection
Enter                               # copy selection
q                                   # exit copy mode
/                                   # search forward
?                                   # search backward
g                                   # top of buffer
G                                   # bottom of buffer
```

## fa-file-code Config (~/.tmux.conf)

```bash
set -g prefix C-a                  # change prefix to Ctrl+a
set -g base-index 1                # windows start at 1
setw -g pane-base-index 1          # panes start at 1
set -g mouse on                    # enable mouse
set -g default-terminal "tmux-256color"
set -g history-limit 10000
set -g renumber-windows on
set -g display-time 2000
setw -g mode-keys vi               # vi key bindings
bind | split-window -h -c "#{pane_current_path}"
bind - split-window -v -c "#{pane_current_path}"
bind r source-file ~/.tmux.conf \; display "Reloaded!"
```

## fa-keyboard Key Bindings

```bash
bind-key -T prefix "'" command-prompt -p "window:" "select-window -t '%%'"
bind-key -T prefix m set mouse \; display "mouse: #{mouse}"
bind-key -T copy-mode-vi v send-keys -X begin-selection
bind-key -T copy-mode-vi y send-keys -X copy-selection-and-cancel

# Bind custom commands
bind-key -T prefix S run "tmux split-window -l 40 'tmux list-sessions'"
bind-key -T prefix P run "tmux capture-pane -p -S -50 | fzf"

# List all bindings
tmux list-keys
```

## fa-table-columns Layouts

```bash
# Main-horizontal layout
select-layout main-horizontal

# Predefined layouts
Ctrl+b M-1                         # even-horizontal
Ctrl+b M-2                         # even-vertical
Ctrl+b M-3                         # main-horizontal
Ctrl+b M-4                         # main-vertical
Ctrl+b M-5                         # tiled

# Custom layout with break-pane
tmux split-window -h -p 30
tmux split-window -v -p 50
tmux select-layout main-vertical
```

## fa-file-lines Scripting (tmuxinator)

```bash
# tmuxinator project config (~/.tmuxinator/dev.yml)
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

tmuxinator start dev               # start project
tmuxinator stop dev                # stop project
tmuxinator list                    # list projects
tmuxinator edit dev                # edit config

# Plain tmux scripting
tmux new-session -d -s work -c ~/src
tmux send-keys -t work 'vim' C-m
tmux split-window -v -t work
tmux send-keys -t work 'npm test' C-m
tmux attach -t work
```

## fa-rotate Resurrect / Continuum

```bash
# tmux-resurrect (prefix + Ctrl+s / Ctrl+r)
set -g @plugin 'tmux-plugins/tmux-resurrect'
set -g @resurrect-capture-pane-contents 'on'
set -g @resurrect-processes ':all:'

# tmux-continuum (auto-save/restore)
set -g @plugin 'tmux-plugins/tmux-continuum'
set -g @continuum-save-interval '15'
set -g @continuum-restore 'on'

# Manual resurrect
Ctrl+b Ctrl+s                      # save session
Ctrl+b Ctrl+r                      # restore session

# Restore neovim sessions
set -g @resurrect-strategy-nvim 'session'
```

## fa-computer-mouse Mouse Support

```bash
set -g mouse on                    # enable all mouse

# Mouse settings breakdown
set -g mouse-select-pane on
set -g mouse-resize-pane on
set -g mouse-select-window on

# Scroll with mouse in copy mode
bind -T root WheelUpPane \
  if-shell -F "#{alternate_on}" \
    "send-keys -M" \
    "if-shell -F '#{pane_in_mode}' 'send-keys -M' 'select-pane -t= ; copy-mode -e ; send-keys -M'"

# Toggle mouse on/off
bind m set -g mouse \; display "mouse: #{mouse}"
```

## fa-bars Status Bar

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

## fa-network-wired Remote Sessions

```bash
# SSH into tmux on remote host
ssh -t user@host "tmux attach -t dev || tmux new -s dev"

# Persistent SSH tmux session
ssh -t production 'tmux new -A -s main'

# Auto-attach on SSH login (add to .bashrc)
if [ -z "$TMUX" ] && [ -n "$SSH_TTY" ]; then
  tmux attach-session -t ssh || tmux new-session -s ssh
  exit
fi

# Mosh + tmux for unreliable connections
mosh user@host -- tmux attach -d
```

## fa-layer-group Synchronized Panes

```bash
Ctrl+b :                          # enter command mode
:setw synchronize-panes on        # enable sync
:setw synchronize-panes off       # disable sync

# Bind toggle
bind S setw synchronize-panes

# Broadcast commands to all panes
# Useful for running same command on multiple servers
# 1. Open multiple SSH panes
# 2. Enable sync
# 3. Type once, executed on all panes
```

## fa-plug Useful Plugins

```bash
# TPM (Tmux Plugin Manager)
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'tmux-plugins/tmux-sensible'
set -g @plugin 'tmux-plugins/tmux-pain-control'
set -g @plugin 'tmux-plugins/tmux-yank'
set -g @plugin 'tmux-plugins/tmux-open'
set -g @plugin 'tmux-plugins/tmux-copycat'
set -g @plugin 'tmux-plugins/tmux-resurrect'
set -g @plugin 'tmux-plugins/tmux-continuum'
set -g @plugin 'tmux-plugins/tmux-sessionist'

# Install TPM and plugins
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm
~/.tmux/plugins/tpm/bin/install_plugins

# TPM key bindings
Ctrl+b I                           # install plugins
Ctrl+b U                           # update plugins
Ctrl+b M-u                         # uninstall plugins
```
