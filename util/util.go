package util

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"math/rand"
	"net/url"
	"sort"
	"strings"
	"time"
)

// GeneralNumber 生成指定长度的首位不为1的随机数字
func GeneralNumber(width int, seeds ...int64) string {
	seed := time.Now().UnixNano()
	if len(seeds) > 0 {
		seed = seed + seeds[0]
	}
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(seed)

	var sb strings.Builder
	for i := 0; i < width; i++ {
		num := numeric[rand.Intn(r)]
		if i == 0 && num == 0 {
			num = 1
		}
		fmt.Fprintf(&sb, "%d", num)
	}
	return sb.String()
}

//func Utf8ToGbk(s []byte) ([]byte, error) {
//	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
//	d, e := ioutil.ReadAll(reader)
//	if e != nil {
//		return nil, e
//	}
//	return d, nil
//}

func UrlEncode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}
func MapEncode(v map[string]interface{}) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	for key, val := range v {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(key)
		buf.WriteByte('=')
		buf.WriteString(val.(string))
	}
	return buf.String()
}

func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func XmlToJson(p string, v interface{}) (string, error) {
	xmlParam, err := url.PathUnescape(p)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader([]byte(xmlParam))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(v)
	if err != nil {
		//logger.Error("unmarshal xml", err)
		return "", err
	}

	result, err := json.Marshal(v)
	if err != nil {
		//logger.Error("unmarshal xml", err)
		return "", err
	}
	return string(result), nil
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GenValidateCode 生成指定长度的首位不为1的随机数字
func GenValidateCode(width int, seeds ...int64) string {
	seed := time.Now().UnixNano()
	if len(seeds) > 0 {
		seed = seed + seeds[0]
	}
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(seed)

	var sb strings.Builder
	for i := 0; i < width; i++ {
		num := numeric[rand.Intn(r)]
		if i == 0 && num == 0 {
			num = 1
		}
		fmt.Fprintf(&sb, "%d", num)
	}
	return sb.String()
}

type CompareFunc func(interface{}, interface{}) int

func IndexOf(a []interface{}, e interface{}, cmp CompareFunc) int {
	n := len(a)
	var i int = 0
	for ; i < n; i++ {
		if cmp(e, a[i]) == 0 {
			return i
		}
	}
	return -1
}

func IsContainArr(whiteList []interface{}, requestUrl string) (isRelease bool) {
	for _,white := range whiteList {
		if requestUrl == white {
			isRelease = true
		}else {
			continue
		}
	}
	return isRelease
}