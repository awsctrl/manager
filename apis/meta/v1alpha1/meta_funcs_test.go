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

package v1alpha1_test

import (
	"testing"

	"go.awsctrl.io/manager/apis/meta/v1alpha1"

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/scheme"
)

func TestObjectReference_String(t *testing.T) {
	stack := &cloudformationv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foobar-blah-baz-default",
			Namespace: "default",
		},
		Status: cloudformationv1alpha1.StackStatus{
			StatusMeta: metav1alpha1.StatusMeta{
				Outputs: map[string]string{
					"key": "value",
				},
			},
		},
	}

	scheme.AddToScheme(scheme.Scheme)
	cloudformationv1alpha1.AddToScheme(scheme.Scheme)

	k8sclient := fake.NewSimpleDynamicClient(scheme.Scheme, []runtime.Object{stack}...)
	type fields struct {
		ObjectRef v1alpha1.ObjectRef
		Id        string
		Arn       string
	}
	type args struct {
		k8sclient dynamic.Interface
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"TestWhenIdSet", fields{
				ObjectRef: v1alpha1.ObjectRef{},
				Id:        "id",
				Arn:       "",
			}, args{k8sclient}, "id", false,
		},
		{
			"TestWhenArnSet", fields{
				ObjectRef: v1alpha1.ObjectRef{},
				Id:        "",
				Arn:       "arn",
			}, args{k8sclient}, "arn", false,
		},
		{
			"TestWhenObjectReferenceSet", fields{
				ObjectRef: v1alpha1.ObjectRef{
					Name:       "baz",
					APIVersion: "foobar.awsctrl.io/v1alpha1",
					Kind:       "Blah",
					Namespace:  stack.Namespace,
					Key:        "key",
				},
				Id:  "",
				Arn: "",
			}, args{k8sclient}, "value", false,
		},
		{
			"TestWhenObjectReferenceNotSet", fields{
				ObjectRef: v1alpha1.ObjectRef{
					Name:       "baz",
					APIVersion: "foobar.awsctrl.io/v1alpha1",
					Kind:       "Blah",
					Namespace:  stack.Namespace,
					Key:        "test",
				},
				Id:  "",
				Arn: "",
			}, args{k8sclient}, "", true,
		},
		{
			"TestWhenNothingSet", fields{
				ObjectRef: v1alpha1.ObjectRef{},
				Id:        "",
				Arn:       "",
			}, args{k8sclient}, "", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			so := &v1alpha1.ObjectReference{
				ObjectRef: tt.fields.ObjectRef,
				Id:        tt.fields.Id,
				Arn:       tt.fields.Arn,
			}
			got, err := so.String(tt.args.k8sclient)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectReference.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ObjectReference.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
