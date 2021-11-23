package utils

import (
	"fmt"
	"strings"

	"github.com/fanux/sealos/pkg/logger"
)

func Clone(pwd, gitUrl, version string) {
	logger.Info("当前执行阶段是: Branch克隆外部代码")
	cloneShell := `cd %s && git clone %s  Kubernetes && \
cd %s && git checkout %s`
	dir := PathToFileName(gitUrl)
	dir = strings.ReplaceAll(dir, ".git", "")
	cloneShellResult := fmt.Sprintf(cloneShell, pwd, gitUrl, dir, version)
	if err := ExecForPipe("/bin/bash", "-c", cloneShellResult); err != nil {
		logger.Fatal("执行shell报错: %s", err.Error())
	}
}
