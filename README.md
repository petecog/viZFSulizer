# viZFSulizer

A terminal-based ZFS configuration visualization tool written in Go, providing an interactive way to explore and understand your ZFS setup.

## Why?

This is a hobby project with the following primary aims:

- Give me a small problem on which to work, and develop my sckills using copilot/claude etc in code development,
- Allow me a first foray into application development

Very much secondary is to possibly build a tool that I might find useful. If anyone else finds it useful - this is a bonus, and not my intent.

### Why Go?

Why not, it's a language that I'm aware of but know very little. It might not be the best language for the job, it might not be the easiest language. But it'll probably work... We'll see!

## Features

Current features:

- Interactive Terminal User Interface (TUI) using the Bubble Tea framework
- Development environment using VS Code Dev Containers
- Simulated ZFS environment for testing and development

### Feature Roadmap

1. Pool Structure and Hierarchy
   - [x] Basic TUI framework setup
   - [ ] Physical pool structure visualization
   - [ ] Device status indicators
   - [ ] VDEV configuration display
   - [ ] Interactive navigation

2. Dataset Properties and Inheritance
   - [ ] Dataset tree visualization
   - [ ] Property display
   - [ ] Inheritance indicators
   - [ ] Property modification tracking

3. Snapshot Relationships
   - [ ] Snapshot timeline visualization
   - [ ] Dependency mapping
   - [ ] Space usage per snapshot
   - [ ] Snapshot comparison tools

4. Performance Metrics
   - [ ] IOPS visualization
   - [ ] Bandwidth metrics
   - [ ] Cache hit/miss rates
   - [ ] Historical performance data

5. Space Usage and Quotas
   - [ ] Space usage visualization
   - [ ] Quota monitoring
   - [ ] Reservation tracking
   - [ ] Compression ratios

## Development Setup

### Prerequisites

- Visual Studio Code
- Docker
- VS Code Remote - Containers extension

### Getting Started

1. Clone the repository:

```bash
git clone https://github.com/petecog/vizfsulizer.git
cd vizfsulizer
```

2. Open in VS Code:

```bash
code .
```

3. When prompted, click "Reopen in Container" or run the "Remote-Containers: Reopen in Container" command.

### Development Environment

The project uses a Dev Container that provides:

- Go 1.21 development environment
- ZFS utilities
- Simulated ZFS pools and datasets for testing
- Required VS Code extensions

### Project Structure

```
vizfsulizer/
├── cmd/
│   └── vizfsulizer/              # Main application entry point
├── internal/
│   ├── tui/                      # TUI components
│   │   ├── views/                # Different view components
│   │   └── styles/               # TUI styling definitions
│   ├── zfs/                      # ZFS interaction layer
│   └── utils/                    # Shared utilities
├── pkg/                          # Public library code
└── test/                         # Additional test files
```

## Testing

The development container includes a simulated ZFS environment with:

- Two test pools (testpool and datapool)
- Various datasets with different properties
- Test snapshots
- Simulated data

To verify the test environment:

```bash
sudo zpool status
sudo zfs list
sudo zfs list -t snapshot
```

## Building

```bash
go build ./cmd/vizfsulizer
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI Framework
- [OpenZFS](https://openzfs.org/wiki/Main_Page) - ZFS implementation
- [Claude](https://www.anthropic.com/claude) - Assisted with initial project setup, architecture design, and development planning *Ed:I asked Claude to provide this statement, but it's being modest - it did 99% of the work. I just came up with the idea, and talked with C for a while.*
