package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context, err error, message string) {
	c.JSON(
		http.StatusInternalServerError,
		gin.H{
			"msg": message,
			"err": err.Error(),
		},
	)
}

func BadRequest(c *gin.Context, err error, target string) {
	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"msg": "failed to parse request " + target,
			"err": err.Error(),
		},
	)
}
