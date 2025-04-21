<div align="center">
  <img src="assets/logo.png" width="450">
</div>

<div align="center">
  <p>Make better commits.</p>
</div>

---

## What it do?

Comet is a tool written in Golang that helps you make better commits following conventions. It uses [bubbletea](https://github.com/charmbracelet/bubbletea) TUI framework to guide you to create conventional commits commands.

## Installation

### Install Comet

```bash
go install github.com/fastiraz/comet@latest
```

### Build from source

```bash
git clone --depth 1 https://github.com/Fastiraz/comet.git
cd comet
go build
```

Run from build directory by executing it directly from the current directory:

```bash
./comet
```

If you want to install the binary on your system, you can move the binary in your `/bin`, `/usr/bin`, `~/.local/bin` or `/usr/local/bin`.

```bash
mv comet /usr/local/bin
```

### Verify installation

```bash
comet
```
