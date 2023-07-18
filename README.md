# kubesw [![CI](https://github.com/spideyz0r/kubesw/workflows/gotester/badge.svg)][![CI](https://github.com/spideyz0r/kubesw/workflows/goreleaser/badge.svg)][![CI](https://github.com/spideyz0r/kubesw/workflows/rpm-builder/badge.svg)]
Kubeswitch is a dynamic tool designed to enhance your kubernetes workflow by enabling isolated context and namespace management for each terminal

## Install

### RPM
```
dnf copr enable brandfbb/kubesw
dnf install kubesw
```

### From source
```
go build -v -o kubesw

### Download the binary from the release section

## Usage
```
# kubesw --help
Kubeswitch is a dynamic tool designed to enhance your kubernetes workflow by enabling isolated context and namespace management for each terminal

Usage:
  kubesw [flags]
  kubesw [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Get current context or namespace
  help        Help about any command
  list        List all contexts or namespaces
  set         Set context or namespace

Flags:
  -h, --help   help for kubesw

Use "kubesw [command] --help" for more information about a command.
```