// +build !linux

package file

import (
	"github.com/x0rzkov/buildkit/solver/llbsolver/ops/fileoptypes"
	"github.com/x0rzkov/buildkit/solver/pb"
	"github.com/pkg/errors"
	copy "github.com/tonistiigi/fsutil/copy"
)

func readUser(chopt *pb.ChownOpt, mu, mg fileoptypes.Mount) (*copy.ChownOpt, error) {
	return nil, errors.New("only implemented in linux")
}
