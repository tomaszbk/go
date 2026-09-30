// -lang=go1.17

package p

func f[P /* ERROR "requires go1.18" */ interface{}](P) {}

var v func(int) = f /* ERROR "requires go1.18" */
