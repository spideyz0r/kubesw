# kubesw [![CI](https://github.com/spideyz0r/kubesw/workflows/gotester/badge.svg)][![CI](https://github.com/spideyz0r/kubesw/workflows/goreleaser/badge.svg)][![CI](https://github.com/spideyz0r/kubesw/workflows/rpm-builder/badge.svg)]
`kubesw` is a versatile option for context switching, namespace switching, and prompt customization. It ensures that each shell operates independently, serving as an alternative to tools like Kubectx or Kubens.

It's a dynamic tool designed to enhance your Kubernetes workflow by enabling isolated context and namespace management for each terminal.

## Install

### RPM
```
dnf copr enable brandfbb/kubesw
dnf install kubesw
```

### From source
```
go build -v -o kubesw
```
### Download the binary from the release section
```
https://github.com/spideyz0r/kubesw/releases
```

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

## Examples
```
kubesw get context
kubesw list context
kubesw set context minikube
kubesw get namespace
kubesw list namespace
kubesw set namespace kube-system
kubesw version
```
## TODO
- Fix the debug flag inside common
- Read extra/optional configurations like rc files and PS1 from the configuration file (maybe use viper)
- Add support to zsh
- Remove namespace string when listing namespaces
- Improve error checks and messages
- Add --global flag for updating the namespace or context globally
- Move history along with the new shell instance (if possible)
- Add shell autocompletion
- Allow short names, ctx for context, and ns for namespace
- Investigate the use of eval instead of spawning shells
