package middlewares

import (
	"errors"

	"github.com/Zaida-3dO/goblin/internal/services"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ts = services.NewTokenService("psql")

		accessToken := c.Request.Header["Token"][0]

		user, err := ts.GetUserFromAccessToken(accessToken)
		if err != nil {
			c.Error(errors.New(err.Message))
			c.Status(err.StatusCode)
			return
		}

		c.Set("user", user)
	}
}
