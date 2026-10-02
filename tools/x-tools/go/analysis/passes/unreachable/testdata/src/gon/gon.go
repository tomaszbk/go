package gon
var a func() = func(){
 return
 println(1) // want "unreachable code"
}
var b func() = () => {
 return
 println(1) // want "unreachable code"
}
