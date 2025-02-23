# Color Mode Implementation

## Overview

Implement black & white mode for accessibility, switchable via command line flag.

## Requirements

- Command line flag: `--color=bw` or `--color=rgb` (default)
- Different border styles for status indication:
  - ONLINE: Normal borders `─│╭╮╰╯`
  - DEGRADED: Dashed borders `╌╌╌╌`
  - FAULTED: Double-line borders `═║╔╗╚╝`

## Implementation Steps

1. Add config package:

   ```go
   type DisplayMode string
   const (
       DisplayModeRGB DisplayMode = "rgb"
       DisplayModeBW DisplayMode = "bw"
   )
   ```

2. Modify theme.go:
   - Add BW border configurations
   - Add mode-aware styling functions
   - Create separate border sets for each status

3. Update main.go:
   - Add flag parsing
   - Pass display mode to model

4. Files to modify:
   - `/internal/tui/styles/theme.go`
   - `/cmd/vizfsulizer/main.go`
   - `/internal/tui/model.go`
   - Create `/internal/config/display.go`

## Border Characters Reference

```
Normal:   ─ │ ╭ ╮ ╰ ╯
Dashed:   ╌ ┊ ┌ ┐ └ ┘
Double:   ═ ║ ╔ ╗ ╚ ╝
```

## Notes

- Consider users with color blindness
- Ensure clear visual distinction between states
- Document border meanings in help text
- Add configuration persistence option for future
