package client

import (
	"errors"
	"fmt"
	"time"

	gHttpClient "github.com/872409/ghttpclient"
)

const (
	URLAccessToken = "/oauth/2.0/token"
)

type JSONResponseBodyInterface interface {
	SetError(err error)
	AfterResponseJSON(resp *gHttpClient.JSONResponse)
}

func (j *AccessTokenResp) SetError(err error) {
	j.Error = err
}

func (j *AccessTokenResp) AfterResponseJSON(resp *gHttpClient.JSONResponse) {
	j.ResponseAt = time.Now()
	j.ResponseBody = resp.GetBodyText()

	if j.Error != nil {
		j.ErrorCode = j.Error.Error()
		j.ErrorDescription = j.Error.Error()
	} else if j.ErrorCode != "" {
		j.Error = errors.New(j.ErrorCode + ":" + j.ErrorDescription)
	}

	if j.Error != nil {
		j.ExpiresAt = time.Now().Add(time.Duration(j.ExpiresIn-120) * time.Second)
	}
}

type AccessTokenResp struct {
	Error            error     `json:"-"`
	ErrorCode        string    `json:"error"`
	ErrorDescription string    `json:"error_description"`
	ResponseBody     string    `json:"-"`
	ResponseAt       time.Time `json:"-"`

	RefreshToken  string `json:"refresh_token"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
	ExpiresIn     int    `json:"expires_in"`

	ExpiresAt time.Time `json:"-"`
}

type Config struct {
	ClientId     string
	ClientSecret string
}

func NewHttpClient(config Config) (*BaiduClient, error) {
	client, err := gHttpClient.New("https://aip.baidubce.com", config.ClientId, config.ClientSecret)

	return &BaiduClient{
		config: config,
		client: client,
	}, err
}

type BaiduClient struct {
	client          *gHttpClient.Client
	config          Config
	accessTokenResp *AccessTokenResp
}

func (b *BaiduClient) GetAccessToken() *AccessTokenResp {
	if b.accessTokenResp != nil && b.accessTokenResp.ExpiresAt.After(time.Now()) {
		return b.accessTokenResp
	}

	urlParam := make(map[string]interface{})
	urlParam["grant_type"] = "client_credentials"
	urlParam["client_id"] = b.config.ClientId
	urlParam["client_secret"] = b.config.ClientSecret

	body := &AccessTokenResp{}

	b.POST(URLAccessToken, urlParam, nil, body)
	if body.Error != nil {
		b.accessTokenResp = body
	}

	fmt.Printf("GetAccessToken :%s\n", body.AccessToken)

	return body
}

func (b *BaiduClient) AuthPOST(url string, urlParam, postData map[string]interface{}, jsonResponseBodyInterface JSONResponseBodyInterface) *gHttpClient.JSONResponse {
	_urlParam := urlParam

	if _urlParam == nil {
		_urlParam = make(map[string]interface{})
	}

	_urlParam["access_token"] = b.GetAccessToken().AccessToken

	return b.POST(url, _urlParam, postData, jsonResponseBodyInterface)
}

func (b *BaiduClient) POST(url string, urlParam, postData map[string]interface{}, jsonResponseBodyInterface JSONResponseBodyInterface) *gHttpClient.JSONResponse {
	resp, err := b.client.Conn.DoJSONResponse("POST", url, urlParam, nil, postData, jsonResponseBodyInterface)

	// fmt.Printf("%v", resp)
	if err != nil {
		jsonResponseBodyInterface.SetError(err)
	}

	jsonResponseBodyInterface.AfterResponseJSON(resp)

	return resp
}
