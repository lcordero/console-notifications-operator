package controller

import (
	"github.com/lcordero/console-notifications-operator/pkg/controller/consolenotification"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, consolenotification.Add)
}
