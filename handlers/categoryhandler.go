package handlers

import (
	"encoding/json"
	"github.com/karldoenitz/Tigo/TigoWeb"
	"polar/global"
	"polar/utils/redis"
)

type CategoryHandler struct {
	TigoWeb.BaseHandler
}

func (c *CategoryHandler) Get() {
	c.SetHeader("Access-Control-Allow-Origin", "*")
	data, isFound := redis.Get(global.CATEGORY_REDIS_KEY)
	if !isFound {
		result := []global.Category{
			{1, "hot", "热点"},
		}
		c.ResponseAsJson(global.RespSuccess(result))
		return
	}
	var result []global.Category
	if err := json.Unmarshal(data, &result); err != nil {
		c.ResponseAsJson(global.RespServerError(err.Error()))
		return
	}
	c.ResponseAsJson(global.RespSuccess(result))
	return
}
