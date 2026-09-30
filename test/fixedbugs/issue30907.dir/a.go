package a

type UUID string

func New() UUID {
	return Must(NewRandom())
}

func NewRandom() (UUID, error) {
	return "", nil
}

func Must(uuid UUID, err error) UUID {
	return uuid
}
