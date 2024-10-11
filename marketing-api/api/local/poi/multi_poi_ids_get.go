package poi

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/poi"
)

// MultiPoiIDsGet 根据多门店ID拉取门店ID
func MultiPoiIDsGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *poi.MultiPoiIDsGetRequest) (*poi.MultiPoiIDsGetResult, error) {
	var resp poi.MultiPoiIDsGetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/multi_poi_id/poi_ids/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
