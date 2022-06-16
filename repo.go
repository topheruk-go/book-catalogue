package bookcatalogue

type Repo interface {
	ReadRepo
	WriteRepo
}

type ReadRepo interface {
	Find(id int) (any, error)
	List() ([]any, error)
}

type WriteRepo interface {
	Insert(any) error
}
