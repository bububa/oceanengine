package site

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/site"
)

// FormsList 获取落地页预约表单信息
// 通过此接口，用户可以获取橙子建站落地页中的特殊的表单类型，比如附带下载类型。
// 预约表单信息包括落地页表单ID、落地页表单位置、落地页表单名字等信息。
func FormsList(ctx context.Context, clt *core.SDKClient, accessToken string, req *site.FormsListRequest) ([]site.Form, error) {
	var resp site.FormsListResponse
	if err := clt.Get(ctx, "2/tools/site/forms/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
