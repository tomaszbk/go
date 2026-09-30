package chacha8rand

var Block = block
var Block_generic = block_generic

func Seed(s *State) [4]uint64 {
	return s.seed
}
