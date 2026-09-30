// DO NOT EDIT (use 'go test -v -update-expected' instead.)
// See cmd/compile/internal/inline/inlheur/testdata/props/README.txt
// for more information on the format of this file.
// <endfilepreamble>
package params

// acrosscall.go T_feeds_indirect_call_via_call_toplevel 15 0 1
// ParamFlags
//   0 ParamFeedsIndirectCall
// <endpropsdump>
// {"Flags":0,"ParamFlags":[8],"ResultFlags":null}
// callsite: acrosscall.go:16:12|0 flagstr "" flagval 0 score 20 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_feeds_indirect_call_via_call_toplevel(f func(int)) {
	callsparam(f)
}

// acrosscall.go T_feeds_indirect_call_via_call_conditional 27 0 1
// ParamFlags
//   0 ParamMayFeedIndirectCall
// <endpropsdump>
// {"Flags":0,"ParamFlags":[16],"ResultFlags":null}
// callsite: acrosscall.go:29:13|0 flagstr "" flagval 0 score 20 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_feeds_indirect_call_via_call_conditional(f func(int)) {
	if G != 101 {
		callsparam(f)
	}
}

// acrosscall.go T_feeds_conditional_indirect_call_via_call_toplevel 41 0 1
// ParamFlags
//   0 ParamMayFeedIndirectCall
// <endpropsdump>
// {"Flags":0,"ParamFlags":[16],"ResultFlags":null}
// callsite: acrosscall.go:42:23|0 flagstr "" flagval 0 score 24 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_feeds_conditional_indirect_call_via_call_toplevel(f func(int)) {
	callsparamconditional(f)
}

// acrosscall.go T_feeds_if_via_call 53 0 1
// ParamFlags
//   0 ParamFeedsIfOrSwitch
// <endpropsdump>
// {"Flags":0,"ParamFlags":[32],"ResultFlags":null}
// callsite: acrosscall.go:54:9|0 flagstr "" flagval 0 score 8 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_feeds_if_via_call(x int) {
	feedsif(x)
}

// acrosscall.go T_feeds_if_via_call_conditional 65 0 1
// ParamFlags
//   0 ParamMayFeedIfOrSwitch
// <endpropsdump>
// {"Flags":0,"ParamFlags":[64],"ResultFlags":null}
// callsite: acrosscall.go:67:10|0 flagstr "" flagval 0 score 8 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_feeds_if_via_call_conditional(x int) {
	if G != 101 {
		feedsif(x)
	}
}

// acrosscall.go T_feeds_conditional_if_via_call 79 0 1
// ParamFlags
//   0 ParamMayFeedIfOrSwitch
// <endpropsdump>
// {"Flags":0,"ParamFlags":[64],"ResultFlags":null}
// callsite: acrosscall.go:80:20|0 flagstr "" flagval 0 score 12 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_feeds_conditional_if_via_call(x int) {
	feedsifconditional(x)
}

// acrosscall.go T_multifeeds1 93 0 1
// ParamFlags
//   0 ParamFeedsIndirectCall|ParamMayFeedIndirectCall
//   1 ParamNoInfo
// <endpropsdump>
// {"Flags":0,"ParamFlags":[24,0],"ResultFlags":null}
// callsite: acrosscall.go:94:12|0 flagstr "" flagval 0 score 20 mask 0 maskstr ""
// callsite: acrosscall.go:95:23|1 flagstr "" flagval 0 score 24 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_multifeeds1(f1, f2 func(int)) {
	callsparam(f1)
	callsparamconditional(f1)
}

// acrosscall.go T_acrosscall_returnsconstant 106 0 1
// ResultFlags
//   0 ResultAlwaysSameConstant
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[8]}
// callsite: acrosscall.go:107:24|0 flagstr "" flagval 0 score 2 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_acrosscall_returnsconstant() int {
	return returnsconstant()
}

// acrosscall.go T_acrosscall_returnsmem 118 0 1
// ResultFlags
//   0 ResultIsAllocatedMem
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[2]}
// callsite: acrosscall.go:119:19|0 flagstr "" flagval 0 score 2 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_acrosscall_returnsmem() *int {
	return returnsmem()
}

// acrosscall.go T_acrosscall_returnscci 130 0 1
// ResultFlags
//   0 ResultIsConcreteTypeConvertedToInterface
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[4]}
// callsite: acrosscall.go:131:19|0 flagstr "" flagval 0 score 7 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_acrosscall_returnscci() I {
	return returnscci()
}

// acrosscall.go T_acrosscall_multiret 140 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":[0],"ResultFlags":[0]}
// callsite: acrosscall.go:142:25|0 flagstr "" flagval 0 score 2 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_acrosscall_multiret(q int) int {
	if q != G {
		return returnsconstant()
	}
	return 0
}

// acrosscall.go T_acrosscall_multiret2 154 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":[0],"ResultFlags":[0]}
// callsite: acrosscall.go:156:25|0 flagstr "" flagval 0 score 2 mask 0 maskstr ""
// callsite: acrosscall.go:158:25|1 flagstr "" flagval 0 score 2 mask 0 maskstr ""
// <endcallsites>
// <endfuncpreamble>
func T_acrosscall_multiret2(q int) int {
	if q == G {
		return returnsconstant()
	} else {
		return returnsconstant()
	}
}

func callsparam(f func(int)) {
	f(2)
}

func callsparamconditional(f func(int)) {
	if G != 101 {
		f(2)
	}
}

func feedsif(x int) int {
	if x != 101 {
		return 42
	}
	return 43
}

func feedsifconditional(x int) int {
	if G != 101 {
		if x != 101 {
			return 42
		}
	}
	return 43
}

func returnsconstant() int {
	return 42
}

func returnsmem() *int {
	return new(int)
}

func returnscci() I {
	var q Q
	return q
}

type I interface {
	Foo()
}

type Q int

func (q Q) Foo() {
}

var G int
