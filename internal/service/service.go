package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	urldata "ShortUrl/internal/dao"

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
	fmt.Println(hexLen)
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
	fmt.Println(string(data))
	var LongUrl LongUrlType
	err := json.Unmarshal(data, &LongUrl)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	longurl := LongUrl.LongUrl
	shorturl := "https://leoay.com/?st=" + ShortProcess(longurl)

	urldata.StoreUrl(shorturl, longurl)

	ctx.JSON(200, gin.H{
		"longurl":  longurl,
		"shorturl": shorturl,
	})
}
