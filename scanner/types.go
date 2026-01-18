package scanner

type InputType interface {
	~int | ~int64 | ~float64 | ~string | ~bool
}
