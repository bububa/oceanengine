# 巨量引擎MarketingAPI Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/oceanengine.svg)](https://pkg.go.dev/github.com/bububa/oceanengine)
[![Go](https://github.com/bububa/oceanengine/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/oceanengine/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/oceanengine/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/oceanengine/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/oceanengine.svg)](https://github.com/bububa/oceanengine)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/oceanengine)](https://goreportcard.com/report/github.com/bububa/oceanengine)
[![GitHub license](https://img.shields.io/github/license/bububa/oceanengine.svg)](https://github.com/bububa/oceanengine/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/oceanengine.svg)](https://GitHub.com/bububa/oceanengine/releases/)

- Oauth2 授权 (api/oauth)
  - 生成授权链接 [ Url(clt *core.SDKClient, redirectUrl string, state string) string ]
  - 获取AccessToken [ AccessToken(clt *core.SDKClient, authCode String) (*oauth.AccessTokenResponse, error) ]
  - 刷新Token [ RefreshToken(clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponse, error)]
  - 获取已授权账户 [ Advertiser(clt *core.SDKClient, accessToken string) (*oauth.AdvertiserGetResponse, error) ]
  - 获取授权User信息 [ UserInfo(clt *core.SDKClient, accessToken string) (*oauth.UserInfoResponse, error) ]
- 账号服务
  - 广告主信息与资质管理 (api/advertiser)
    - 广告主信息 [ Info(clt *core.SDKClient, accessToken string, req *advertiser.InfoRequest) (*advertiser.InfoResponse, error) ]
    - 广告主公开信息 [ PublicInfo(clt *core.SDKClient, accessToken string, req *advertiser.PublicInfoRequest) (*advertiser.PublicInfoResponse, error) ]
    - 获取广告主头像信息 [ Avatar(clt *core.SDK, accessToken string, advertiserID uint64) (*advertiser.AvatarGetResponse, error) ] 
  - 代理商账号管理 (api/agent)
    - 广告主列表 [ AdvertiserSelect(clt *core.SDKClient, accessToken string, req *agent.AdvertiserSelectRequest) (*agent.AdvertiserSelectResponse, error) ] 
    - 修改广告主 [ AdvertiserUpdate(clt *core.SDKClient, accessToken string, req *agent.AdvertiserUpdateRequest) (*agent.AdvertiserUpdateResponse, error) ]
    - 二级代理商列表 [ ChildAgentSelect(clt *core.SDKClient, accessToken string, req *agent.ChildAgentSelectRequest) (*agent.ChildAgentSelectResponse, error) ]
    - 获取代理商信息 [ Info(clt *core.SDKClient, accessToken string, req *agent.InfoRequest) (*agent.InfoResponse, error) ] 
  - 账号管家管理 (api/majordomo)
    - 广告主列表 [ MajordomoSelect(clt *core.SDKClient, accessToken string, req *majordomo.AdvertiserSelectRequest) (*majordomo.AdvertiserSelectResponse, error) ]
  - 资金和流水管理 (api)
    - 查询账号余额 [ advertiser.Fund(clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.FundGetResponse, error) ]
    - 查询账号日流水 [ advertiser.FundDailyStat(clt *core.SDKClient, accessToken string, req *advertiser.FundDailyStatRequest) (*advertiser.FundDailyStatResponse, error) ]
    - 查询账号流水明细 [ advertiser.FundTransaction(clt *core.SDKClient, accessToken string, req *advertiser.FundTransactionGetRequest) (*advertiser.FundTransactionGetResponse, error) ]
    - 代理商转账 [ agent.AdvertiserRecharge(clt *core.SDKClient, accessToken string, req *agent.AdvertiserRechargeRequest) (*agent.AdvertiserRechargeResponse, error) ]
    - 代理商退款 [ agent.AdvertiserRefund(clt *core.SDKClient, accessToken string, req *agent.AdvertiserRefundRequest) (*agent.AdvertiserRefundResponse, error) ]
- 广告投放
  - 广告账户预算 (api)
    - 获取账户日预算 [ advertiser.BudgetGet(clt *core.SDKClient, accessToken string, req *advertiser.BudgetGetRequest) (*advertiser.BudgetGetResponse, error) ]
    - 更新账户日预算 [ advertiser.UpdateBudget(clt *core.SDKClient, accessToken string, req *advertiser.UpdateBudgetRequest) (*model.BaseResponse, error) ]
  - 广告组 (api/campaign)
    - 获取广告组 [ Get(clt *core.SDKClient, accessToken string, req *campaign.GetRequest) (*campaign.GetResponse, error) ]
    - 创建广告组 [ Create(clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (*campaign.CreateResponse, error) ]
    - 修改广告组 [ Update(clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (*campaign.UpdateResponse, error) ]
    - 广告组更新状态 [ UpdateStatus(clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) (*campaign.UpdateStatusResponse, error) ]
  - 广告计划模块 (api/ad)
    - 获取广告计划 [ Get(clt *core.SDKClient, accessToken string, req *ad.GetRequest) (*ad.GetResponse, error) ]
    - 更新计划状态 [ UpdateStatus(clt *core.SDKClient, accessToken string, req *ad.UpdateStatusRequest) (*ad.UpdateResponse, error) ]
    - 更新计划预算 [ UpdateBudget(clt *core.SDKClient, accessToken string, req *ad.UpdateBudgetRequest) (*ad.UpdateBudgetResponse, error) ]
    - 更新计划出价 [ UpdateBid(clt *core.SDKClient, accessToken string, req *ad.UpdateBidRequest) (*ad.UpdateBidResponse, error) ]
- 数据报表
  - 广告数据报表 (api/report)
    - 广告主数据 [ Advertiser(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponse, error) ]
    - 广告组数据 [ Campaign(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponse, error) ]
    - 广告计划数据 [ Ad(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponse, error) ]
    - 广告创意数据 [ Creative(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponse, error) ]
    - 多合一数据报表接口 [ Integrated(clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponse, error) ]
    - 视频素材报表 [ Video(clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponse, error) ]
    - 视频互动流失数据 [ VideoFrame(clt *core.SDKClient, accessToken string, req *report.VideoFrameRequest) (*report.VideoFrameResponse, error) ]
    - 分级模糊数据 [ Misty(clt *core.SDKClient, accessToken string, req *report.VideoFrameRequest) (*report.VideoFrameResponse, error) ]
  - 受众分析数据报表 (api/report/audience)
    - 行为兴趣数据 [ InterestActionList(clt *core.SDKClient, accessToken string, req *audience.ListRequest) (*audience.ListResponse, error) ]
    - 抖音达人数据 [ AwemeList(clt *core.SDKClient, accessToken string, req *audience.ListRequest) (*audience.ListResponse, error) ] 
    - 省级数据 [ Province(clt *core.SDKClient, accessToken string, req *audience.Request) (*audience.Response, error) ]
    - 市级数据 [ City(clt *core.SDKClient, accessToken string, req *audience.Request) (*audience.Response, error) ]
    - 性别数据 [ Gender(clt *core.SDKClient, accessToken string, req *audience.Request) (*audience.Response, error) ]
    - 兴趣数据 [ Tag(clt *core.SDKClient, accessToken string, req *audience.Request) (*audience.Response, error) ]
    - 年龄数据 [ Age(clt *core.SDKClient, accessToken string, req *audience.Request) (*audience.Response, error) ]
  - 电商直播数据报表 (api/report/liveroom)
    - 直播间属性报表 [ Attribute(clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.Response, error) ]
    - 直播间分析报表 [ Analysis(clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.Response, error) ]
    - 直播间流量来源报表 [ FlowCategory(clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.Response, error) ]
    - 直播间商品分析报表 [ Product(clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.Response, error) ]
    - 直播受众分析报表 [ AudiencePortrait(clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.Response, error) ]
- DMP人群管理 (api/dmp)
  - 数据源文件上传 [ datasource.FileUpload(clt *core.SDKClient, accessToken string, req *datasource.FileUploadRequest) (string, error) ]
  - 数据源创建 [ datasource.Create(clt *core.SDKClient, accessToken string, req *datasource.CreateRequest) (string, error) ] 
  - 数据源更新 [ datasource.Update(clt *core.SDKClient, accessToken string, req *datasource.UpdateRequest) error ] 
  - 数据源详细信息 [ datasource.Read(clt *core.SDKClient, accessToken string, req *datasource.ReadRequest) ([]datasource.DataSource, error) ]
  - 人群包列表 [ customaudience.Select(clt *core.SDKClient, accessToken string, req *customaudience.SelectRequest) (*customaudience.SelectResponseData, error) ]
  - 人群包详细信息 [ customaudience.Read(clt *core.SDKClient, accessToken string, req *customaudience.ReadRequest) ([]customaudience.CustomAudience, error) ]
  - 发布人群包 [ customaudience.Publish(clt *core.SDKClient, accessToken string, req *customaudience.PublishRequest) error ]
  - 推送人群包 [ customaudience.Push(clt *core.SDKClient, accessToken string, req *customaudience.PushRequest) error ]
  - 删除人群包 [ customaudience.Delete(clt *core.SDKClient, accessToken string, req *customaudience.DeleteRequest) error ]
- 素材管理 (api/file)
  - 上传广告主图片 [ ImageAdvertiser(clt *core.SDKClient, accessToken string, req *file.ImageAdvertiserRequest) (*file.Image, error) ]
  - 上传广告图片 [ ImageAd(clt *core.SDKClient, accessToken string, req *file.ImageAdRequest) (*file.Image, error) ]
  - 上传视频 [ VideoAd(clt *core.SDKClient, accessToken string, req *file.VideoAdRequest) (*file.Video, error) ]
  - 获取图片素材 [ ImageGet(clt *core.SDKClient, accessToken string, req *file.ImageGetRequest) (*file.ImageGetResponseData, error) ]
  - 获取视频素材 [ VideoGet(clt *core.SDKClient, accessToken string, req *file.VideoGetRequest) (*file.VideoGetResponseData, error) ]
  - 获取视频智能封面 [ VideoCoverSuggest(clt *core.SDKClient, accessToken string, req *file.VideoCoverSuggestRequest) (*file.VideoCoverSuggestResponseData, error) ]
  - 获取同主体下广告主图片素材 [ ImageAdGet(clt *core.SDKClient, accessToken string, req *file.ImageAdGetRequest) ([]file.Image, error) ]
  - 获取同主体下广告主视频素材 [ VideoAdGet(clt *core.SDKClient, accessToken string, req *file.VideoAdGetRequest) ([]file.Video, error) ]
  - 素材推送 [ MaterialBind(clt *core.SDKClient, accessToken string, req *file.MaterialBindRequest) ([]file.FailedMaterialBind, error) ]
  - 批量删除视频素材 [ VideoDelete(clt *core.SDKClient, accessToken string, req *file.VideoDeleteRequest) ([]string, error) ]
  - 更新视频 [ VideoUpdate(clt *core.SDKClient, accessToken string, req *file.VideoUpdateRequest) ([]file.VideoForUpdate, error) ] 
- 数据上报管理 (api/track)
  - 转化回传 [ Active(req *track.ActiveRequest) error ]
- 事件管理(api/conversion)
  - 转化回传 [ Conversion(req *conversion.Request) error ]
