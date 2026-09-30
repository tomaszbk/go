// errorcheck


package p

func _(s []int) {
        var i, j, k, l int
        _, _, _, _ = i, j, k, l

        for range s {}
        for i = range s {}
        for i, j = range s {}
        for i, j, k = range s {} // ERROR "range clause permits at most two iteration variables"
        for i, j, k, l = range s {} // ERROR "range clause permits at most two iteration variables"
}

func _(s chan int) {
        var i, j, k, l int
        _, _, _, _ = i, j, k, l

        for range s {}
        for i = range s {}
        for i, j = range s {} // ERROR "range over .* permits only one iteration variable"
        for i, j, k = range s {} // ERROR "range over .* permits only one iteration variable"
        for i, j, k, l = range s {} // ERROR "range over .* permits only one iteration variable"
}
