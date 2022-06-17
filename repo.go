package bookcatalogue

type Repository[T any, U any] interface {
	ReadRepository[T, U]
	WriteRepository[T, U]
}

type ReadRepository[T any, U any] interface {
	Find(U) (T, error)
	List() ([]T, error)
	ListByIndex(U) ([]T, error)
	// Filter() ([]T, error)
}

type WriteRepository[T any, U any] interface {
	Insert(*T) error
	// Remove(U)error
}
