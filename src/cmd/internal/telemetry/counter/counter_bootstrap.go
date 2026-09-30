//go:build cmd_go_bootstrap || compiler_bootstrap

package counter

import "flag"

type dummyCounter struct{}

func (dc dummyCounter) Inc() {}

func Open()                                                               {}
func Inc(name string)                                                     {}
func New(name string) dummyCounter                                        { return dummyCounter{} }
func NewStack(name string, depth int) dummyCounter                        { return dummyCounter{} }
func CountFlags(name string, flagSet flag.FlagSet)                        {}
func CountFlagValue(prefix string, flagSet flag.FlagSet, flagName string) {}
