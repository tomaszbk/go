//simdgen:category Shuffles

package spec

// Permute permutes x.
//
//	z[i] = x[indices[i] % len(x)]
//
//specgen:require indices=Uint{xN}x{xL}
func Permute[E Elt, W Width, E2 Uints](x Vec[E, W], indices Vec[E2, W]) (z Vec[E, W]) {
	z = makeVec[E, W]()
	for i := range z {
		z[i] = x[uint(indices[i])%uint(x.len())]
	}
	return z
}
