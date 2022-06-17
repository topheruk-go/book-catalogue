package bookcatalogue

type Repo[T any, U any] interface {
	ReadRepo[T, U]
	WriteRepo[T, U]
}

type ReadRepo[T any, U any] interface {
	Find(U) (T, error)
	List() ([]T, error)
	// Filter() ([]T, error)
}

type WriteRepo[T any, U any] interface {
	Insert(*T) error
	// Remove(U)error
}
