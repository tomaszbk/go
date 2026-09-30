package lru_test

import (
	"testing"

	"golang.org/x/tools/gopls/internal/util/lru"
)

func TestSetUntypedNil(t *testing.T) {
	cache := lru.New[any, any](100 * 1e6)
	cache.Set(nil, nil, 1)
	if got, ok := cache.Get(nil); !ok || got != nil {
		t.Errorf("cache.Get(nil) = %v, %v, want nil, true", got, ok)
	}
}
