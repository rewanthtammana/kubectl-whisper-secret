<h1 align="center">kubectl-whisper-secret</h1>

<h4 align="center">kubectl-whisper-secret plugin allows users to create secrets with secure input prompt to prevent information leakages through terminal history, shoulder surfing attacks, etc</h4>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/rewanthtammana/kubectl-whisper-secret">
    <img src="https://goreportcard.com/badge/github.com/rewanthtammana/kubectl-whisper-secret">
  </a>
  <a href="https://github.com/rewanthtammana/kubectl-whisper-secret/blob/main/LICENSE">
    <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg">
  </a>
  <!-- <a href="https://github.com/rewanthtammana/kubectl-whisper-secret/releases">
    <img src="https://img.shields.io/github/downloads/rewanthtammana/kubectl-whisper-secret/total.svg?style=for-the-badge">
  </a> -->
</p>

## Installation

Installing with krew
```console
kubectl krew install whisper-secret
```

Installing with wget
```console
wget https://github.com/rewanthtammana/kubectl-whisper-secret/releases/download/v1.1.0/kubectl-whisper-secret-Linux-x86_64.tar.gz
sudo tar xvf kubectl-whisper-secret-Linux-x86_64.tar.gz -C /usr/bin/
```

## Why to use

### Problem statement

Analyze the examples below

```console
$ kubectl create secret generic my-secret --from-literal key1=value1 --from-literal key2=value2
```

```console
$ kubectl create secret docker-registry my-secret --docker-password s3cur3D0ck3rP@ssw0rD
```

`kubectl create secret` has a few sub-commands we use most often that can possibly leak sensitive information in multiple ways like terminal history, shoulder surfing, etc.

### Proposed solution with examples

`kubectl whisper-secret` plugin allows users to create secrets with secure input prompt for fields like `--from-literal` and `--docker-password` that contain sensitive information

```console
$ kubectl whisper-secret generic my-secret --from-literal key1 --from-literal key2
Enter value for key1: 
Enter value for key2: 
secret/my-secret created
```

```console
$ kubectl whisper-secret docker-registry my-secret --docker-password -- -n test --docker-username admin
Enter value for docker password: 
secret/my-secret created
```

```console
$ kubectl whisper-secret docker-registry my-secret --docker-password --verbose -- -n test --docker-username admin
Enter value for docker password: 
[+] Command: kubectl create docker-registry my-secret --docker-password abcdef -n test --docker-username admin
secret/my-secret created
```

```console
$ kubectl whisper-secret docker-registry my-secret --docker-password --print-only
Enter value for docker password: 
[*] Command: kubectl create docker-registry my-secret --docker-password abcdef
```

## Developer build

Build from Makefile
```console
make build
```

Build only for Linux
```console
make build-linux
```

Build with go
```console
go get ./...
go build -o kubectl-whisper-secret main.go
mv kubectl-whisper-secret /usr/bin
```

Cross platform builds with go
```console
go get ./...
GOOS=windows GOARCH=amd64 go build -o kubectl-whisper-secret.exe main.go
```


## Usage

```console
rewanth@ubuntu:~/go/src/kubectl-whisper-secret$ kubectl whisper-secret -h
"kubectl whisper-secret" creates kubectl secrets by taking sensitive input from console.
More info: https://github.com/rewanthtammana/kubectl-whisper-secret

Usage:
  kubectl-whisper-secret [flags]
  kubectl-whisper-secret [command]

Examples:

Create generic secret in default namespace:
$ kubectl whisper-secret generic my-secret --from-literal key1 --from-literal key2

Provide non-existing/unknown flags after double hypen (--)

Create generic secret in test namespace:
$ kubectl whisper-secret generic my-secret --from-literal key1 --from-literal key2 -- -n test

Create docker-registry secret in default namespace:
$ kubectl whisper-secret docker-registry my-secret --docker-password -- --docker-server=DOCKER_REGISTRY_SERVER --docker-username=DOCKER_USER --docker-email=DOCKER_EMAIL

Available Commands:
  docker-registry Take docker-registry password input from console
  generic         Create generic secrets by taking input from console
  help            Help about any command

Flags:
  -h, --help   help for kubectl-whisper-secret

Use "kubectl-whisper-secret [command] --help" for more information about a command.
```

