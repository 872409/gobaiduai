package face

import (
	"fmt"
	"testing"

	"github.com/872409/gobaiduai/client"
)

func NewFaceClient() *FaceClient {
	c, e := NewFace(client.Config{

	})

	fmt.Printf("e %v\n", e)
	return c
}

func TestFace_IDMatch(t *testing.T) {
	c := NewFaceClient()
	resp := c.IDMatch("谢梅", "432121197208265941")
	fmt.Printf("e %v\n", resp)
}
