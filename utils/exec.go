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
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/fanux/sealos/pkg/logger"
)

func S(v interface{}, args ...interface{}) string {
	return fmt.Sprintf(fmt.Sprint(v), args...)
}

func execCode(err error) int {
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus()
		}
	}
	return 0
}

/*
Exec runs a command. exe is the path to the executable and args are arugment
passed to it.

If the command is executed successfuly without mistakes, (nil, 0) will be
returned. Otherwise, the error and error code will be returned.
NOTE the error code could be 0 with a non-nil error.

Stdout/stderr are directed to the current stdout/stderr.
*/
func Exec(exe interface{}, args ...string) (error, int) {
	cmd := exec.Command(S(exe), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err, execCode(err)
}

func ExecForPipe(exe interface{}, args ...string) error {
	cmdDbg := []string{S(exe)}
	cmdDbg = append(cmdDbg, args...)
	logger.Debug("执行命令是: %s", strings.Join(cmdDbg, " "))
	cmd := exec.Command(S(exe), args...)
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating StdoutPipe for cmd: #%v", err)
	}

	errReader, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating StderrPipe for cmd: #%v", err)
	}

	outScanner := bufio.NewScanner(outReader)
	go func() {
		for outScanner.Scan() {
			logger.Info(outScanner.Text())
		}
	}()

	errScanner := bufio.NewScanner(errReader)
	go func() {
		for errScanner.Scan() {
			logger.Info(errScanner.Text())
		}
	}()

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("error starting cmd: #%v", err)
	}

	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for cmd: #%v", err)
	}

	return nil
}
func ExecForPipeFlag(flag, exe interface{}, args ...string) error {
	cmdDbg := []string{S(exe)}
	cmdDbg = append(cmdDbg, args...)
	logger.Debug("执行命令是: %s", strings.Join(cmdDbg, " "))
	cmd := exec.Command(S(exe), args...)
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating StdoutPipe for cmd: #%v", err)
	}

	errReader, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating StderrPipe for cmd: #%v", err)
	}

	outScanner := bufio.NewScanner(outReader)
	go func() {
		for outScanner.Scan() {
			logger.Info(S("[%v]", flag) + outScanner.Text())
		}
	}()

	errScanner := bufio.NewScanner(errReader)
	go func() {
		for errScanner.Scan() {
			logger.Info(S("[%v]", flag) + errScanner.Text())
		}
	}()

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("error starting cmd: #%v", err)
	}

	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for cmd: #%v", err)
	}

	return nil
}

/*
ExecWithStdout is similar to Exec but the stdout is captured and returned as
the first return value.
*/
func ExecWithStdout(exe interface{}, args ...string) (stdout string, err error, errCode int) {
	var stdoutBuf bytes.Buffer

	cmd := exec.Command(S(exe), args...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	return string(stdoutBuf.Bytes()), err, execCode(err)
}

/*
ExecWithStdout is similar to Exec but the stdout/stderr are captured and
returned as the first/second return values.
*/
func ExecWithStdErrOut(exe interface{}, args ...string) (stdout, stderr string, err error, errCode int) {
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command(S(exe), args...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	return string(stdoutBuf.Bytes()), string(stderrBuf.Bytes()), err, execCode(err)
}

/*
Eval is similar to ExecWithStdout but with stdout captured and returned as a
string. Trainling newlines are deleted.
*/
func Eval(exe interface{}, args ...string) string {
	out, _, _ := ExecWithStdout(exe, args...)
	return strings.TrimRight(out, "\r\n")
}

/*
Bash runs a command with bash. Return values are defined in Exec.
*/
func Bash(cmd interface{}, args ...interface{}) (error, int) {
	return Exec("bash", "-c", S(cmd, args...))
}

/*
BashWithStdout is similar to Bash but with stdout captured and returned as a
string.
*/
func BashWithStdout(cmd interface{}, args ...interface{}) (string, error, int) {
	return ExecWithStdout("bash", "-c", S(cmd, args...))
}

/*
BashEval is similar to BashWithStdout but only returns captured stdout as a
string. Trainling newlines are deleted. It's like the backtick substitution in
Bash.
*/
func BashEval(cmd interface{}, args ...interface{}) string {
	out, _, _ := BashWithStdout(cmd, args...)
	return strings.TrimRight(out, "\r\n")
}
