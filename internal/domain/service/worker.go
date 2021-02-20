package service

import (
	"fmt"
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

func (w *workerService) GetAll() (d.WorkerResponse, error) {
	var response = d.WorkerResponse{}
	result, err := w.worker.GetAll()
	if err != nil {
		return response, err
	}
	return result, nil
}

func (w *workerService) InsertMessage() error {
	fmt.Println("in service")
	err := w.worker.InsertMessage()
	if err != nil {
		return err
	}
	return nil
}

func Multiply(x, y int) int {
	return x * y
}
