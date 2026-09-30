package bar

import "issue44732.dir/foo"

type Bar struct {
	Foo *foo.Foo
}
