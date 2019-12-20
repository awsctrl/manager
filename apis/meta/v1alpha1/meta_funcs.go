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

package v1alpha1

import (
	"fmt"
	"strings"
	"unicode"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
)

// Empty will allow controllers to ignore these references
func (so *ObjectReference) Empty() bool {
	return so.ObjectRef.Name == "" &&
		so.ObjectRef.Key == "" &&
		so.Id == "" &&
		so.Arn == ""
}

// String will return stack object string
// TODO(christopherhein) This overall will need to be updated to make sure that
// the person has RBAC or something to allow this to happen
func (so *ObjectReference) String(client dynamic.Interface) (string, error) {
	if so.Empty() {
		return "", nil
	}

	if so.ObjectRef.Name != "" &&
		so.ObjectRef.Kind != "" &&
		so.ObjectRef.APIVersion != "" &&
		so.ObjectRef.Key != "" {

		stackGVR := schema.GroupVersionResource{
			Group:    "cloudformation.awsctrl.io",
			Version:  "v1alpha1",
			Resource: "stacks",
		}
		domainslice := strings.Split(so.ObjectRef.APIVersion, ".")

		key := types.NamespacedName{
			Name:      strings.Join([]string{domainslice[0], strings.ToLower(so.ObjectRef.Kind), so.ObjectRef.Name, so.ObjectRef.Namespace}, "-"),
			Namespace: so.ObjectRef.Namespace,
		}
		stackClient := client.Resource(stackGVR)

		stack, err := stackClient.Namespace(key.Namespace).Get(key.Name, metav1.GetOptions{})
		if err != nil {
			return "", fmt.Errorf("error %+v fetching %+v", err, key)
		}

		m, _, err := unstructured.NestedStringMap(stack.Object, "status", "outputs")
		if err != nil {
			return "", err
		}

		value, ok := m[so.ObjectRef.Key]
		if !ok {
			return "", fmt.Errorf("could not find stack object %v in %+v", lowerfirst(so.ObjectRef.Key), m)
		}
		return value, nil
	}

	if so.Id != "" {
		return so.Id, nil
	}

	if so.Arn != "" {
		return so.Arn, nil
	}
	return "", fmt.Errorf("could not find object reference %+v", so)
}

func lowerfirst(str string) string {
	a := []rune(str)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
