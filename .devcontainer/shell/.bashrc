# Add bash-specific customizations here
export PS1='\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '
export HISTSIZE=10000
export HISTFILESIZE=20000
export HISTCONTROL=ignoreboth

# Enable bash completion
if [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
fi

# Useful aliases
alias ll='ls -alF'
alias la='ls -A'
alias l='ls -CF'
alias ..='cd ..'
alias ...='cd ../..'

# Go development aliases
alias gt='go test ./...'
alias gr='go run .'
alias gm='go mod tidy'

# Git aliases
alias gs='git status'
alias gl='git log --oneline'
alias gd='git diff'

