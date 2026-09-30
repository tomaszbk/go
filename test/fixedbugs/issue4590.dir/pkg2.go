package pkg2

import "./pkg1"

var T = struct{ pkg1.A }{nil}
var U = struct{ pkg1.B }{nil}
var V pkg1.A = struct{ *pkg1.C }{nil}
var W = interface {
	Write() error
	Hello()
}(nil)
