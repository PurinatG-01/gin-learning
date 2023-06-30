package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Paginator struct {
	Page  int
	Limit int
}

func (s *Paginator) Bind(c *gin.Context) error {
	str_page := c.Query("page")
	str_limit := c.Query("limit")
	var page, limit int
	if str_page == "" {
		page = 1
	} else {
		conv_page, page_err := strconv.Atoi(str_page)
		page = conv_page
		if page_err != nil {
			return page_err
		}
	}
	if str_limit == "" {
		limit = 10
	} else {
		conv_limit, limit_err := strconv.Atoi(str_limit)
		limit = conv_limit
		if limit_err != nil {
			return limit_err
		}
	}
	s.Page = page
	s.Limit = limit
	return nil
}
