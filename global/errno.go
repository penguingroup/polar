package global

// 新定义返回结果函数
func NewResp(code int, msg string) func(...interface{}) *JsonResponse {
	return func(i ...interface{}) *JsonResponse {
		length := len(i)
		jsonResponse := &JsonResponse{
			Status:  code,
			Message: msg,
		}
		if length > 0 {
			jsonResponse.Data = i[0]
		}
		return jsonResponse
	}
}

var (
	RespSuccess     = NewResp(0, "请求成功")
	RespNotAuth     = NewResp(1, "未授权")
	RespFailed      = NewResp(2, "请求失败")
	RespServerError = NewResp(3, "服务异常")
	RespNotFound    = NewResp(4, "资源未找到")
	RespOptFailed   = NewResp(5, "操作失败")
	RespNotLogin    = NewResp(6, "未登录")
	RespIllegal     = NewResp(7, "名词非法")
)
