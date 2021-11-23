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

package cmd

import (
	"fmt"
	"os"

	"github.com/cuisongliu/kube-100years/pkg"
	"github.com/cuisongliu/kube-100years/utils"
	"github.com/fanux/sealos/pkg/logger"
	"github.com/spf13/cobra"
)

var k8sVersion, platform, image, mirrorRepo string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kube-100years",
	Short: "Using docker build k8s binary. Hotfix 100years cert, output kubelet、kubeadm、kubectl.",
	Run: func(cmd *cobra.Command, args []string) {
		tips := fmt.Sprintf("Using docker build k8s binary and fix 100years cert.Current branch/tag is %s , platform is %s?(y/n)", k8sVersion, platform)
		if utils.Confirm(tips) {
			b := &pkg.Version{
				K8sVersion: k8sVersion,
				Platform:   platform,
				Pwd:        utils.Pwd(),
				Image:      image,
				MirrorRepo: mirrorRepo,
			}
			b.K8s100y()
		} else {
			logger.Warn("User cancel build 100years cert kubernetes.")
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.Flags().StringVarP(&k8sVersion, "version", "", "master", "kubernetes branch or tag")
	rootCmd.Flags().StringVarP(&platform, "platform", "p", "linux/amd64", "kubernetes platform, ex linux/amd64 linux/arm64")
	rootCmd.Flags().StringVarP(&image, "image", "i", "cuisongliu/kube-build:alpine-high", "build kubernetes binary (kubectl、kubelet、kubeadm)")
	rootCmd.Flags().StringVarP(&mirrorRepo, "mirror-repo", "", "https://gitee.com/mirrors/Kubernetes.git", "kubernetes mirror repos addr")

}
