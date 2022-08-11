package password

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestAesDecryptECB(t *testing.T) {
	//加密
	str := []byte("@dataon@cnzdf@2018")
	key := []byte("dataon$123")
	res, err := AesEncryptECB(str, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hex.EncodeToString(res))
	fmt.Println(strings.ToUpper(hex.EncodeToString(res)))

	//解密
	jmh := "B95BBFDF658388E329D0F5B2CD222ED2"
	content, err := hex.DecodeString(jmh)
	if err != nil {
		t.Fatal(err)
	}
	de, err := AesDecryptECB(content, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(de))
	fmt.Println(string(de))
}

func TestEncrypt(t *testing.T) {
	val := "123456"
	result, err := Encrypt(val)
	if err != nil {
		fmt.Println("密码加密出错")
	}
	fmt.Println(result)
}
