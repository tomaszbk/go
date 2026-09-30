package p

var _ = (*interface /* ERROR "interface contains type constraints" */ {int})(nil)

// abbreviated test case from issue

type TypeSet interface{ int | string }

func _() {
	f((*TypeSet /* ERROR "interface contains type constraints" */)(nil))
}

func f(any) {}