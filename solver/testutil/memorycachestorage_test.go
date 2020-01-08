package testutil

import (
	"testing"

	"github.com/x0rzkov/buildkit/solver"
)

func TestMemoryCacheStorage(t *testing.T) {
	RunCacheStorageTests(t, func() (solver.CacheKeyStorage, func()) {
		return solver.NewInMemoryCacheStorage(), func() {}
	})
}
