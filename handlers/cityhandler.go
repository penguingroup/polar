package handlers

import (
	"encoding/json"
	"github.com/karldoenitz/Tigo/TigoWeb"
	"polar/global"
	"polar/utils/redis"
)

type CityHandler struct {
	TigoWeb.BaseHandler
}

func (c *CityHandler) Get() {
	data, isFound := redis.Get(global.CITY_REDIS_KEY)
	if !isFound {
		result := []global.City{
			{1, "beijing", "北京"},
		}
		c.ResponseAsJson(global.RespSuccess(result))
		return
	}
	var result []global.City
	if err := json.Unmarshal(data, &result); err != nil {
		c.ResponseAsJson(global.RespServerError(err.Error()))
		return
	}
	c.ResponseAsJson(global.RespSuccess(result))
	return
}
