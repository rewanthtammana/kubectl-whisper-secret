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

// Package execCmd takes command as input that needs to be executed in console
package execCmd

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

// Run method takes the command as input that needs to be executed in console
func Run(finalCmd string) (error, bytes.Buffer, bytes.Buffer) {
	var execCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		execCmd = exec.Command("cmd", "/c", finalCmd)
	} else if runtime.GOOS == "linux" {
		execCmd = exec.Command("sh", "-c", finalCmd)
	} else {
		osError := fmt.Sprintf("Unrecognized OS: %s", runtime.GOOS)
		var osStderrMessage bytes.Buffer
		osStderrMessage.WriteString("Please create an issue here, https://github.com/rewanth1997/kubectl-whisper-secret/issues")
		return errors.New(osError), osStderrMessage, osStderrMessage
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	execCmd.Stdout = &out
	execCmd.Stderr = &stderr
	err := execCmd.Run()

	return err, stderr, out
}
