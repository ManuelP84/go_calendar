package task

import (
	"net/http"

	"github.com/ManuelP84/calendar/business/task/usecase"
	"github.com/gin-gonic/gin"
)

func getAllTasks(usecase *usecase.GetAllTasks) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		tasks, error := usecase.GetAllTasks(ctx)

		if error != nil {
			c.JSON(http.StatusInternalServerError, error.Error())
			return
		}

		c.JSON(http.StatusOK, tasks)
	}
}
