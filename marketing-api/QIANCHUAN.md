# 巨量千川开放平台 Golang SDK

- Oauth2 授权 (api/oauth)
  - 生成授权链接 [ Url(clt *core.SDKClient, redirectUrl string, state string, materialAuth bool) string ]
  - 获取 AccessToken [ AccessToken(clt *core.SDKClient, authCode String) (*oauth.AccessTokenResponseData, error) ]
  - 刷新 Token [ RefreshToken(clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponseData, error)]

- 账户关系获取
  - 获取已授权的账户（店铺/代理商）[ oauth.AdvertiserGet(clt *core.SDKClient, accessToken string) ([]oauth.Advertiser, error) ]
  - 获取店铺账户关联的广告账户列表 [ qianchuan.shop.AdvertiserList(clt *core.SDKClient, accessToken string, req *shop.AdvertiserListRequest) (*shop.AdvertiserListResponseData, error) ]
  - 获取代理商账户关联的广告账户列表 [ advertiser.AdvertiserSelect(clt *core.SDKClient, accessToken string, req *agent.AdvertiserSelectRequest) (*agent.AdvertiserSelectResponseData, error) ]

- 账户信息获取
  - 获取授权 User 信息 [ oauth.UserInfo(clt *core.SDKClient, accessToken string) (*oauth.UserInfoResponseData, error) ]
  - 获取代理商信息 [ agent.Info(clt *core.SDKClient, accessToken string, req *agent.InfoRequest) ([]agent.Info, error) ]
  - 获取店铺账户信息 [ Get(clt *core.SDKClient, accessToken string, req *shop.GetRequest) ([]shop.Shop, error) ]
  - 广告主公开信息 [ advertiser.PublicInfo(clt *core.SDKClient, accessToken string, req *advertiser.PublicInfoRequest) ([]advertiser.PublicInfo, error) ]
  - 广告主信息 [ advertiser.Info(clt *core.SDKClient, accessToken string, req *advertiser.InfoRequest) ([]advertiser.Info, error) ]

