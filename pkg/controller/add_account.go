package controller

import (
	"awsctrl.io/pkg/controller/account"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, account.Add)
}
