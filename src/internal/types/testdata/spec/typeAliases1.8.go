// -lang=go1.8


package aliasTypes

type _ = /* ERROR "type alias requires go1.9 or later" */ int
type _[P /* ERROR "generic type alias requires go1.23 or later" */ interface{}] = int
