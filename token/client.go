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

// Package token generates random strings
package token

import (
	uuid "github.com/satori/go.uuid"
)

// Token defines what it means to generate random strings
type Token interface {
	// Generate creates a random string
	Generate() string
}

type token struct{}

// New initialized a new token
func New() token {
	return token{}
}

// Generate will generate a UUID V4 string
func (t token) Generate() string {
	return uuid.NewV4().String()
}
