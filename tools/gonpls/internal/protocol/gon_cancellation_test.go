package protocol

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"golang.org/x/tools/internal/jsonrpc2"
	jsonrpc2_v2 "golang.org/x/tools/internal/jsonrpc2_v2"
)

type cancellationServer struct {
	Server
	cancel context.CancelFunc
	err    error
}

func (s cancellationServer) CodeAction(context.Context, *CodeActionParams) ([]CodeAction, error) {
	if s.cancel != nil {
		s.cancel()
	}
	return []CodeAction{{Title: "available fix"}}, s.err
}

// Exercise the actual generated codeAction dispatch, not just the classifier.
// Both transports must preserve real failures and valid code actions.
func TestGonCodeActionCancellation(t *testing.T) {
	realError := errors.New("analysis failed")
	for _, test := range []struct {
		name   string
		cancel bool
		err    error
		code   int64
	}{
		{"success", false, nil, 0},
		{"request cancelled", true, context.Canceled, int64(RequestCancelled)},
		{"cancelled successful result", true, nil, int64(RequestCancelled)},
		{"snapshot cancelled", false, context.Canceled, int64(ServerCancelled)},
		{"wrapped snapshot", false, fmt.Errorf("loading snapshot: %w", context.Canceled), int64(ServerCancelled)},
		{"real error", false, realError, 0},
		{"real error after cancellation", true, realError, 0},
	} {
		t.Run(test.name, func(t *testing.T) {
			for _, v2 := range []bool{false, true} {
				ctx, cancel := context.WithCancel(context.Background())
				server := cancellationServer{err: test.err}
				if test.cancel {
					server.cancel = cancel
				}
				var result any
				var got error
				if v2 {
					req, err := jsonrpc2_v2.NewCall(jsonrpc2_v2.Int64ID(1), "textDocument/codeAction", CodeActionParams{})
					if err != nil {
						t.Fatal(err)
					}
					result, got = ServerHandlerV2(server).Handle(ctx, req)
				} else {
					req, err := jsonrpc2.NewCall(jsonrpc2.NewIntID(1), "textDocument/codeAction", CodeActionParams{})
					if err != nil {
						t.Fatal(err)
					}
					handler := ServerHandler(server, func(context.Context, jsonrpc2.Replier, jsonrpc2.Request) error {
						t.Fatal("unhandled codeAction")
						return nil
					})
					err = handler(ctx, func(replyCtx context.Context, res any, err error) error {
						if replyCtx.Err() != nil {
							t.Error("reply must survive request cancellation")
						}
						result, got = res, err
						return nil
					}, req)
					if err != nil {
						t.Fatal(err)
					}
				}
				cancel()
				if test.code != 0 {
					var code int64
					if v2 {
						var wire *jsonrpc2_v2.WireError
						if !errors.As(got, &wire) {
							t.Fatalf("v2=%v: %v", v2, got)
						}
						code = wire.Code
					} else {
						var wire *jsonrpc2.WireError
						if !errors.As(got, &wire) {
							t.Fatalf("v2=%v: %v", v2, got)
						}
						code = wire.Code
					}
					if code != test.code || result != nil {
						t.Fatalf("v2=%v: result=%v error=%v code=%d want=%d", v2, result, got, code, test.code)
					}
				} else if got != test.err {
					t.Fatalf("v2=%v: genuine error changed: %v", v2, got)
				} else if got == nil && len(result.([]CodeAction)) != 1 {
					t.Fatalf("lost code actions: %v", result)
				}
			}
		})
	}
}
