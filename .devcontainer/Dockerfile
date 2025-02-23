FROM ubuntu:22.04

# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive

# Install Go and other utilities
RUN apt-get update && apt-get install -y \
    wget \
    git \
    sudo \
    vim \
    curl \
    htop \
    jq \
    ripgrep \
    zfsutils-linux \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install Go
RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz \
    && rm go1.21.6.linux-amd64.tar.gz

# Add Go to path
ENV PATH=$PATH:/usr/local/go/bin

# Install Go tools
RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

# Create the user
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user and add sudo support
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Create Go workspace directory and set permissions
RUN mkdir -p /go/src /go/bin \
    && chown -R $USERNAME:$USERNAME /go

# [Optional] Set the default user
USER $USERNAME

# Set up shell customization
COPY --chown=vscode:vscode shell/.bashrc /home/vscode/.bashrc
COPY --chown=vscode:vscode shell/.profile /home/vscode/.profile
COPY --chown=vscode:vscode shell/.inputrc /home/vscode/.inputrc

# Set up Go environment
ENV GOPATH=/go
ENV PATH=$PATH:/go/bin

# Create and set proper ownership of GOPATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Set working directory
WORKDIR /workspace
