package middlewares

import (
	"errors"
	"net/http"

	"github.com/Zaida-3dO/goblin/internal/services"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ts = services.NewTokenService("psql")

		accessTokenSlice := c.Request.Header["Token"]
		if len(accessTokenSlice) == 0 {
			c.Error(errors.New("unauthorized access"))
			c.Status(http.StatusUnauthorized)
			return
		}

		user, err := ts.GetUserFromAccessToken(accessTokenSlice[0])
		if err != nil {
			c.Error(errors.New(err.Message))
			c.Status(err.StatusCode)
			return
		}

		c.Set("user", user)
	}
}
