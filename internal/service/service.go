package service

import (
	urldata "ShortUrl/internal/dao"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SubString(str string, begin, length int) (substr string) {
	rs := []rune(str)
	lth := len(rs)

	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}

func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

func ShortProcess(longstring string) string {
	var key = "leoaylearngo"
	var base32 = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	md5s := md5.New()
	md5s.Write([]byte(longstring + key))
	hexstr := hex.EncodeToString(md5s.Sum(nil))

	hexLen := len(hexstr)
	subHexLen := hexLen / 8
	var output []string
	for i := 0; i < subHexLen; i++ {
		subHex := SubString(hexstr, i*8, 8)
		idx := 0x3FFFFFFF & Hex2Dec(subHex)
		out := ""
		for j := 0; j < 6; j++ {
			val := 0x0000003D & idx
			out = out + string(base32[val])
			idx = idx >> 5
		}
		output = append(output, out)
	}
	return output[0]
}

type LongUrlType struct {
	LongUrl string `json:"longurl"`
}

func Long2Short(ctx *gin.Context) {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var LongUrl LongUrlType
	err := json.Unmarshal(data, &LongUrl)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	longurl := LongUrl.LongUrl
	shorturl := ShortProcess(longurl)
	fullshorturl := "https://leoay.com/?st=" + shorturl

	urldata.StoreUrl(shorturl, longurl)

	ctx.JSON(200, gin.H{
		"longurl":  longurl,
		"shorturl": fullshorturl,
	})
}
