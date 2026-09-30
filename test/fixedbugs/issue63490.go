// compile

package main

type ResourceFunc struct {
	junk [8]int
	base assignmentBaseResource
}

type SubscriptionAssignmentResource struct {
	base assignmentBaseResource
}

type assignmentBaseResource struct{}

//go:noinline
func (a assignmentBaseResource) f(s string) ResourceFunc {
	println(s)
	return ResourceFunc{}
}

//go:noinline
func (r SubscriptionAssignmentResource) Hi() ResourceFunc {
	rf := r.base.f("Hello world")
	rf.base = r.base
	return rf
}

func main() {
	var r SubscriptionAssignmentResource
	r.Hi()
}
