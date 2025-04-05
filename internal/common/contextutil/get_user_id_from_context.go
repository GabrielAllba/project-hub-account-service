package contextutil

import "github.com/gin-gonic/gin"

func GetUserIDFromContext(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, false
	}

	id, ok := userID.(uint)
	return id, ok
}
