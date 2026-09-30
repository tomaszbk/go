package p

var _ = interface{
	m()
	m /* ERROR "duplicate method" */ ()
}(nil)
