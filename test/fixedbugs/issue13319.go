// errorcheck


package main

func f(int, int) {
    switch x {
    case 1:
        f(1, g()   // ERROR "expecting \)|possibly missing comma or \)"
    case 2:
        f()
    case 3:
        f(1, g()   // ERROR "expecting \)|possibly missing comma or \)"
    }
}
