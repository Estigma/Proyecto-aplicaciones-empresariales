package tasks

import (
	"net/http"

	"github.com/SmartByt3r/Proyecto-aplicaciones-empresariales/Backend/task-ms/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&task)

	c.Status(http.StatusOK)
}
