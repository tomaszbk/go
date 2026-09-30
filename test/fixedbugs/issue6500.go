// errorcheck

// Issue 6500: missing error when fallthrough appears in a block.

package main

func main() {
	var x int
	switch x {
	case 0:
		{
			fallthrough // ERROR "fallthrough"
		}
	case 1:
		{
			switch x {
			case 2:
				fallthrough
			case 3:
			}
		}
		fallthrough
	default:
	}
}
