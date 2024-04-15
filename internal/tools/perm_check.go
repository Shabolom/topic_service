package tools

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func PermCheck(c *gin.Context) error {
	claims, err := ParseTokenClaims(c)
	if err != nil {
		return err
	}

	if claims.UserPerm != 3 {
		return errors.New("не достаточно прав")
	}

	return nil
}
