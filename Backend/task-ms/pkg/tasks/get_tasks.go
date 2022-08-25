package tasks

import (
	"net/http"

	"github.com/SmartByt3r/Proyecto-aplicaciones-empresariales/Backend/task-ms/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTasks(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := ValidateToken(tokenString)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var tasks []models.Task

	if result := h.DB.Where("user_id = ?", userId).Find(&tasks); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tasks)
}
