package frontend

import (
	"context"
	"io"

	"github.com/x0rzkov/buildkit/cache"
	"github.com/x0rzkov/buildkit/client"
	"github.com/x0rzkov/buildkit/executor"
	gw "github.com/x0rzkov/buildkit/frontend/gateway/client"
	digest "github.com/opencontainers/go-digest"
)

type Frontend interface {
	Solve(ctx context.Context, llb FrontendLLBBridge, opt map[string]string) (*Result, error)
}

type FrontendLLBBridge interface {
	Solve(ctx context.Context, req SolveRequest) (*Result, error)
	ResolveImageConfig(ctx context.Context, ref string, opt gw.ResolveImageConfigOpt) (digest.Digest, []byte, error)
	Exec(ctx context.Context, meta executor.Meta, rootfs cache.ImmutableRef, stdin io.ReadCloser, stdout, stderr io.WriteCloser) error
}

type SolveRequest = gw.SolveRequest

type CacheOptionsEntry = gw.CacheOptionsEntry

type WorkerInfos interface {
	WorkerInfos() []client.WorkerInfo
}
