package tasks

import (
	"net/http"

	"github.com/SmartByt3r/Proyecto-aplicaciones-empresariales/Backend/task-ms/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskRequestBody struct {
	Title       string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func (h handler) AddTask(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := ValidateToken(tokenString)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	body := AddTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	task.Title = body.Title
	task.Status = body.Status
	task.Description = body.Description
	task.UserId = userId

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &task)
}
