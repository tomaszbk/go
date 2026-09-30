// compile


package bug066

type Scope struct {
	entries map[string] *Object;
}


type Type struct {
	scope *Scope;
}


type Object struct {
	typ *Type;
}


func Lookup(scope *Scope) *Object {
	return scope.entries["foo"];
}
