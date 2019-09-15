package controller

import (
	"awsctrl.io/pkg/controller/stack"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, stack.Add)
}
