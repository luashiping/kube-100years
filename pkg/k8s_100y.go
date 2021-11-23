/*
Copyright © 2021 cuisongliu@qq.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"fmt"

	utils2 "github.com/cuisongliu/kube-100years/utils"
	"github.com/fanux/sealos/pkg/logger"
)

//init repos

type Version struct {
	K8sVersion string
	Platform   string //linux/amd64
	Pwd        string
	Image      string
	MirrorRepo string
}

func (v *Version) K8s100y() { //1.clone code
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
	hotfix100yearsFmt := `cat > hotfix_100years.sh <<EOF
#!/bin/bash
export KUBE_GIT_TREE_STATE="clean"
export KUBE_GIT_VERSION=$VERSION
export KUBE_BUILD_PLATFORMS=%s
make all WHAT=cmd/kubeadm GOFLAGS=-v
make all WHAT=cmd/kubelet GOFLAGS=-v
make all WHAT=cmd/kubectl GOFLAGS=-v
EOF

chmod a+x hotfix_100years.sh`
	hotfix100yearsShell := fmt.Sprintf(hotfix100yearsFmt, v.Platform)
	if err := utils2.ExecForPipe("/bin/bash", "-c", fmt.Sprintf("cd %s/%s && %s", v.Pwd, "kubernetes", hotfix100yearsShell)); err != nil {
		logger.Fatal("执行write shell报错: %s", err.Error())
	}
	//4.
	//cuisongliu/kube-build:alpine-high
	buildFmt := `docker run -i --rm -v $(pwd):/go/src/k8s.io/kubernetes -w /go/src/k8s.io/kubernetes \
%s \
bash -c /go/src/k8s.io/kubernetes/hotfix_100years.sh`
	buildShell := fmt.Sprintf(buildFmt, v.Image)
	if err := utils2.ExecForPipe("/bin/bash", "-c", fmt.Sprintf("cd %s/%s && %s", v.Pwd, "kubernetes", buildShell)); err != nil {
		logger.Fatal("执行build shell报错: %s", err.Error())
	}
	//
	finalFmt := `cp _output/local/bin/%s/kubeadm .
cp _output/local/bin/%s/kubelet .
cp _output/local/bin/%s/kubectl .`
	finalShell := fmt.Sprintf(finalFmt, v.Platform, v.Platform, v.Platform)
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
