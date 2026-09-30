package strconv

var (
	Log10Pow2        = log10Pow2
	Log2Pow10        = log2Pow10
	ParseFloatPrefix = parseFloatPrefix
)

func NewDecimal(i uint64) *decimal {
	d := new(decimal)
	d.Assign(i)
	return d
}

func SetOptimize(b bool) bool {
	old := optimize
	optimize = b
	return old
}
