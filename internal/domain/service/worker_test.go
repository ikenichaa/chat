package service
import (
	"gotest.tools/assert"
	d "practice/golang/internal/domain"
	"practice/golang/internal/mocks"
	"testing"
)

func TestGetAll(t *testing.T) {
	type testCase struct {
		name  string
		prepareStub func(t *testCase, customer *mocks.WorkerRepository)
		expect d.WorkerResponse
	}
	worker := d.Worker{
		ID: 1,
		Name:"Nicha",
		Age: 20,
	}
	testCases := []testCase{
		{
			name: "success case",
			prepareStub: func(t *testCase, customer *mocks.WorkerRepository) {
				response := d.WorkerResponse{
					Data: []d.Worker{
						worker,
					},
				}
				customer.On("GetAll").Return(response, nil)
			},
			expect: d.WorkerResponse{
				Data: []d.Worker{
					worker,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			workRepo := new(mocks.WorkerRepository)
			tc.prepareStub(&tc, workRepo)
			expect := tc.expect
			workService := NewWorkerService(WorkerConfig{
				Worker: workRepo,
			})
			result, _ := workService.GetAll()
			assert.DeepEqual(t, expect, result)
		})
	}
}

func TestMultiply(t *testing.T) {
	get := Multiply(10,20)
	expect := 200
	if get!=expect {
		t.Error("Multiply wrong")
	}
}