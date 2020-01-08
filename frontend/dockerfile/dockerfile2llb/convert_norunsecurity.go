// +build !dfrunsecurity

package dockerfile2llb

import (
	"github.com/x0rzkov/buildkit/client/llb"
	"github.com/x0rzkov/buildkit/frontend/dockerfile/instructions"
)

func dispatchRunSecurity(c *instructions.RunCommand) (llb.RunOption, error) {
	return nil, nil
}
