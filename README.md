### build kubernetes binary and fix 100 years cert

```
Using docker build k8s binary. Hotfix 100years cert, output kubelet、kubeadm、kubectl.

Usage:
  kube-100years [flags]

Flags:
  -h, --help                 help for kube-100years
  -i, --image string         build kubernetes binary (kubectl、kubelet、kubeadm) (default "cuisongliu/kube-build:alpine-high")
      --mirror-repo string   kubernetes mirror repos addr (default "https://gitee.com/mirrors/Kubernetes.git")
  -p, --platform string      kubernetes platform, ex linux/amd64 linux/arm64 (default "linux/amd64")
      --version string       kubernetes branch or tag (default "master")

```


Dependents:

- docker
- git
