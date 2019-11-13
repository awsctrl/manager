/*
Copyright © 2019 AWS Controller authors

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

package input

import (
	kbinput "sigs.k8s.io/kubebuilder/pkg/scaffold/input"
)

// File inputs the file interface for processing files
type File interface {
	// returns the file input for the file generator
	GetInput() Input
}

// ProjectFile loads project file from kubebuilder
type ProjectFile struct {
	kbinput.ProjectFile
}

// Input is the input for scaffolding a file
type Input struct {
	kbinput.Input
}

// Options is the main place for passed in options
type Options struct {
	kbinput.Options
}
