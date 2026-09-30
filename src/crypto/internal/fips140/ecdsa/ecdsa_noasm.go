//go:build !s390x || purego

package ecdsa

func sign[P Point[P]](c *Curve[P], priv *PrivateKey, drbg *hmacDRBG, hash []byte) (*Signature, error) {
	return signGeneric(c, priv, drbg, hash)
}

func verify[P Point[P]](c *Curve[P], pub *PublicKey, hash []byte, sig *Signature) error {
	return verifyGeneric(c, pub, hash, sig)
}
