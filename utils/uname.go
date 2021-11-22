package utils

import "github.com/fanux/sealos/pkg/logger"

func GetSystemOS() string {
	os := Eval("/bin/bash", "-c", "echo $(uname -s) | tr '[A-Z]' '[a-z]'")
	logger.Debug("当前操作系统的OS是: %s", os)
	return os
}

func GetSystemArch() string {
	arch := Eval("/bin/bash", "-c", "echo $(uname -m)")
	logger.Debug("当前操作系统的ARCH是: %s", arch)
	return arch
}
