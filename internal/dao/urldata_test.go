package urldata

import (
	"fmt"
	"testing"
)

func TestReadLong(t *testing.T) {
	var (
		in       = "23233"
		expected = "sdsdsd"
	)
	actual := ReadLongUrl(in)
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("ReadLongUrl(%s) = %s; expected %s", in, actual, expected)
	}
}

func TestReadShort(t *testing.T) {
	var (
		in       = "sdsdsd"
		expected = "232333"
	)
	actual := ReadShortUrl(in)
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("ReadLongUrl(%s) = %s; expected %s", in, actual, expected)
	}
}

func TestStoreUrlMap(t *testing.T) {
	var (
		in       = "111111"
		expected = "222222"
	)
	StoreUrl(in, expected)
}
