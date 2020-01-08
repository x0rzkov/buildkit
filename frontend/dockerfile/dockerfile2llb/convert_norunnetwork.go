// +build !dfrunnetwork

package dockerfile2llb

import (
	"github.com/x0rzkov/buildkit/client/llb"
	"github.com/x0rzkov/buildkit/frontend/dockerfile/instructions"
)

func dispatchRunNetwork(c *instructions.RunCommand) (llb.RunOption, error) {
	return nil, nil
}
