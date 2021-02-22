package store

type ErrNotFound struct {
	resource string
	Id       string
}

func NewErrNotFound(resource string, id string) *ErrNotFound {
	return &ErrNotFound{
		resource: resource,
		Id:       id,
	}
}

func (e *ErrNotFound) Error() string {
	return "resource: " + e.resource + " id: " + e.Id
}
