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

package utils

import (
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/davecgh/go-spew/spew"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"

	metav1alpha1 "awsctrl.io/apis/meta/v1alpha1"
)

var (
	ControllerOwnerKey        = ".metadata.controller"
	StackTemplateVersionLabel = "stack-template-version"
)

func IsStatusComplete(status metav1alpha1.ConditionStatus) bool {
	return status == metav1alpha1.CreateCompleteStatus ||
		status == metav1alpha1.UpdateCompleteStatus ||
		status == metav1alpha1.UpdateRollbackCompleteStatus ||
		status == metav1alpha1.RollbackCompleteStatus
}

// ContainsFinalizer will check if the finalizer exists
func ContainsFinalizer(obj metav1.ObjectMeta, finalizer string) bool {
	return SliceContains(obj.GetFinalizers(), finalizer)
}

// SliceContains will check a list for something
func SliceContains(items []string, name string) bool {
	for _, f := range items {
		if f == name {
			return true
		}
	}
	return false
}

// AddFinalizer will add the finalizer from the list
func AddFinalizer(obj metav1.ObjectMeta, finalizer string) metav1.ObjectMeta {
	obj.SetFinalizers(append(obj.GetFinalizers(), finalizer))
	return obj
}

// RemoveFinalizer will remove the finalizer from the list
func RemoveFinalizer(obj metav1.ObjectMeta, finalizer string) metav1.ObjectMeta {
	var finalizers []string
	for _, f := range obj.GetFinalizers() {
		if f == finalizer {
			continue
		}
		finalizers = append(finalizers, f)
	}
	obj.SetFinalizers(finalizers)
	return obj
}

// ComputeHash will take in an object and generate a repetable hash for versioning
func ComputeHash(obj interface{}) string {
	objhasher := fnv.New32a()
	deepHashObject(objhasher, obj)

	return rand.SafeEncodeString(fmt.Sprint(objhasher.Sum32()))
}

func deepHashObject(hasher hash.Hash, objectToWrite interface{}) {
	hasher.Reset()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	printer.Fprintf(hasher, "%#v", objectToWrite)
}
