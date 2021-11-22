package pkg

import (
	"fmt"
	utils2 "github.com/cuisongliu/kube-100years/utils"
	"github.com/fanux/sealos/pkg/logger"
)

//init repos

type Version struct {
	K8sVersion string
	IsArm      bool
	Pwd        string
}

func (v *Version) K8s100y() {

	//1.clone code
	utils2.Clone(v.Pwd, "https://gitee.com/mirrors/Kubernetes.git", v.K8sVersion)
	if err := utils2.ExecForPipe("/bin/bash", "-c", fmt.Sprintf("cd %s && mv %s %s", v.Pwd, "Kubernetes", "kubernetes")); err != nil {
		logger.Fatal("执行mv shell报错: %s", err.Error())
	}
	//2.sed shell
	sedShell := `sed -i "s#CertificateValidity.*#CertificateValidity = time.Hour * 24 * 365 * 100#g"  cmd/kubeadm/app/constants/constants.go
sed -i "s#now.Add.*#now.Add(duration365d * 100).UTC(),#g"  staging/src/k8s.io/client-go/util/cert/cert.go
sed -i "s#maxAge :=.*#maxAge :=time.Hour * 24 * 365 * 100#g"  staging/src/k8s.io/client-go/util/cert/cert.go
`
	if err := utils2.ExecForPipe("/bin/bash", "-c", fmt.Sprintf("cd %s/%s && %s", v.Pwd, "kubernetes", sedShell)); err != nil {
		logger.Fatal("执行sed shell报错: %s", err.Error())
	}

	//3.
	hotfix100yearsShell := `cat > hotfix_100years.sh <<EOF
#!/bin/bash
export KUBE_GIT_TREE_STATE="clean"
export KUBE_GIT_VERSION=$VERSION
export KUBE_BUILD_PLATFORMS=linux/amd64
make all WHAT=cmd/kubeadm GOFLAGS=-v
make all WHAT=cmd/kubelet GOFLAGS=-v
make all WHAT=cmd/kubectl GOFLAGS=-v
EOF

chmod a+x hotfix_100years.sh`
	if err := utils2.ExecForPipe("/bin/bash", "-c", fmt.Sprintf("cd %s/%s && %s", v.Pwd, "kubernetes", hotfix100yearsShell)); err != nil {
		logger.Fatal("执行write shell报错: %s", err.Error())
	}
	//4.
	buildShell := `docker run -i --rm -v $(pwd):/go/src/k8s.io/kubernetes -w /go/src/k8s.io/kubernetes \
cuisongliu/kube-build:alpine-high \
bash -c /go/src/k8s.io/kubernetes/hotfix_100years.sh`
	if err := utils2.ExecForPipe("/bin/bash", "-c", fmt.Sprintf("cd %s/%s && %s", v.Pwd, "kubernetes", buildShell)); err != nil {
		logger.Fatal("执行build shell报错: %s", err.Error())
	}
	//
	finalShell := `cp _output/local/bin/linux/amd64/kubeadm .
cp _output/local/bin/linux/amd64/kubelet .
cp _output/local/bin/linux/amd64/kubectl .`
	if err := utils2.ExecForPipe("/bin/bash", "-c", fmt.Sprintf("cd %s/%s && %s", v.Pwd, "kubernetes", finalShell)); err != nil {
		logger.Fatal("执行copy shell报错: %s", err.Error())
	}
	cpShell := `cp kubernetes/kubeadm .
cp kubernetes/kubelet .
cp kubernetes/kubectl .
rm -rf kubernetes`
	if err := utils2.ExecForPipe("/bin/bash", "-c", fmt.Sprintf("cd %s && %s", v.Pwd, cpShell)); err != nil {
		logger.Fatal("执行copy shell报错: %s", err.Error())
	}
}
