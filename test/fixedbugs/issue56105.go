// compile -d=libfuzzer

package p

func f() {
	_ = [...][]int{{}, {}, {}, {}, {}}
}
