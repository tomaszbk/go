package a

type Here struct{ stuff int }
type Info struct{ Dir string }

func New() Here { return Here{} }
func (h Here) Dir(p string) (Info, error)

type I interface{ M(x string) }

type T = struct {
	Here
	I
}

var X T

var A = (*T).Dir
var B = T.Dir
var C = X.Dir
var D = (*T).M
var E = T.M
var F = X.M
