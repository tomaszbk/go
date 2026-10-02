package noder

import "cmd/compile/internal/syntax"

// Verify typed lambdas before the passes that recognize function boundaries.
// The checker supplies ordinary function literals with the original parameter
// objects; no printed types or second type check is involved.
func prepareLambdas(files []*syntax.File) {
	for _, file := range files {
		syntax.Inspect(file, func(n syntax.Node) bool {
			if lambda, ok := n.(*syntax.LambdaExpr); ok {
				assert(lambda.Lowered != nil)
			}
			return true
		})
	}
}
