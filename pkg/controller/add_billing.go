package controller

import (
	"github.com/prnvkv/billing-operator-example/pkg/controller/billing"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, billing.Add)
}
