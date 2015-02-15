package gocmdo

import (
	"fmt"
)

func hello(args []string) string {
  image := "http://www.yosemitepark.com/Images/Header-Plan-1.1_img1.jpg"
  return fmt.Sprintf("Hello yoruself! %s", image)
}

func init() {
	registerHandler("hello", hello)
}
