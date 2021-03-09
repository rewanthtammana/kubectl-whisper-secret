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
	"os"
)

var (
	verboseFlag             bool
	printOnlyFlag           bool
	rootCmdDescriptionShort = "Create kubectl secrets by taking sensitive input from console"
	rootCmdDescriptionLong  = `"kubectl whisper-secret" allows users to create kubectl secrets by taking sensitive input from the console.
More info: https://github.com/rewanth1997/kubectl-whisper-secret`
	rootCmdExamples = `
Create generic secret in default namespace:
$ kubectl whisper-secret generic my-secret --from-literal key1 --from-literal key2

Provide non-supported whisper-secret flags/options after double hypen (--)

Create generic secret in test namespace:
$ kubectl whisper-secret generic my-secret --from-literal key1 --from-literal key2 -- -n test

More examples:
$ kubectl whisper-secret generic -h
$ kubectl whisper-secret docker-registry -h`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "kubectl-whisper-secret",
	Short:   rootCmdDescriptionShort,
	Long:    rootCmdDescriptionLong,
	Example: rootCmdExamples,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Initiates ignore-case and stdin flags
func init() {}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
