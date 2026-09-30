// errorcheck

package main

type T int

func (T) m() {} // GCCGO_ERROR "previous"
func (T) m() {} // ERROR "T\.m already declared|redefinition"

func (*T) p() {} // GCCGO_ERROR "previous"
func (*T) p() {} // ERROR "T\.p already declared|redefinition"
