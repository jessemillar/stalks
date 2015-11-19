package controllers

import "github.com/jessemillar/stalks/accessors"

// ControllerGroup holds all config information for the controllers
type ControllerGroup struct {
	Accessors *accessors.AccessorGroup
}
