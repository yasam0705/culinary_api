package pkg

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	limit  uint64
	offset uint64
	filter map[string]string
}

func GetQueryParams(c *gin.Context) (*QueryParams, error) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		return nil, err
	}
	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		return nil, err
	}
	filter := make(map[string]string)

	if err = c.ShouldBindQuery(&filter); err != nil {
		return nil, err
	}
	delete(filter, "limit")
	delete(filter, "offset")

	return &QueryParams{
		limit:  limit,
		offset: offset,
		filter: filter,
	}, nil
}

func (q *QueryParams) GetLimit() uint64 {
	return q.limit
}

func (q *QueryParams) GetOffset() uint64 {
	return q.offset
}

func (q *QueryParams) GetFilter() map[string]string {
	return q.filter
}
