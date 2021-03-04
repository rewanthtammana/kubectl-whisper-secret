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
	genericCmdDescriptionShort = "Create generic secrets by taking input from console"
	genericCmdDescriptionLong  = `"kubectl ccsecret generic" takes secret values for given keys as input from user console 
More info: https://github.com/rewanth1997/kubectl-ccsecret`
	genericCmdExamples = `
Create generic secret in default namespace:
$ kubectl ccsecret generic my-secret --from-literal key1 --from-literal key2
	
Provide required non-existing/unknown options after double hypen (--)

Create generic secret in test namespace:
$ kubectl ccsecret generic my-secret --from-literal key1 --from-literal key2 -- -n test`
	fromLiteralArr []string
)

// Take user input from cli
func joinArgsWithKey(keyword string, argsArr []string) string {
	output := ""
	for _, key := range argsArr {
		fmt.Println("Enter value for " + key + " : ")
		output += fmt.Sprintf(" %s=%s=%s ", keyword, key, stdin.GetStdInput())
	}
	return output
}

// genericCmd represents the base command when called without any subcommands
var genericCmd = &cobra.Command{
	Use: "generic",
	Short:   genericCmdDescriptionShort,
	Long:    genericCmdDescriptionLong,
	Example: genericCmdExamples,
	Run: func(cmd *cobra.Command, args []string) {

		tail := joinArgsWithKey("--from-literal", fromLiteralArr)

		if printOnlyFlag {
			// kubectl create-secret generic <secret-name> --from-literal secret1 --from-literal secret2 -n test
			fmt.Println("[*] Generated:", "kubectl", "create", "secret", "generic", args[0], tail, strings.Join(args[1:], " "))
			return
		}

		if verboseFlag {
			// kubectl create-secret generic <secret-name> --from-literal secret1 --from-literal secret2 -n test
			fmt.Println("[+] Executing", "kubectl", "create", "secret", "generic", args[0], tail, strings.Join(args[1:], " "))
		}

		output, err := exec.Command("kubectl", "create", "secret", "generic", args[0], tail, strings.Join(args[1:], " ")).Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(output)
	},
}

// Initiates from-literal and verbose flags
func init() {
	rootCmd.AddCommand(genericCmd)
	genericCmd.Flags().StringArrayVarP(&fromLiteralArr, "from-literal", "", []string{}, "From literal")
	genericCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "Prints the final kubectl execution command")
	genericCmd.Flags().BoolVarP(&printOnlyFlag, "print-only", "p", false, "Only prints the final execution command (dry run)")
}
