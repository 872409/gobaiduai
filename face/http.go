package face

import (
	"errors"
	"fmt"
	"time"

	"github.com/872409/gobaiduai/client"
)

type AuthResponseJSON struct {
	Succeed    bool      `json:"-"`
	Error      error     `json:"-"`
	ErrorCode  int       `json:"error_code,int"`
	ErrorMsg   string    `json:"error_msg"`
	ResponseAt time.Time `json:"-"`
}

var _ client.JSONResponseBodyInterface = &AuthResponseJSON{}

func (j *AuthResponseJSON) SetError(err error) {
	j.Error = err
}

func (j *AuthResponseJSON) AfterResponse() {
	j.ResponseAt = time.Now()

	if j.Error != nil {
		j.ErrorCode = -500
		j.ErrorMsg = j.Error.Error()
	} else if j.ErrorCode != 0 {
		j.Error = errors.New(fmt.Sprintf("[%d] %s", j.ErrorCode, j.ErrorMsg))
	}

	j.Succeed = j.ErrorCode == 0
}

func NewFace(config client.Config) (*FaceClient, error) {
	client, err := client.NewHttpClient(config)
	return &FaceClient{
		httpClient: client,
	}, err
}

type FaceClient struct {
	httpClient *client.BaiduClient
}
