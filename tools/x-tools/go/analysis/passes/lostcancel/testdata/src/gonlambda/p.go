package gonlambda

import "context"

func legacy(parent context.Context) {
	callback := func() {
		_, cancel := context.WithCancel(parent) // want "cancel function is not used on all paths"
		if parent.Err() != nil {
			cancel()
		}
	} // want "this return statement may be reached without using the cancel var"
	callback()
}

func modern(parent context.Context) {
	var callback func() = () => {
		_, cancel := context.WithCancel(parent) // want "cancel function is not used on all paths"
		if parent.Err() != nil {
			cancel()
		}
	} // want "this return statement may be reached without using the cancel var"
	callback()
}

func safe(parent context.Context) {
	var callback func() = () => {
		_, cancel := context.WithCancel(parent)
		defer cancel()
	}
	callback()
}
