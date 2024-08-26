package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// BookingRecordsGet 查询应用预约记录
// 查询应用预约记录接口
// 其中应答参数是按creative_time正序返回的。如果您有数据处理的需求，可以参考此排序逻辑。
func BookingRecordsGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.BookingRecordsGetRequest) (*appmanagement.BookingRecordsGetData, error) {
	var resp appmanagement.BookingRecordsGetResponse
	if err := clt.Get(ctx, "2/tools/app_management/booking_records/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
