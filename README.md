# kubesw [![CI](https://github.com/spideyz0r/kubesw/workflows/gotester/badge.svg)][![CI](https://github.com/spideyz0r/kubesw/workflows/goreleaser/badge.svg)][![CI](https://github.com/spideyz0r/kubesw/workflows/rpm-builder/badge.svg)] [[![Go Report Card](https://goreportcard.com/badge/github.com/spideyz0r/kubesw)](https://goreportcard.com/report/github.com/spideyz0r/kubesw)]
`kubesw` is a versatile option for context switching, namespace switching, and prompt customization. It ensures that each shell operates independently, serving as an alternative to tools like Kubectx or Kubens.

It's a dynamic tool designed to enhance your Kubernetes workflow by enabling isolated context and namespace management for each terminal.

## Install

### [RPMs](https://copr.fedorainfracloud.org/coprs/brandfbb/kubesw/): Fedora, CentOS, OpenSuse, Rocky, Rhel
```
dnf copr enable brandfbb/kubesw
dnf install kubesw
```

### Binary: MacOS (amd64/arm64), Windows, Linux
```
https://github.com/spideyz0r/kubesw/releases
```
### From source
```
git checkout https://github.com/spideyz0r/kubesw
cd kubesw; go build -v -o kubesw
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
  version     Get version of kubesw

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

## FZF - Fuzzy Finder
Running the set option without actually providing an argument will allow the context or namespace to be selected in a fzf fashion.

Example:
```
kubesw set ns
> Filter...
5/5 ─────────────────
> default
  kube-node-lease
  kube-public
  kube-system
  kyll
```

## Aliases / shortnames
You can also use the short form for each command:
```
command   => aliases
namespace => namespaces, ns
context   => contexts, ctx
list      => ls, l
get       => current, g
set       => switch, s
```
Some examples:
```kubesw get ns
kubesw ls ns
kubesw get ns
kubesw switch ns "somenamespace"
kubesw set ctx "somecontext"
kubesw s ctx "somecontext"
```

## Configuration file
The configuration file is optional. You can see the default values below.

It should be installed in ~/.kubesw/config.yaml
```yaml
---
# Use a custom PS1 when using kubesw
# Default: unset
PS1: '[\u@\h \W {$(kubesw get ctx) @ $(kubesw get ns)}]\\$ '
# Set default shell
# Default: auto
shell: /bin/bash
# Configure rc files to be read
# Note: they are read from your home directory
# The default values are listed below
default_rc:
  bash:
    - .bashrc
    - .bash_profile
    - .profile
    - .bash_login
    - .bash_logout
  zsh:
    - .zshrc
    - .zprofile
    - .zlogin
    - .zlogout
```

## Autocompletion
The autocompletion script can be generated with the following:
```
kubesw completion bash
```

You can add the following line to your ~/.bashrc:
```
source <(`which kubesw` completion bash)
```

For zshell, pick a path under $fpath and put the contents of `kubesw completion zsh`:
```
kubesw completion zsh >$(echo $fpath | cut -d " " -f1)/_kubesw
```

## Shell history
In order to have the shell's history, some extra configuration needs to be appended to the rc files.
This will make each command run to be persisted to the history file right-away.


### Bash
~/.bashrc
```
shopt -s histappend
PROMPT_COMMAND="history -a"
```

### Zsh
~/.zshrc
```
setopt inc_append_history
```

## TODO
- Improve error checks and messages
- Add --global flag for updating the namespace or context globally
- Investigate the use of eval instead of spawning shells
