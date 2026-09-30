// Package nonewvars defines an Analyzer that applies suggested fixes
// to errors of the type "no new variables on left side of :=".
//
// # Analyzer nonewvars
//
// nonewvars: suggested fixes for "no new vars on left side of :="
//
// This checker provides suggested fixes for type errors of the
// type "no new vars on left side of :=". For example:
//
//	z := 1
//	z := 2
//
// will turn into
//
//	z := 1
//	z = 2
package nonewvars
