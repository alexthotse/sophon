---
sidebar_position: 1
sidebar_label: Install
---

# Install Sophon

## Quick Install

```bash
curl -sL https://sophon.ai/install.sh | bash
```

## Manual install

Grab the appropriate binary for your platform from the latest [release](https://github.com/plandex-ai/sophon/releases) and put it somewhere in your `PATH`.

## Build from source

```bash
git clone https://github.com/plandex-ai/sophon.git
cd sophon/app/cli
go build -ldflags "-X sophon/version.Version=$(cat version.txt)"
mv sophon /usr/local/bin # adapt as needed for your system
```

## Windows

Windows is supported via [WSL](https://learn.microsoft.com/en-us/windows/wsl/about).

Sophon only works correctly in the WSL shell. It doesn't work in the Windows CMD prompt or PowerShell.