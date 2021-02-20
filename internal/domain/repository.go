package domain

// WorkerRepository ...
type (
	WorkerRepository interface {
		GetAll() (WorkerResponse, error)
		InsertMessage() error
		GetMessage() (MessageResponse, error)
	}
)
