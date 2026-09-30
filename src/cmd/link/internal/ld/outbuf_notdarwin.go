//go:build !darwin

package ld

func (out *OutBuf) purgeSignatureCache() {}
