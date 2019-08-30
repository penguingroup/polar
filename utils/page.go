package utils

import "reflect"

// PageBuilder 分页生成器
//  - data 任意类型数组，本页的有效数据
//  - page 页码
//  - total 总条数
func PageBuilder(data interface{}, page, total int64) map[string]interface{} {
	result := map[string]interface{}{}
	vField := reflect.ValueOf(data)
	if vField.Kind() == reflect.Slice || vField.Kind() == reflect.Array {
		pageSize := vField.Len()
		result["page_info"] = map[string]int64{
			"page":  page,            // 当前页数
			"count": int64(pageSize), // 当前每页请求的条数
			"size":  total,           // 总数
		}
		result["data_list"] = data
	}
	return result
}
