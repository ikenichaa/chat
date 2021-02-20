package repository

import (
	"fmt"
	"gorm.io/gorm"
	d "practice/golang/internal/domain"
)

type workerRepository struct {
	db *gorm.DB
}

type workers struct {
	ID   int
	Name string
	Age  int
}

func (worker *workers) TableName() string {
	return "workers"

}

// NewWorkerRepository ...
func NewWorkerRepository(db *gorm.DB) *workerRepository {
	return &workerRepository{
		db: db,
	}
}

func (r *workerRepository) GetAll() (d.WorkerResponse,error) {
	var (
		user d.Worker
		userAll d.WorkerResponse
	)
	rows, err := r.db.Model(&workers{}).Rows()
	if err!=nil {
		return userAll, err
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println(rows)
		r.db.ScanRows(rows, &user)
		userAll.Data = append(userAll.Data, user)
	}
	return userAll,nil
}
