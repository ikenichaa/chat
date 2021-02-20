package domain

// Worker struct
type Worker struct {
	ID   int
	Name string
	Age  int
}

type WorkerResponse struct {
	Data []Worker `json:"worker"`
}

type InsertMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
