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
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := usecase.InsertTask(ctx, &task); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "OK")
	}
}

func deleteTaskById(usecase *usecase.DeleteTaskById) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		id := c.Param("id")

		if err := usecase.DeleteTaskById(ctx, id); err != nil {
			c.JSON(http.StatusNoContent, err.Error())
		}
		c.JSON(http.StatusAccepted, "OK")
	}
}

func getTaskById(usecase *usecase.SearchTaskById) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		id := c.Param("id")

		task, err := usecase.SearchTaskById(ctx, id)

		if err != nil {
			c.JSON(http.StatusNoContent, err.Error())
		}
		c.JSON(http.StatusOK, task)
	}
}

func updateTask(usecase *usecase.UpdateTask) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var task models.Task

		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := usecase.UpdateTask(ctx, &task); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, "OK")
	}
}
