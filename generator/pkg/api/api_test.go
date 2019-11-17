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

package api_test

import (
	"testing"

	"awsctrl.io/generator/pkg/api"
	"awsctrl.io/generator/pkg/input"
	"awsctrl.io/generator/pkg/resource"
	"github.com/spf13/afero"

	kbinput "sigs.k8s.io/kubebuilder/pkg/scaffold/input"
	kbresource "sigs.k8s.io/kubebuilder/pkg/scaffold/resource"
)

func TestAPI_Build(t *testing.T) {
	fs := afero.NewMemMapFs()
	afs := afero.Afero{Fs: fs}

	afs.WriteFile("./hack/boilerplate.go.txt", []byte("// LICENSE"), 0644)

	a := api.New(fs, input.Options{Options: kbinput.Options{BoilerplatePath: "./hack/boilerplate.go.txt"}})

	r := &resource.Resource{
		Resource: kbresource.Resource{
			Namespaced: true,
			Group:      "ecr",
			Version:    "v1alpha1",
			Kind:       "Repository",
			ShortNames: []string{"repo"},
		},
		ResourceType: &resource.BaseResource{
			Attributes: map[string]map[string]string{
				"Arn": map[string]string{
					"PrimitiveType": "string",
				},
			},
			Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html",
			Properties: map[string]resource.Property{
				"RepositoryName": &resource.BaseProperty{
					Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositoryname",
					Type:          "string",
					Required:      false,
					UpdateType:    resource.ImmutableType,
				},
				"LifecyclePolicy": &resource.BaseProperty{
					Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-lifecyclepolicy",
					Required:      false,
					Type:          "LifecyclePolicy",
					UpdateType:    resource.MutableType,
				},
			},
		},
		PropertyTypes: map[string]resource.ResourceType{
			"LifecyclePolicy": &resource.BaseResource{
				Properties: map[string]resource.Property{
					"LifecyclePolicyText": &resource.BaseProperty{
						Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-lifecyclepolicytext",
						Type:          "string",
						Required:      false,
						UpdateType:    resource.MutableType,
					},
					"RegistryId": &resource.BaseProperty{
						Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-registryid",
						Type:          "string",
						Required:      false,
						UpdateType:    resource.MutableType,
					},
				},
			},
		},
	}

	type fields struct {
		api      *api.API
		resource *resource.Resource
	}
	tests := []struct {
		name     string
		fields   fields
		wantErr  bool
		wantFile string
	}{
		{"TestCreatingTypesFile", fields{a, r}, false, "apis/ecr/v1alpha1/repository_types.go"},
		{"TestCreatingtypesFile", fields{a, r}, false, "apis/ecr/v1alpha1/repository_types.go"},
		{"TestCreatingStackObjectFile", fields{a, r}, false, "apis/ecr/v1alpha1/zz_generated.repository.stackobject.go"},
		{"TestCreatingControllerFile", fields{a, r}, false, "controllers/ecr/repository_controller.go"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.api.Build(tt.fields.resource); (err != nil) != tt.wantErr {
				t.Errorf("API.Build() error = %v, wantErr %v", err, tt.wantErr)
			}

			if _, err := afs.Stat(tt.wantFile); err != nil {
				t.Errorf("API.Build() didn't create file %v", tt.wantFile)
			}
		})
	}
}

// TODO: Tests that test the contents of the files...
