// Copyright 2026 The Gon Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gonmodernize suggests semantics-preserving uses of Gon syntax.
// Its analyzers are shared by gon fix and gonpls, including gon check.
// They deliberately recognize a conservative subset of possible rewrites:
// an omitted suggestion does not mean the source cannot be modernized by hand.
// Ordinary Go code remains valid; these diagnostics are optional hints.
package gonmodernize

import "golang.org/x/tools/go/analysis"

// Suite contains Gon's syntax modernization analyzers.
var Suite = []*analysis.Analyzer{
	ErrorAnalyzer,
	ConditionalAnalyzer,
	NilAnalyzer,
	LambdaAnalyzer,
}
