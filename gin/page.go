package api

// PageRequest 分页请求参数
type PageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

// OffsetLimit 分页下行
func (p PageRequest) OffsetLimit() PageOffsetLimit {
	var response = PageOffsetLimit{
		Page:  p.Page,
		Limit: p.PageSize,
	}

	if response.Page < 1 {
		response.Page = 1
	}
	if response.Limit < 1 {
		response.Limit = 15
	}
	response.Offset = (response.Page - 1) * response.Limit

	return response
}

// PageOffsetLimit 定义分页参数返回结构体
type PageOffsetLimit struct {
	Page   int
	Limit  int
	Offset int
}

// PageResponse 分页返回结构
type PageResponse struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
