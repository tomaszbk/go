// errorcheck -lang=go1.17


package p

type _ interface {
	int // ERROR "embedding non-interface type int requires go1\.18 or later \(-lang was set to go1\.17; check go.mod\)"
}
