package p

type _ interface {
	int
	(int)
	(*int)
	*([]byte)
	~(int)
	(int) | (string)
	(int) | ~(string)
	(/* ERROR unexpected ~ */ ~int)
	(int /* ERROR unexpected \| */ | /* ERROR unexpected name string */ string /* ERROR unexpected \) */ )
}
