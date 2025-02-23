FROM ubuntu:22.04

# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive

# Install basic tools
RUN apt-get update && apt-get install -y \
    git \
    curl \
    wget \
    gcc \
    zsh \
    bash-completion \
    && rm -rf /var/lib/apt/lists/*

# Install Go
RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz \
    && rm go1.21.6.linux-amd64.tar.gz

# Set up environment
ENV PATH=$PATH:/usr/local/go/bin:/go/bin
ENV GOPATH=/go

# Create user
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Create workspace and Go directories
RUN mkdir -p /workspace /go/bin /go/src \
    && chown -R $USERNAME:$USERNAME /workspace /go

# Configure bash settings
RUN echo 'export HISTSIZE=10000' >> /home/$USERNAME/.bashrc \
    && echo 'export HISTFILESIZE=20000' >> /home/$USERNAME/.bashrc \
    && echo 'source /etc/bash_completion' >> /home/$USERNAME/.bashrc \
    && echo 'bind "set show-all-if-ambiguous on"' >> /home/$USERNAME/.bashrc \
    && echo 'bind "set completion-ignore-case on"' >> /home/$USERNAME/.bashrc

USER $USERNAME
WORKDIR /workspace
