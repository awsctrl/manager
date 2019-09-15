package controllerutils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
