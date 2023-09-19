# 巨量应用市场 Golang SDK

- 连山OpenAPI
  - 视频前置机审/投前分析API
    - 连山视频投前分析提交 [ file.v3.QualitySubmit(clt *core.SDKClient, accessToken string, req *v3.QualitySubmitRequest) (uint64, error) ]
    - 连山投前分析结果查询 [ file.v3.QualityGet(clt *core.SDKClient, accessToken string, req *v3.QualityGetRequest) ([]v3.MaterialQuality, error) ]
  - RDS订阅adv_id增删查 (api/subscribe)
    - 新增 Adv 订阅 [ AccountsAdd(clt *core.SDKClient, accessToken string, req *subscribe.AccountsAddRequest) error ]
    - 删除 Adv 订阅 [ AccountsRemove(clt *core.SDKClient, accessToken string, req *subscribe.AccountsRemoveRequest) error ]
    - 查询订阅 Adv [ AccountsList(clt *core.SDKClient, accessToken string, req *subscribe.AccountsListRequest) (*subscribe.AccountsListResult, error) ]
- 应用市场(api/serve_market)
  - 获取应用订单数据 [ OrderGet(clt *core.SDKClient, accessToken string, req *servemarket.OrderGetRequest) (*servemarket.OrderGetResponseData, error) ]
  - 获取用户已购功能点列表 [ ActiveFuncGet(clt *core.SDKClient, accessToken string, req *servemarket.ActiveFuncGetRequest) ([]servemarket.OrderFunction, error) ]
