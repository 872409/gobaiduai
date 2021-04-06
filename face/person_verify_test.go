package face

import (
	"fmt"
	"testing"
)

func TestFaceClient_PersonVerify(t *testing.T) {
	c := NewFaceClient()
	// resp:=c.IDMatch("谢梅英","352121197408264941")
	base64 := "/9j/4AAQSkZJRgABAQAAkACQAAD//+/++++/////zJgAAB5IAAP2R///7ov////Z"
	resp := c.PersonVerifyDefault(base64, ImageTypeBASE64, "谢梅", "234234234234")
	fmt.Printf("e %v %v\n", resp.IsMatch(), resp.Error)
}
