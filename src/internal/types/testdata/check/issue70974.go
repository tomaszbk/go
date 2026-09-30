package p

func _() {
outer:
	for {
		break outer
	}

	for {
		break outer /* ERROR "invalid break label outer" */
	}
}

func _() {
outer:
	for {
		continue outer
	}

	for {
		continue outer /* ERROR "invalid continue label outer" */
	}
}
