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

package utils_test

import (
	"reflect"
	"testing"

	cloudformationv1alpha1 "awsctrl.io/apis/cloudformation/v1alpha1"
	metav1alpha1 "awsctrl.io/apis/meta/v1alpha1"
	"awsctrl.io/controllers/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsStatusComplete(t *testing.T) {
	type args struct {
		status metav1alpha1.ConditionStatus
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestCheckingCreateStatus", args{metav1alpha1.CreateCompleteStatus}, true},
		{"TestCheckingUpdateStatus", args{metav1alpha1.UpdateCompleteStatus}, true},
		{"TestCheckingUpdatingStatus", args{metav1alpha1.UpdateInProgressStatus}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsStatusComplete(tt.args.status); got != tt.want {
				t.Errorf("IsStatusComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsFinalizer(t *testing.T) {
	stack := &cloudformationv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "stack",
			Namespace:  "default",
			Finalizers: []string{"delete"},
		},
	}
	type args struct {
		obj       metav1.ObjectMeta
		finalizer string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestEmptyFinalizers", args{stack.ObjectMeta, "test"}, false},
		{"TestEmptyFinalizers", args{stack.ObjectMeta, "delete"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ContainsFinalizer(tt.args.obj, tt.args.finalizer); got != tt.want {
				t.Errorf("ContainsFinalizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceContains(t *testing.T) {
	type args struct {
		items []string
		name  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestSliceContainsHelper", args{[]string{"delete"}, "delete"}, true},
		{"TestSliceNotContainsHelper", args{[]string{"delete"}, "test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.SliceContains(tt.args.items, tt.args.name); got != tt.want {
				t.Errorf("SliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddFinalizer(t *testing.T) {
	stack := &cloudformationv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "stack",
			Namespace:  "default",
			Finalizers: []string{"delete"},
		},
	}
	type args struct {
		obj       metav1.ObjectMeta
		finalizer string
	}
	tests := []struct {
		name string
		args args
		want metav1.ObjectMeta
	}{
		{"TestNewFinalizer", args{stack.ObjectMeta, "test"}, func() metav1.ObjectMeta {
			newMeta := stack.ObjectMeta.DeepCopy()
			newMeta.Finalizers = append(newMeta.Finalizers, "test")
			return *newMeta
		}()},
		{"TestSameFinalizers", args{stack.ObjectMeta, "delete"}, func() metav1.ObjectMeta {
			newMeta := stack.ObjectMeta.DeepCopy()
			newMeta.Finalizers = append(newMeta.Finalizers, "delete")
			return *newMeta
		}()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.AddFinalizer(tt.args.obj, tt.args.finalizer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFinalizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFinalizer(t *testing.T) {
	stack := &cloudformationv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "stack",
			Namespace:  "default",
			Finalizers: []string{"delete"},
		},
	}

	type args struct {
		obj       metav1.ObjectMeta
		finalizer string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"TestRemoveNonExistingFinalizer", args{stack.ObjectMeta, "test"}, 1},
		{"TestRemovingFinalizers", args{stack.ObjectMeta, "delete"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.RemoveFinalizer(tt.args.obj, tt.args.finalizer); len(got.GetFinalizers()) != tt.wantCount {
				t.Errorf("RemoveFinalizer() = %v, want %v", len(got.GetFinalizers()), tt.wantCount)
			}
		})
	}
}

func TestComputeHash(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestGeneratingHash", args{struct{ foo string }{foo: "bar"}}, "67848d774f"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ComputeHash(tt.args.obj); got != tt.want {
				t.Errorf("ComputeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
