package handlers

import (
	"github.com/karldoenitz/Tigo/TigoWeb"
	"polar/global"
	"polar/services"
	"polar/utils"
)

type NewsHandler struct {
	TigoWeb.BaseHandler
}

func (n *NewsHandler) Get() {
	param := &global.NewsRequest{}
	if err := n.CheckUrlParamBinding(param); err != nil {
		n.ResponseAsJson(global.RespIllegal(err.Error()))
		return
	}
	result, _, err := services.GetNews(param.City, param.Category, param.Page, param.Size)
	if err != nil {
		n.ResponseAsJson(global.RespServerError(err.Error()))
		return
	}
	response := utils.PageBuilder(result, param.Page, param.Size)
	n.ResponseAsJson(global.RespSuccess(response))
	return
}
