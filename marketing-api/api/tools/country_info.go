package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// CountryInfo 查询国家/区域信息
func CountryInfo(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.CountryInfoRequest) (*tools.CountryInfoResponseData, error) {
	var resp tools.CountryInfoResponse
	if err := clt.Get(ctx, "2/tools/country/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
