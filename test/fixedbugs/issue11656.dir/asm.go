//go:build ppc64 || ppc64le

package main

func syncIcache(p uintptr)
