// +build dfrunsecurity

package dockerfile2llb

import (
	"github.com/pkg/errors"

	"github.com/x0rzkov/buildkit/client/llb"
	"github.com/x0rzkov/buildkit/frontend/dockerfile/instructions"
	"github.com/x0rzkov/buildkit/solver/pb"
)

func dispatchRunSecurity(c *instructions.RunCommand) (llb.RunOption, error) {
	security := instructions.GetSecurity(c)

	switch security {
	case instructions.SecurityInsecure:
		return llb.Security(pb.SecurityMode_INSECURE), nil
	case instructions.SecuritySandbox:
		return llb.Security(pb.SecurityMode_SANDBOX), nil
	default:
		return nil, errors.Errorf("unsupported security mode %q", security)
	}
}
