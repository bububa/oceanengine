package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// ShareAccountList 查看应用共享范围
// 查询应用共享范围，查询巨量纵横组织下某个应用的共享范围
func ShareAccountList(clt *core.SDKClient, accessToken string, req *appmanagement.ShareAccountListRequest) (*appmanagement.ShareAccountListData, error) {
	var resp appmanagement.ShareAccountListResponse
	if err := clt.Get("2/tools/app_management/share_account/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
