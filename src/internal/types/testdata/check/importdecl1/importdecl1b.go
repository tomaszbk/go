package importdecl1

import . /* ERRORx ".unsafe. imported and not used" */ "unsafe"

type B interface {
	A
}
