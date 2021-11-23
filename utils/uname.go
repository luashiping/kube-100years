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
