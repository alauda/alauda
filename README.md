[![Go Report Card](https://goreportcard.com/badge/github.com/alauda/alauda)](https://goreportcard.com/report/github.com/alauda/alauda)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/alauda/alauda)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/alauda/alauda/blob/master/LICENSE)

# alauda
A command-line interface for the Alauda Container Platform.

## Usage
```
$ alauda
Alauda CLI

Usage:
  alauda [command]

Available Commands:
  help        Help about any command
  kubectl     kubectl controls the Kubernetes cluster manager
  version     Display version of Alauda CLI

Flags:
  -h, --help   help for alauda

Use "alauda [command] --help" for more information about a command.
```

## Building alauda

To build and install the alauda CLI:
```
make build
```

To uninstall:
```
make clean
```

## Testing alauda
1. Place a valid kubeconfig file named `config` in the `$HOME/.kube` directory.
2. Run:
```
make test
```
