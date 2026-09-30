// run


// Issue 4748.
// This program used to complain because inlining created two exit labels.

package main

func jump() {
        goto exit
exit:
        return
}
func main() {
        jump()
        jump()
}
