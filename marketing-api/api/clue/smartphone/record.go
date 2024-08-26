package smartphone

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/smartphone"
)

// Record 查询智能电话拨打记录
func Record(ctx context.Context, clt *core.SDKClient, accessToken string, req *smartphone.RecordRequest) (*smartphone.RecordResponseData, error) {
	var resp smartphone.RecordResponse
	if err := clt.Get(ctx, "2/clue/smartphone/record/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
