/*
Copyright © 2021 Rewanth Cool

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

// Package cmd creates cli interface for this application
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"

	"github.com/rewanth1997/kubectl-ccsecret/pkg/stdin"
)

var (
	dockerPasswordFlag                bool
	dockerRegistryCmdDescriptionShort = "Take docker-registry password input from console"
	dockerRegistryCmdDescriptionLong  = `"kubectl ccsecret docker-registry" takes docker-password value from console 
More info: https://github.com/rewanth1997/kubectl-ccsecret`
	dockerRegistryCmdExamples = `
Provide required non-existing/unknown options after double hypen (--)

Create docker-registry secret in default namespace:
$ kubectl ccsecret docker-registry my-secret --docker-password -- --docker-server=DOCKER_REGISTRY_SERVER --docker-username=DOCKER_USER --docker-email=DOCKER_EMAIL

Create docker-registry secret in test namespace:
$ kubectl ccsecret docker-registry my-secret --docker-password -- -n test --docker-server=DOCKER_REGISTRY_SERVER --docker-username=DOCKER_USER --docker-email=DOCKER_EMAIL`
	tail           = ""
	dockerPassword string
)

// dockerRegistryCmd represents the base command when called without any subcommands
var dockerRegistryCmd = &cobra.Command{
	Use: "docker-registry",
	Short:   dockerRegistryCmdDescriptionShort,
	Long:    dockerRegistryCmdDescriptionLong,
	Example: dockerRegistryCmdExamples,
	Run: func(cmd *cobra.Command, args []string) {

		if dockerPasswordFlag {
			fmt.Print("Enter value for docker password: ")
			dockerPassword = stdin.GetStdInput()
			tail = fmt.Sprintf(" --docker-password %s ", dockerPassword)
		}

		if printOnlyFlag {
			fmt.Println("[*] Generated:", "kubectl", "create", "secret", "docker-registry", args[0], tail, strings.Join(args[1:], " "))
			return
		}

		if verboseFlag {
			fmt.Println("[+] Executing", "kubectl", "create", "secret", "docker-registry", args[0], tail, strings.Join(args[1:], " "))
		}

		output, err := exec.Command("kubectl", "create", "secret", "docker-registry", args[0], tail, strings.Join(args[1:], " ")).Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(output)
	},
}

// Initiates docker password flag and verbose flags
func init() {
	rootCmd.AddCommand(dockerRegistryCmd)
	dockerRegistryCmd.Flags().BoolVarP(&dockerPasswordFlag, "docker-password", "", false, "Docker password")
	dockerRegistryCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "Prints the final kubectl execution command")
	dockerRegistryCmd.Flags().BoolVarP(&printOnlyFlag, "print-only", "p", false, "Only prints the final execution command (dry run)")
}
