package server

import (
	"context"

	"golang.org/x/tools/gopls/internal/golang"
	"golang.org/x/tools/gopls/internal/protocol"
)

func (s *server) ResolveCommand(ctx context.Context, param *protocol.ExecuteCommandParams) (*protocol.ExecuteCommandParams, error) {
	return golang.ResolveCommand(ctx, param, s.Options().ClientOptions)
}
