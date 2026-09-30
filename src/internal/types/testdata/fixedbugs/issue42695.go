package issue42695

const _ = 6e5518446744 // ERROR "malformed constant"
const _ uint8 = 6e5518446744 // ERROR "malformed constant"

var _ = 6e5518446744 // ERROR "malformed constant"
var _ uint8 = 6e5518446744 // ERROR "malformed constant"

func f(x int) int {
        return x + 6e5518446744 // ERROR "malformed constant"
}

var _ = f(6e5518446744 /* ERROR "malformed constant" */ )
