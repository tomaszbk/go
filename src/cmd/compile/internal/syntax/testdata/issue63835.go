package p

func (x string) /* ERROR syntax error: unexpected \[, expected name */ []byte {
        return []byte(x)
}
