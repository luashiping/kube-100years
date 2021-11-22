package utils

import (
	"fmt"
	"github.com/fanux/sealos/pkg/logger"
)

func Download(pwd, remoteUrl, outputFile string) {
	logger.Info("当前执行阶段是: 下载上传的产物")
	cloneShell := `cd %s  && wget %s -O %s`
	cloneShellResult := fmt.Sprintf(cloneShell, pwd, remoteUrl, outputFile)
	if err := ExecForPipe("/bin/bash", "-c", cloneShellResult); err != nil {
		logger.Fatal("执行shell报错: %s", err.Error())
	}
}
