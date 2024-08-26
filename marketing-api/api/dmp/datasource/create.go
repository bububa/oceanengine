package datasource

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/datasource"
)

// Create 数据源创建
// 通过【数据源文件上传】接口得到file_path文件路径后，调用当前接口将数据源文件创建成一个数据源，创建成功后会返回一个数据源id，作为数据源的唯一标识。
// 数据源创建完成后，系统会进行一个数据源解析的过程，将数据源解析成对应的人群包，这个过程大概持续20-60分钟；（同时使用该功能的用户数量多时可能造成解析时间较长，超出60分钟，具体时长将由具体任务量决定，请耐心等待）。 人群包的生成进度可通过【数据源详细信息】查询。
// 目前每个广告主账号一天只能创建100次数据源，超出报错！
// 一次上传的file_paths文件路径限制个数为1000，超出报错！
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *datasource.CreateRequest) (string, error) {
	var resp datasource.CreateResponse
	err := clt.Post(ctx, "2/dmp/data_source/create/", req, &resp, accessToken)
	if err != nil {
		return "", err
	}
	return resp.Data.DataSourceID, nil
}
