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

// Package stdin takes input from cli via pipes. Built especially to run build tests on drone CI
package stdin

import (
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

// GetStdInput method takes user input from the cli
func GetStdInput() string {
	byteInput, _ := terminal.ReadPassword(int(syscall.Stdin))
	return string(byteInput)
}