- 投放管理 (api/qianchuan)
  - 广告组管理 (api/qianchuan/campaign)
    - 广告组创建 [ Create(clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (uint64, error) ]
    - 广告组更新 [ Update(clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (uint64, error) ]
    - 广告组状态更新 [ UpdateStatus(clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) (*campaign.UpdateStatusResponseData, error) ]
    - 广告组列表获取 [ ListGet(clt *core.SDKClient, accessToken string, req *campaign.ListGetRequest) (*campaign.ListGetResponseData, error) ]
  - 广告计划管理 (api/qianchuan/ad)
    - 创建广告计划 [ Create(clt *core.SDKClient, accessToken string, req *ad.CreateRequest) (uint64, error) ]
    - 更新广告计划 [ Update(clt *core.SDKClient, accessToken string, req *ad.UpdateRequest) (*ad.UpdateResponseData, error) ] 
    - 获取计划详情 [ DetailGet(clt *core.SDKClient, accessToken string, req *ad.DetailGetRequest) (*ad.Ad, error) ]
    - 获取账户下计划列表 [ Get(clt *core.SDKClient, accessToken string, req *ad.GetRequest) (*ad.GetResponseData, error) ]
    - 更新状态 [ UpdateStatus(clt *core.SDKClient, accessToken string, req *ad.UpdateStatusRequest) (*ad.UpdateResponseData, error) ]
    - 更新出价 [ UpdateBid(clt *core.SDKClient, accessToken string, req *ad.UpdateBidRequest) (*ad.UpdateResponseData, error) ]
    - 更新预算 [ UpdateBudget(clt *core.SDKClient, accessToken string, req *ad.UpdateBudgetRequest) (*ad.UpdateResponseData, error) ]
    - 获取计划审核建议 [ RejectReason(clt *core.SDKClient, accessToken string, req *ad.RejectReasonRequest) ([]ad.RejectReasonList, error) ]
  - 广告创意管理 (api/qianchuan/creative)
    - 批量更新广告创意状态 [ UpdateStatus(clt *core.SDKClient, accessToken string, req *creative.UpdateStatusRequest) (*creative.UpdateResponseData, error) ]
    - 获取账户下创意列表 [ Get(clt *core.SDKClient, accessToken string, req *creative.GetRequest) (*creative.GetResponseData, error) ]
    - 获取计划审核建议 [ RejectReason(clt *core.SDKClient, accessToken string, req *creative.RejectReasonRequest) ([]creative.RejectReasonList, error) ]
  - 商品/直播间管理
    - 获取可投商品列表 [ AvailableGet(clt *core.SDKClient, accessToken string, req *product.AvailableGetRequest) (*product.AvailableGetResponseData, error) ]
    - 获取千川账户下已授权抖音号 [ AuthorizedGet(clt *core.SDKClient, accessToken string, req *aweme.AuthorizedGetRequest) (*aweme.AuthorizedGetResponseData, error) ]

- 数据报表 （api/qianchuan/report)
  - 获取广告账户数据 [ AdvertiserGet(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) ]
  - 获取广告计划数据 [ AdGet(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) ]
  - 获取广告创意数据 [ CreativeGet(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) ]

- 素材管理 (api/file)
  - 上传广告图片 [ ImageAd(clt *core.SDKClient, accessToken string, req *file.ImageAdRequest) (*file.Image, error) ]
  - 上传视频 [ VideoAd(clt *core.SDKClient, accessToken string, req *file.VideoAdRequest) (*file.Video, error) ]
  - 获取图片素材 [ ImageGet(clt *core.SDKClient, accessToken string, req *file.ImageGetRequest) (*file.ImageGetResponseData, error) ]
  - 获取视频素材 [ VideoGet(clt *core.SDKClient, accessToken string, req *file.VideoGetRequest) (*file.VideoGetResponseData, error) ]
  - 获取抖音号下的视频 [ VideoAwemeGet(clt *core.SDKClient, accessToken string, req *file.VideoAwemeGetRequest) (*file.VideoAwemeGetResponseData, error) ]

- 工具
  - 查询工具
    - 获取行业列表 [ tools.IndustryGet(clt *core.SDKClient, accessToken string, req *tools.IndustryGetRequest) ([]tools.Industry, error) ]
  - 抖音达人 (tools/aweme)
    - 查询抖音类目列表 [ AwemeMultiLevelCategoryGet(clt *core.SDKClient, accessToken string, req *aweme.AwemeMultiLevelCategoryGetRequest) ([]aweme.Category, error) ]
    - 查询抖音类目下的推荐达人 [ AwemeCategoryTopAuthorGet(clt *core.SDKClient, accessToken string, req *aweme.AwemeCategoryTopAuthorGetRequest) ([]aweme.Author, error) ]
  - 行为兴趣词管理 (tools/interestaction)
    - 行为类目查询 [ ActionCategory(clt *core.SDKClient, accessToken string, req *interestaction.ActionCategoryRequest) ([]interestaction.Object, error) ]
    - 行为关键词查询 [ ActionKeyword(clt *core.SDKClient, accessToken string, req *interestaction.ActionKeywordRequest) ([]interestaction.Object, error) ]
    - 兴趣类目查询 [ InterestCategory(clt *core.SDKClient, accessToken string, req *interestaction.InterestCategoryRequest) ([]interestaction.Object, error) ]
    - 兴趣关键词查询 [ InterestKeyword(clt *core.SDKClient, accessToken string, req *interestaction.InterestKeywordRequest) ([]interestaction.Object, error) ]
  - 动态创意词包管理 (tools/creativeword)
    - 查询动态创意词包 [ Select(clt *core.SDKClient, accessToken string, req *creativeword.SelectRequest) ([]creativeword.CreativeWord, error) ]
  - DMP人群管理 (tools/dmp)
    - 查询人群包列表 [ AudiencesGet(clt *core.SDKClient, accessToken string, req *dmp.AudiencesGetRequest) (*dmp.AudiencesGetResponseData, error) ]
