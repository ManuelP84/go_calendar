package task

import (
	"net/http"

	"github.com/ManuelP84/calendar/business/task/usecase"
	"github.com/ManuelP84/calendar/domain/task/models"
	"github.com/gin-gonic/gin"
)

func getAllTasks(usecase *usecase.GetAllTasks) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		tasks, err := usecase.GetAllTasks(ctx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, tasks)
	}
}

func insertTask(usecase *usecase.InsertTask) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var task models.Task

		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if err := usecase.InsertTask(ctx, &task); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "OK")
	}
}
