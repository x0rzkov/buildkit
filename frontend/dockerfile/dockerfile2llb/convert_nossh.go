// +build dfrunmount,!dfssh

package dockerfile2llb

import (
	"github.com/x0rzkov/buildkit/client/llb"
	"github.com/x0rzkov/buildkit/frontend/dockerfile/instructions"
	"github.com/pkg/errors"
)

func dispatchSSH(m *instructions.Mount) (llb.RunOption, error) {
	return nil, errors.Errorf("ssh mounts not allowed")
}
