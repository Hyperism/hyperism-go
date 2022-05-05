/*
Copyright 2019 IBM Corp.

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

package metadata

import (
	"fmt"
	"runtime"
)

// Version config
const Version = "0.1.0"

// ProgramName config
const ProgramName = "swarmkeygen"

// PskHeader config
const PskHeader = "/key/swarm/psk/1.0.0/"

// EncodingType config
const EncodingType = "/base64/"

// GetVersionInfo function
func GetVersionInfo() string {
	return fmt.Sprintf("%s:\n Version: %s\n Go version: %s\n OS/Arch: %s",
		ProgramName, Version, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
}
