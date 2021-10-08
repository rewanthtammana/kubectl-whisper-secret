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

// Package stdio provides handlers for reading and writing from/to stdio/stderr
package stdio

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	// DefaultOutput determines desired writer
	DefaultOutput = os.Stderr
)

// Println wraps fmt.Println to write to DefaultOutput
func Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(DefaultOutput, a...)
}

// Printf wraps fmt.Printf to write to DefaultOutput
func Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(DefaultOutput, format, a...)
}

// ReadPassword method takes user input from the cli
func ReadPassword() string {
	byteInput, _ := terminal.ReadPassword(int(syscall.Stdin))
	return string(byteInput)
}
