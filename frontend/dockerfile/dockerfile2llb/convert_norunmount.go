// +build !dfrunmount

package dockerfile2llb

import (
	"github.com/x0rzkov/buildkit/client/llb"
	"github.com/x0rzkov/buildkit/frontend/dockerfile/instructions"
)

func detectRunMount(cmd *command, allDispatchStates *dispatchStates) bool {
	return false
}

func dispatchRunMounts(d *dispatchState, c *instructions.RunCommand, sources []*dispatchState, opt dispatchOpt) ([]llb.RunOption, error) {
	return nil, nil
}
