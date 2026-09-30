package ssacompile

// int128 is a type that stores a 128-bit constant.
// The only allowed constant right now is 0, so we can cheat quite a bit.
type int128 int64
