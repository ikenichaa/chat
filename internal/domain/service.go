package domain

type (
	// WorkerService ...
	WorkerService interface {
		GetAll() (WorkerResponse, error)
		GetMessage() (MessageResponse, error)
		InsertMessage() error
	}
)
