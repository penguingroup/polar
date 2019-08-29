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
	param := global.NewsRequest{}
	if err := n.CheckJsonBinding(&param); err != nil {
		n.ResponseAsJson(global.RespIllegal(err.Error()))
		return
	}
	result, total := services.GetNews(param.City, param.Category, param.Page, param.Size)
	response := utils.PageBuilder(result, param.Page, total)
	n.ResponseAsJson(global.RespSuccess(response))
	return
}
