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
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"awsctrl.io/generator/pkg/api"
	"awsctrl.io/generator/pkg/input"
	"awsctrl.io/generator/pkg/resource"

	kbinput "sigs.k8s.io/kubebuilder/pkg/scaffold/input"
	kbresource "sigs.k8s.io/kubebuilder/pkg/scaffold/resource"
)

var boilerplatePath string
var projectPath string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run will process the CloudFormation Resource Spec and generate files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()

		options := input.Options{
			Options: kbinput.Options{
				BoilerplatePath: boilerplatePath,
				ProjectPath:     projectPath,
			},
		}

		// Load CloudFormation Resource Specification
		r := &resource.Resource{
			Resource: kbresource.Resource{
				Namespaced: true,
				Group:      "ecr",
				Version:    "v1alpha1",
				Kind:       "Repository",
				ShortNames: []string{"repo"},
			},
			ResourceType: &resource.BaseResource{
				Properties: map[string]resource.Property{
					"RepositoryName": &resource.BaseProperty{
						Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositoryname",
						Type:          "String",
						Required:      false,
						UpdateType:    resource.ImmutableType,
					},
					"LifecyclePolicy": &resource.BaseProperty{
						Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-lifecyclepolicy",
						Required:      false,
						Type:          "LifecyclePolicy",
						UpdateType:    resource.MutableType,
					},
					"RepositoryPolicyText": &resource.BaseProperty{
						Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositorypolicytext",
						Required:      false,
						Type:          "Json",
						UpdateType:    resource.MutableType,
					},
				},
			},
			PropertyTypes: map[string]resource.ResourceType{
				"LifecyclePolicy": &resource.BaseResource{
					Properties: map[string]resource.Property{
						"LifecyclePolicyText": &resource.BaseProperty{
							Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-lifecyclepolicytext",
							Type:          "String",
							Required:      false,
							UpdateType:    resource.MutableType,
						},
						"RegistryId": &resource.BaseProperty{
							Documentation: "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-registryid",
							Type:          "String",
							Required:      false,
							UpdateType:    resource.MutableType,
						},
					},
				},
			},
		}

		builder := api.New(fs, options)
		if err := builder.Build(r); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&boilerplatePath, "boilerplate-path", "b", "./hack/boilerplate.go.txt", "Path to the boilerplate header.")
	runCmd.Flags().StringVarP(&projectPath, "project-path", "p", "./PROJECT", "Path to the project file.")
}
