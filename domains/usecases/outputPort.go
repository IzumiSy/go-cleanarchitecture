package usecases

type OutputPort interface {
	Raise(reason error)
}
