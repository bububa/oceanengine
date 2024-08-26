package datasource

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/datasource"
)

// Update 数据源更新
// 用户可以调用该接口在原有的数据源上进行添加、删除、重置操作。 数据源更新不会导致数据源id发生变化。用户可以在【数据源详细信息】查看更新是否完成。当lastest_published_time:数据源最近一次发布时间返回的数据被覆盖为最新的更新发布时间后，则说明最近一次更新已经完成。
// 添加：将需要添加的数据源文件内容通过调用【数据源文件上传】的方式获得文件路径file_path，作为【数据源更新】的请求参数file_paths添加到当前数据源内； 删除：将需要删除的数据源文件内容通过调用【数据源文件上传】的方式获得文件路径file_path，作为【数据源更新】的请求参数file_paths进行删除操作，删除掉当前数据源内这部分内容； 重置：将需要替换的新数据源内容通过调用【数据源文件上传】的方式获得文件路径file_path，作为【数据源更新】的请求参数file_paths进行重置操作，替换掉当前数据源内的内容；
// 每一个数据源一天只能更新50次,建议合理使用更新次数，减少无效的更新！
// 每次更新数据源后，都需要重新调用【发布人群包】接口发布人群包才能让对应的更新生效！否则人群包将使用更新前的数据源！
// 在调用【发布人群包】接口发布人群包的过程中，建议不要更新数据源，这会导致人群包发布过程延后！
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *datasource.UpdateRequest) error {
	return clt.Post(ctx, "2/dmp/data_source/update/", req, nil, accessToken)
}
