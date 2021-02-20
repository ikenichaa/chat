package route

import (
	"practice/golang/internal/handler"

	"github.com/labstack/echo/v4"
)

//InitRoute ...
func InitRoute(e *echo.Echo, h *handler.Handler) {
	r := e.Group("/plate")

	worker := r.Group("/worker")
	worker.POST("/createworker", h.CreateWorker)
	worker.GET("/workers/:id", h.GetWorker)
	worker.GET("/workers", h.GetWorkers)
	worker.PUT("/workers/:id", h.UpdateWorker)
	worker.DELETE("/workers/:id", h.DeleteWorker)
	worker.POST("/input", h.Input)
	worker.GET("/get-all", h.GetAll)

	l := e.Group("/chat")
	l.POST("/message", h.InsertMessage)
	l.GET("/message", h.GetMessage)
}
