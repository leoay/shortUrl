package service

import (
	"fmt"
	"testing"
)

func TestShortProcess(t *testing.T) {
	var (
		in       = "https://blog.csdn.net/weixin_39172380/article/details/95059107"
		expected = "7VJFvu"
	)
	actual := ShortProcess(in)
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("ReadLongUrl(%s) = %s; expected %s", in, actual, expected)
	}
}
