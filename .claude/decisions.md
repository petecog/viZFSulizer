# Architecture and Implementation Decisions

## DevContainer Strategy

### Decision: Use Dockerfile approach over devcontainer features
**Rationale**: More control over package versions and installation order. Node.js 20.x from NodeSource provides better Claude CLI compatibility than Ubuntu's older nodejs package.

### Decision: Bash over Zsh for default shell
**Rationale**: Simpler configuration for persistent history. Bash is more predictable across different environments and better supported for scripting.

### Decision: Skip ZFS setup in container environment
**Rationale**: Containers don't have access to loop devices by default. ZFS testing should be done on host system or dedicated test environment where proper kernel modules and permissions are available.

## Data Persistence Strategy

### Decision: Use bind mounts for persistent development data
**Rationale**: 
- `.local_dev_stuff` mount provides persistent bash history and development files
- Host `.claude` mount enables seamless Claude CLI integration with existing memory
- More reliable than volume mounts for development workflows

### Decision: Git structure for .local_dev_stuff
**Rationale**: Use `.gitkeep` to preserve directory structure while ignoring contents. Allows team members to have the directory but keeps personal development data private.

## Development Tooling

### Decision: Install Claude CLI in container
**Rationale**: Provides AI assistance directly in development environment. Host `.claude` mount ensures continuity with host Claude usage and memory.

### Decision: Maintain Go toolchain in Dockerfile
**Rationale**: Ensures consistent Go versions across development environments. Installing in Dockerfile is faster than in postCreateCommand for frequently used tools.

## Environment Management

### Decision: Use .env file support
**Rationale**: Provides standard way to manage environment-specific configuration while keeping secrets out of git. `.env.example` serves as documentation.