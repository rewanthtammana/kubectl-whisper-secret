# kubectl-ccsecret (create-console-secret)

<h4 align="center">kubectl-ccsecret extension prompts input box to user while creating secrets to prevent leakages through terminal history, etc</h4>

## Usage

```console
rewanth@ubuntu:~/go/src/kubectl-ccsecret$ kubectl ccsecret -h
"kubectl ccsecret" creates kubectl secrets by taking sensitive input from console.
More info: https://github.com/rewanth1997/kubectl-ccsecret

Usage:
  kubectl-ccsecret [flags]
  kubectl-ccsecret [command]

Examples:

Create generic secret in default namespace:
$ kubectl ccsecret generic my-secret --from-literal key1 --from-literal key2

Provide non-existing/unknown flags after double hypen (--)

Create generic secret in test namespace:
$ kubectl ccsecret generic my-secret --from-literal key1 --from-literal key2 -- -n test

Create docker-registry secret in default namespace:
$ kubectl ccsecret docker-registry my-secret --docker-password -- --docker-server=DOCKER_REGISTRY_SERVER --docker-username=DOCKER_USER --docker-email=DOCKER_EMAIL

Available Commands:
  docker-registry Take docker-registry password input from console
  generic         Create generic secrets by taking input from console
  help            Help about any command

Flags:
  -h, --help   help for kubectl-ccsecret

Use "kubectl-ccsecret [command] --help" for more information about a command.
```

## Examples

TBD

## Why to use

### Problem statement

Analyze the examples below

```console
$ kubectl create secret generic my-secret --from-literal key1=value1 --from-literal key2=value2
```

```console
$ kubectl create secret docker-registry my-secret --docker-password s3cur3D0ck3rP@ssw0rD
```

`kubectl create secret` has a few sub-commands we use most often that leak sensitive information in multiple ways like terminal history, shoulder surfing, etc.

### Current solution

This plugin prompts users to provide inputs for fields that might possibly contain sensitive information like `--from-literal` and `--docker-password`

```console
$ kubectl ccsecret generic my-secret --from-literal key1 --from-literal key2
Enter value for key1: ******
Enter value for key2: ******
```

```console
$ kubectl ccsecret docker-password my-secret --docker-password
Enter value for docker password: ******
```

