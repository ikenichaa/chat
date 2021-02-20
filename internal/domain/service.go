package domain

type (
	// WorkerService ...
	WorkerService interface {
		GetAll() (WorkerResponse, error)
		InsertMessage() error
	}
)
