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

package scaffold

import (
	"bytes"
	"fmt"
	"unicode"

	"html/template"
	"path/filepath"

	"awsctrl.io/generator/pkg/input"
	"awsctrl.io/generator/pkg/resource"
	"github.com/spf13/afero"
	"golang.org/x/tools/imports"

	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/flect"
)

// Scaffold contains the functions for generating files
type Scaffold struct {
	fs afero.Fs
	r  *resource.Resource
}

// New initializes the scaffolder
func New(fs afero.Fs, r *resource.Resource) *Scaffold {
	return &Scaffold{
		fs: fs,
		r:  r,
	}
}

// Execute will generate the file
func (s *Scaffold) Execute(files ...input.File) error {
	afs := afero.Afero{
		Fs: s.fs,
	}

	for _, file := range files {
		fileinput := file.GetInput()
		path := fileinput.Path

		contents, err := s.doTemplate(fileinput, file)
		if err != nil {
			return err
		}

		dir := filepath.Dir(path)
		if err := afs.MkdirAll(dir, 0700); err != nil {
			return err
		}

		if err := afs.WriteFile(path, contents, 0600); err != nil {
			return err
		}
	}

	return nil
}

func (s *Scaffold) doTemplate(i input.Input, e input.File) ([]byte, error) {
	temp, err := newTemplate(e).Parse(i.TemplateBody)
	if err != nil {
		return nil, err
	}

	out := &bytes.Buffer{}
	err = temp.Execute(out, e)
	if err != nil {
		return nil, err
	}
	b := out.Bytes()

	if filepath.Ext(i.Path) == ".go" {
		b, err = imports.Process(i.Path, b, nil)
		if err != nil {
			fmt.Printf("%s\n", out.Bytes())
			return nil, err
		}
	}

	return b, nil
}

func newTemplate(t input.File) *template.Template {
	return template.New(fmt.Sprintf("%T", t)).Funcs(sprig.FuncMap()).Funcs(funcMap())
}

func funcMap() template.FuncMap {
	funcs := map[string]interface{}{
		"lowerfirst": lowerfirst,
		"pluralize": flect.Pluralize,
	}

	return funcs
}

func lowerfirst(str string) string {
	a := []rune(str)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

