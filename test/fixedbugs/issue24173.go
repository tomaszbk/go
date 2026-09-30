// compile

package p

type arrayAlias = [10]int
type mapAlias = map[int]int
type sliceAlias = []int
type structAlias = struct{}

func Exported() {
	_ = arrayAlias{}
	_ = mapAlias{}
	_ = sliceAlias{}
	_ = structAlias{}
}
