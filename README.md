# kubesw [![CI](https://github.com/spideyz0r/kubesw/workflows/gotester/badge.svg)][![CI](https://github.com/spideyz0r/kubesw/workflows/goreleaser/badge.svg)][![CI](https://github.com/spideyz0r/kubesw/workflows/rpm-builder/badge.svg)] [[![Go Report Card](https://goreportcard.com/badge/github.com/spideyz0r/kubesw)](https://goreportcard.com/report/github.com/spideyz0r/kubesw)]
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
$ kubesw list context
gke_us-central
k3s
minikube
```
```
$ kubesw set context minikube
$ kubesw get context
minikube
```
```
$ kubesw list namespace
default
kube-node-lease
kube-public
kube-system
```
```
$ kubesw set namespace kube-system
$ kubesw get namespace
kube-system
```
## Aliases / shortnames
You can also use the short form for each command:
```
namespace => namespaces ns
context   => contexts ctx
list      => ls l
get       => current g
set       => switch s
```
Some examples:
```kubesw get ns
kubesw set ctx "somecontext"
kubesw ls ns
kubesw switch ns "somenamespace"
```

## Auto-completion
The autocompletion script can be generated with the following:
```
kubesw completion bash
```

You can add the following line to your ~/.bashrc:
```
source <(`which kubesw` completion bash)
```

## TODO
- Read extra/optional configurations like rc files and PS1 from the configuration file (maybe use viper)
- Add support to zsh
- Improve error checks and messages
- Add --global flag for updating the namespace or context globally
- Investigate the use of eval instead of spawning shells
