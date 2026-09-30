// compile


// Issue 2576
package bug

type T struct { a int }

func f(t T) {
        switch _, _ = t.a, t.a; {}
}