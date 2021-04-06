package face

import (
	"fmt"
	"testing"

	"gobaiduai/client"
)

func NewFaceClient() *FaceClient {
	c, e := NewFace(client.Config{
		ClientId:     "dd",
		ClientSecret: "vv",
	})

	fmt.Printf("e %v\n", e)
	return c
}

func TestFace_IDMatch(t *testing.T) {
	c := NewFaceClient()
	resp := c.IDMatch("谢梅", "432121197208265941")
	fmt.Printf("e %v\n", resp)
}
