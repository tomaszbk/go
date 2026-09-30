// -lang=go1.22


package aliasTypes

type _ = int
type _[P /* ERROR "generic type alias requires go1.23 or later" */ any] = int
