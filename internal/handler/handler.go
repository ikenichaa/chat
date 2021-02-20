package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/golang/internal/domain"
	"practice/golang/internal/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Handler ...
type Handler struct {
	worker domain.WorkerService
}

// Worker struct
type Worker struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Chat struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

type AllChat struct {
	Data []Chat `json:"data"`
}

var (
	// Workers ...
	Workers = map[int]*Worker{}
	seq     = 1
)

// NewHandler ...
func NewHandler(worker domain.WorkerService) *Handler {
	return &Handler{worker: worker}
}

// GetWorker ...
func (h *Handler) GetWorker(c echo.Context) error {
	name := c.FormValue("name")
	fmt.Println(name)
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, Workers[id])
}

//CreateWorker ...
func (h *Handler) CreateWorker(c echo.Context) error {
	u := &Worker{
		ID: seq,
	}
	fmt.Println("This is u", u)
	if err := c.Bind(u); err != nil {
		return err
	}
	Workers[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

// GetWorkers ...
func (h *Handler) GetWorkers(c echo.Context) error {
	return c.JSON(http.StatusOK, Workers)
}

// UpdateWorker ...
func (h *Handler) UpdateWorker(c echo.Context) error {
	u := new(Worker)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	Workers[id].Name = u.Name
	return c.JSON(http.StatusOK, Workers[id])
}

// DeleteWorker ...
func (h *Handler) DeleteWorker(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(Workers, id)
	return c.NoContent(http.StatusNoContent)
}

//Input ...
func (h *Handler) Input(c echo.Context) error {
	u := &Worker{
		Name: "hi",
	}
	fmt.Println("This is u", u)
	if err := c.Bind(u); err != nil {
		return err
	}
	Workers[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) GetAll(c echo.Context) error {
	data, err := h.worker.GetAll()
	if err != nil {
		return err
	}
	return response.Success(c, data.Data)
}

func (h *Handler) InsertMessage(c echo.Context) error {
	var req = Chat{}
	tmp := echo.Map{}
	c.Bind(&tmp)
	jsonData, err := json.Marshal(tmp)
	if err != nil {
		return response.Error(c, err)
	}

	err = json.Unmarshal(jsonData, &req)

	if err != nil {
		return response.Error(c, err)
	}
	fmt.Println(req.User)
	fmt.Println(req.Message)
	err = h.worker.InsertMessage()

	return c.JSON(http.StatusOK, err)
}

func (h *Handler) GetMessage(c echo.Context) error {
	data, err := h.worker.GetMessage()
	if err != nil {
		return err
	}
	return response.Success(c, data.Data)
}
