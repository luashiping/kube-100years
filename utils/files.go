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
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//FetchAllFiles is return dir all files abs path
func FetchAllFiles(pathName string, filters []string) ([]string, error) {
	dirs, err := ioutil.ReadDir(pathName)
	if err != nil {
		return nil, err
	}
	data := make([]string, 0)
	for _, fd := range dirs {
		if !NotIn(fd.Name(), filters) {
			continue
		}
		fName := path.Join(pathName, fd.Name())
		if fd.IsDir() {
			files, err := FetchAllFiles(fName, filters)
			if err != nil {
				return nil, err
			}
			if len(files) > 0 {
				data = append(data, files...)
			}
		} else {
			data = append(data, fName)
		}
	}
	return data, nil
}

//Pwd is similar to os.Getwd() without error returned.
func Pwd() string {
	pwd, _ := os.Getwd()
	return pwd
}

//ScriptDir returns the folder of the current script.
func ScriptDir() string {
	return path.Dir(os.Args[0])
}

//ScriptName returns the filename of the current script not including the path.
func ScriptName() string {
	return path.Base(os.Args[0])
}

//Exists checks whether the path exists
func Exists(p interface{}, args ...interface{}) bool {
	_, err := os.Stat(S(p, args...))
	return err == nil
}

//IsDir returns true only if the path exists and indicates a directory
func IsDir(p interface{}, args ...interface{}) bool {
	info, err := os.Stat(S(p, args...))
	if err != nil {
		// the path does not exist
		return false
	}
	return info.Mode().IsDir()
}

//IsFile returns true only if the path exists and indicates a file
func IsFile(p interface{}, args ...interface{}) bool {
	info, err := os.Stat(S(p, args...))
	if err != nil {
		// the path does not exist
		return false
	}
	return !info.Mode().IsDir()
}

func PathToFileName(pathName string) string {
	data := strings.Split(pathName, "/")
	return data[len(data)-1]
}

func MkdirByShell(pwd, dir string) error {
	if err, _ := Bash("cd %s && mkdir -p %s", pwd, dir); err != nil {
		return err
	}
	return nil
}
