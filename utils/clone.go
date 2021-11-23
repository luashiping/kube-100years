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

import (
	"fmt"
	"strings"

	"github.com/fanux/sealos/pkg/logger"
)

func Clone(pwd, gitUrl, version string) {
	logger.Info("当前执行阶段是: Branch克隆代码")
	cloneShell := `cd %s && git clone %s  Kubernetes && \
cd %s && git checkout %s`
	dir := PathToFileName(gitUrl)
	dir = strings.ReplaceAll(dir, ".git", "")
	cloneShellResult := fmt.Sprintf(cloneShell, pwd, gitUrl, dir, version)
	if err := ExecForPipe("/bin/bash", "-c", cloneShellResult); err != nil {
		logger.Fatal("执行shell报错: %s", err.Error())
	}
}
