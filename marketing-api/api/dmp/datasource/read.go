package datasource

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/datasource"
)

// Read 数据源详细信息
// 通过数据源id，查询该数据源相关信息和其对应的人群包信息。数据源解析完成后会返回default_audience：数据源对应的人群包相关信息，其中人群包ID是该人群包的唯一标识。解析未完成不返回default_audience：数据源对应的人群包相关信息。
// 数据源解析成人群包后，调用该接口会返回该数据源对应的人群包的人群包id，作为该人群包的标识；
// 数据源在创建后被解析成人群包需要20-60分钟。高峰期可能造成解析时间延长，请您耐心等待；
// 本接口可以批量查询多个数据源的信息。其中应答参数是按create_time：数据源创建时间倒序返回的。如果您有数据处理的需求，可以参考此排序逻辑。
func Read(ctx context.Context, clt *core.SDKClient, accessToken string, req *datasource.ReadRequest) ([]datasource.DataSource, error) {
	var resp datasource.ReadResponse
	err := clt.Get(ctx, "2/dmp/data_source/read/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.DataList, nil
}
