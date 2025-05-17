package core

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bububa/oceanengine/marketing-api/core/internal/debug"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient sdk client
type SDKClient struct {
	client     *http.Client
	tracer     *Otel
	limiter    RateLimiter
	Secret     string
	operatorIP string
	preReqs    []PreRequest
	AppID      uint64
	debug      bool
	sandbox    bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appID uint64, secret string) *SDKClient {
	return &SDKClient{
		AppID:  appID,
		Secret: secret,
		client: defaultHttpClient(),
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHttpClient 设置http.Client
func (c *SDKClient) SetHttpClient(client *http.Client) {
	c.client = client
}

// UseSandbox 启用sandbox
func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

// DisableSandbox 禁用sandbox
func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

// SetOperatorIP 设置操作者IP, 支持ipv4/ipv6
func (c *SDKClient) SetOperatorIP(ip string) {
	c.operatorIP = ip
}

// SetRateLimiter 设置限流
func (c *SDKClient) SetRateLimiter(limiter RateLimiter) {
	c.limiter = limiter
}

func (c *SDKClient) WithTracer(namespace string) {
	c.tracer = NewOtel(namespace, c.AppID)
}

func (c *SDKClient) WithPreRequests(reqs ...PreRequest) {
	c.preReqs = reqs
}

func (c *SDKClient) AddPreRequests(reqs ...PreRequest) {
	c.preReqs = append(c.preReqs, reqs...)
}

// Copy 复制SDKClient
func (c *SDKClient) Copy() *SDKClient {
	return &SDKClient{
		AppID:      c.AppID,
		Secret:     c.Secret,
		debug:      c.debug,
		sandbox:    c.sandbox,
		operatorIP: c.operatorIP,
		client:     c.client,
		tracer:     c.tracer,
		preReqs:    c.preReqs,
	}
}

// Post post api
func (c *SDKClient) Post(ctx context.Context, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	return c.post(ctx, BASE_URL, gw, req, resp, accessToken)
}

func (c *SDKClient) PostAPI(ctx context.Context, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	return c.post(ctx, API_BASE_URL, gw, req, resp, accessToken)
}

// Get get api
func (c *SDKClient) Get(ctx context.Context, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	return c.get(ctx, BASE_URL, gw, req, resp, accessToken)
}

func (c *SDKClient) GetAPI(ctx context.Context, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	return c.get(ctx, API_BASE_URL, gw, req, resp, accessToken)
}

// OpenGet get api
func (c *SDKClient) OpenGet(ctx context.Context, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	return c.get(ctx, OPEN_URL, gw, req, resp, accessToken)
}

// Upload multipart/form-data post
func (c *SDKClient) Upload(ctx context.Context, gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	return c.upload(ctx, BASE_URL, gw, req, resp, accessToken)
}

func (c *SDKClient) UploadAPI(ctx context.Context, gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	return c.upload(ctx, API_BASE_URL, gw, req, resp, accessToken)
}

func (c *SDKClient) post(ctx context.Context, base string, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	reqUrl := util.StringsJoin(base, gw)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.operatorIP != "" {
		httpReq.Header.Add("Operator-Ip", c.operatorIP)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	return c.WithSpan(ctx, httpReq, resp, reqBytes, c.fetch)
}

func (c *SDKClient) get(ctx context.Context, base string, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	reqUrl := util.StringsJoin(base, gw)
	if req != nil {
		reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.operatorIP != "" {
		httpReq.Header.Add("Operator-Ip", c.operatorIP)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	return c.WithSpan(ctx, httpReq, resp, nil, c.fetch)
}

// GetBytes get bytes api
func (c *SDKClient) GetBytes(ctx context.Context, gw string, req model.GetRequest, accessToken string) ([]byte, error) {
	reqUrl := util.StringsJoin(BASE_URL, gw)
	if req != nil {
		reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.operatorIP != "" {
		httpReq.Header.Add("Operator-Ip", c.operatorIP)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	var ret []byte
	err = c.WithSpan(ctx, httpReq, nil, nil, func(httpReq *http.Request, resp model.Response) (*http.Response, error) {
		httpResp, err := c.client.Do(httpReq)
		if err != nil {
			return httpResp, err
		}
		defer httpResp.Body.Close()
		ret, err = io.ReadAll(httpResp.Body)
		return httpResp, err
	})
	return ret, err
}

func (c *SDKClient) upload(ctx context.Context, base string, gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	mw := multipart.NewWriter(buf)
	params := req.Encode()
	mp := make(map[string]string, len(params))
	for _, v := range params {
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)
		if v.Reader != nil {
			if fw, err = mw.CreateFormFile(v.Key, v.Value); err != nil {
				return err
			}
			r = v.Reader
			builder := util.GetStringsBuilder()
			builder.WriteString("@")
			builder.WriteString(v.Value)
			mp[v.Key] = builder.String()
			util.PutStringsBuilder(builder)
		} else {
			if fw, err = mw.CreateFormField(v.Key); err != nil {
				return err
			}
			r = strings.NewReader(v.Value)
			mp[v.Key] = v.Value
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}
	}
	mw.Close()
	reqUrl := util.StringsJoin(base, gw)
	debug.PrintPostMultipartRequest(reqUrl, mp, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.operatorIP != "" {
		httpReq.Header.Add("Operator-Ip", c.operatorIP)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}

	bs, _ := json.Marshal(mp)
	return c.WithSpan(ctx, httpReq, resp, bs, c.fetch)
}

// TrackActive 转化回传API专用
func (c *SDKClient) TrackActive(ctx context.Context, req model.TrackRequest, resp model.Response) error {
	var (
		reqUrl   = req.RequestURI()
		reqBytes = req.Encode()
	)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if token, err := req.Sign(httpReq, nil); err == nil {
		httpReq.Header.Add("x-rs256-token", token)
	}
	if token := req.GetAppAccessToken(); token != "" {
		httpReq.Header.Add("App-Access-Token", token)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	if resp != nil {
		return c.WithSpan(ctx, httpReq, resp, reqBytes, c.fetch)
	}
	return c.WithSpan(ctx, httpReq, nil, reqBytes, func(httpReq *http.Request, resp model.Response) (*http.Response, error) {
		httpResp, err := c.client.Do(httpReq)
		if err != nil {
			return httpResp, err
		}
		defer httpResp.Body.Close()
		if httpResp.StatusCode != 200 {
			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				return httpResp, err
			}
			return httpResp, model.BaseResponse{Code: httpResp.StatusCode, Message: string(body)}
		}
		return httpResp, nil
	})
}

// AnalyticsPost 转化回传API专用
func (c *SDKClient) AnalyticsPost(ctx context.Context, gw string, req model.ConversionRequest, resp model.Response) error {
	reqBytes := req.Encode()
	reqUrl := util.StringsJoin(ANALYTICS_URL, gw)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if token := req.GetAppAccessToken(); token != "" {
		httpReq.Header.Add("App-Access-Token", token)
	}
	if token, err := req.Sign(httpReq, reqBytes); err == nil {
		httpReq.Header.Add("x-rs256-token", token)
	} else {
		reqBuf := util.GetBufferPool()
		reqBuf.Grow(len(reqBytes) + len(c.Secret))
		reqBuf.Write(reqBytes)
		reqBuf.WriteString(c.Secret)
		sha256Sum := sha256.Sum256(reqBuf.Bytes())
		util.PutBufferPool(reqBuf)
		sign := hex.EncodeToString(sha256Sum[:])
		httpReq.Header.Add("x-signature", sign)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	return c.WithSpan(ctx, httpReq, resp, reqBytes, c.fetch)
}

// AnalyticsV1Post 电话转化回传API专用
func (c *SDKClient) AnalyticsV1Post(ctx context.Context, gw string, req model.PostRequest, resp model.Response) error {
	reqBytes := req.Encode()
	reqUrl := util.StringsJoin(ANALYTICSV1_URL, gw)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	reqBuf := util.GetBufferPool()
	reqBuf.Grow(len(reqBytes) + len(c.Secret))
	reqBuf.Write(reqBytes)
	reqBuf.WriteString(c.Secret)
	sha256Sum := sha256.Sum256(reqBuf.Bytes())
	util.PutBufferPool(reqBuf)
	sign := hex.EncodeToString(sha256Sum[:])
	httpReq.Header.Add("x-signature", sign)
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	return c.WithSpan(ctx, httpReq, resp, reqBytes, c.fetch)
}

type PreRequest func(httpReq *http.Request) error

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) (*http.Response, error) {
	if len(c.preReqs) > 0 {
		for _, req := range c.preReqs {
			if err := req(httpReq); err != nil {
				return nil, err
			}
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return httpResp, err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.BaseResponse{}
	}
	body, err := debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if httpResp.ContentLength <= 0 {
		httpResp.ContentLength = int64(len(body))
	}
	if err != nil {
		debug.PrintError(err, c.debug)
		return httpResp, errors.Join(err, model.BaseResponse{
			Code:    httpResp.StatusCode,
			Message: string(body),
		})
	}
	if resp.IsError() {
		return httpResp, resp
	}
	return httpResp, nil
}

func (c *SDKClient) WithSpan(ctx context.Context, req *http.Request, resp model.Response, payload []byte, fn func(*http.Request, model.Response) (*http.Response, error)) error {
	if c.tracer == nil {
		_, err := fn(req, resp)
		return err
	}
	return c.tracer.WithSpan(ctx, req, resp, payload, fn)
}
