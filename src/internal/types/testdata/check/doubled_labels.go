package p

func _() {
outer:
inner:
	for {
		continue inner
		break inner
	}
	goto outer
}

func _() {
outer:
inner:
	for {
		continue inner
		continue outer /* ERROR "invalid continue label outer" */
		break outer    /* ERROR "invalid break label outer" */
	}
	goto outer
}
