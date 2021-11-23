package utils

import (
	"fmt"

	"github.com/fanux/sealos/pkg/logger"
)

func MD5Sum(path string) string {
	shell := fmt.Sprintf("md5sum %s | awk '{print $1}'", path)
	logger.Debug("md5执行的shell: %s", shell)
	return Eval("/bin/bash", "-c", shell)
}
