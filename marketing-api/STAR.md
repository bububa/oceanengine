# 巨量星图开放平台 Golang SDK

- Oauth2 授权 (api/oauth)
  - 生成授权链接 [ Url(clt *core.SDKClient, redirectUrl string, state string, materialAuth bool) string ]
  - 获取 AccessToken [ AccessToken(clt *core.SDKClient, authCode String) (*oauth.AccessTokenResponseData, error) ]
  - 刷新 Token [ RefreshToken(clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponseData, error)]
- 账号服务
  - 广告主信息与资质管理
    - 修改广告主 [ agent.AdvertiserUpdate(clt *core.SDKClient, accessToken string, req *agent.AdvertiserUpdateRequest) (*agent.AdvertiserUpdateResponseData, error) ]
  - 代理商账号管理 (api/agent)
    - 广告主列表 [ AdvertiserSelect(clt *core.SDKClient, accessToken string, req *agent.AdvertiserSelectRequest) (*agent.AdvertiserSelectResponseData, error) ]
    - 二级代理商列表 [ ChildAgentSelect(clt *core.SDKClient, accessToken string, req *agent.ChildAgentSelectRequest) ([]uint64, error) ]
    - 获取代理商信息 [ Info(clt *core.SDKClient, accessToken string, req *agent.InfoRequest) ([]agent.Info, error) ]
- 巨量星图 (api/star)
  - 获取星图客户任务列表 [ DemandList(clt *core.SDKClient, accessToken string, req *star.DemandListRequest) (*star.DemandListResponseData, error) ]
  - 获取星图客户任务订单列表 [ DemandOrderList(clt *core.SDKClient, accessToken string, req *star.DemandOrderListRequest) (*star.DemandOrderListResponseData, error) ]
  - 获取订单投后分析报表 [ ReportOrderOverviewGet(clt *core.SDKClient, accessToken string, req *star.ReportOrderOverviewGetRequest) (*star.ReportOrderOverviewGetResponseData, error) ]
  - 获取订单投后受众报表 [ ReportOrderUserDistributionGet(clt *core.SDKClient, accessToken string, req *star.ReportOrderUserDistributionGetRequest) (*star.ReportOrderUserDistributionGetResponseData, error) ]
  - 获取星图订单投后线索 [ ClueList(clt *core.SDKClient, accessToken string, req *star.ClueListRequest) (*star.ClueListResponseData, error) ]
