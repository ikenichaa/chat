package service

import (
	d "practice/golang/internal/domain"
)

// WorkerService ...
type workerService struct {
	worker d.WorkerRepository
}

type WorkerConfig struct {
	Worker d.WorkerRepository
}

// NewWorkerService ...
func NewWorkerService(w WorkerConfig) *workerService {
	return &workerService{worker: w.Worker}
}

func (w *workerService) GetAll() (d.WorkerResponse,error) {
var response = d.WorkerResponse{}
	result, err := w.worker.GetAll()
	if err!= nil {
		return response, err
	}
	return result, nil
}

func Multiply(x,y int) int {
	return x*y
}
