package client

import (
	"fmt"
	"testing"
)

func NewC() *BaiduClient {
	c, e := NewHttpClient(Config{
		ClientId:     "vv",
		ClientSecret: "dd",
	})

	fmt.Printf("e %v\n", e)
	return c
}




func TestHttpClient_GetAccessToken(t *testing.T) {
	client := NewC()
	body := client.GetAccessToken()
	fmt.Printf("body:%v \n", body)
	fmt.Printf("AccessToken:%s \n", body.AccessToken)
	fmt.Printf("Error:%s \n", body.Error)

}
