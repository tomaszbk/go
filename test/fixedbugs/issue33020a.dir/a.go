package a

type FArg func(args []string) error

type Command struct {
	Name string
	Arg1 FArg
	Arg2 func(args []string) error
}
