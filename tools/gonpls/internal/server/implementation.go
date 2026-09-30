package server

import (
	"context"

	"golang.org/x/tools/gopls/internal/file"
	"golang.org/x/tools/gopls/internal/golang"
	"golang.org/x/tools/gopls/internal/label"
	"golang.org/x/tools/gopls/internal/protocol"
	"golang.org/x/tools/gopls/internal/telemetry"
	"golang.org/x/tools/internal/event"
)

func (s *server) Implementation(ctx context.Context, params *protocol.ImplementationParams) (_ []protocol.Location, rerr error) {
	recordLatency := telemetry.StartLatencyTimer("implementation")
	defer func() {
		recordLatency(ctx, rerr)
	}()

	ctx, done := event.Start(ctx, "server.Implementation", label.URI.Of(params.TextDocument.URI))
	defer done()

	fh, snapshot, release, err := s.session.FileOf(ctx, params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	defer release()
	if snapshot.FileKind(fh) != file.Go {
		return nil, nil // empty result
	}
	return golang.Implementation(ctx, snapshot, fh, params.Range)
}
