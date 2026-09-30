// compile

package p

func f() {
	var B bool
	B2 := (B || B && !B) && !B
	B3 := B2 || B
	for (B3 || B2) && !B2 && B {
	}
}
