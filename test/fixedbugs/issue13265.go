// errorcheck -0 -race

//go:build (linux && amd64) || (linux && ppc64le) || (darwin && amd64) || (freebsd && amd64) || (netbsd && amd64) || (windows && amd64)


// Issue 13265: nil pointer deref.

package p

func f() {
    var c chan chan chan int
    for ; ; <-<-<-c {
    }
}
