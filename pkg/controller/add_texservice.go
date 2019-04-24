package controller

import (
	"github.com/boernd/tex-operator/pkg/controller/texservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, texservice.Add)
}
