package lib

func Pick(x int) func(int) int    { return (y) => x + y }
func Generic[T any](x T) func() T { return () => x }

type hidden struct{ Value int }

func Hidden(f func(hidden) int) int { return f(hidden{Value: 13}) }
