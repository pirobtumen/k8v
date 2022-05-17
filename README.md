# k8v

Sometimes you need to work with different k8s cluster versions, or maybe you just wants to update
`kubectl`. This tool `k8v` or **_kubernets version manager_** it's what you are looking for.

## Usage

I have this script for personal use, so it's a bit manual:

```
mkdir ~/.k8v
export PATH=$PATH:~/.k8v
git clone https://github.com/pirobtumen/k8v.git
cd k8v
go run cmd/main.go <kubectl version>
kubectl version
```

> The command `export PATH=$PATH:~/.k8v` can be added to your `.zshrc` or `.bashrc` file.
