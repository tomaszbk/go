package p

type A[P int] = struct{}

var _ A[string /* ERROR "string does not satisfy int (string missing in int)" */]
