package tools

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// Pagination return limit and skip
func Pagination(c *gin.Context) (uint64, uint64, error) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return 0, 0, err
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.WithField("component", "service").Debug(err)
		return 0, 0, err
	}

	skip := page*limit - limit
	uLimit := uint64(limit)
	uSkip := uint64(skip)
	return uLimit, uSkip, nil
}
