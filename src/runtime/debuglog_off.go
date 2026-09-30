//go:build !debuglog

package runtime

const dlogEnabled = false

type dlogger = dloggerFake

func dlog1() dloggerFake {
	return dlogFake()
}

type dlogPerM struct{}

func getCachedDlogger() *dloggerImpl {
	return nil
}

func putCachedDlogger(l *dloggerImpl) bool {
	return false
}
