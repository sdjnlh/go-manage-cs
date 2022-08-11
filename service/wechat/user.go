package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"jinbao-cs/cs"
	"jinbao-cs/dict"
	"jinbao-cs/model"
	"net/http"
	"strconv"
)

var User user

type user int

func (user) Get(form *model.User) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.Id).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (user) List(form *model.User, page *model.Page, list *[]model.User) error {
	// 分页查询
	if cnt, err := cs.Sql.Limit(page.Limit(), page.Skip()).Desc("status", "created_at").FindAndCount(list, form); err != nil {
		return err
	} else {
		page = page.GetPager(cnt)
	}

	return nil
}

// Update 更新新的纪录
func (user) Update(form *model.User) error {
	if _, err := cs.Sql.ID(form.Id).Update(form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t user) Delete(form *model.User) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.Id).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (t user) Save(form *model.User) error {
	if form.Id == "" {
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}

func (t user) Login(form *model.User) error {
	// Login 登录或者新建
	resp, err := http.Get(fmt.Sprintf(dict.WxLogin, dict.LibrarayId, dict.LibrarySecret, form.Code))
	if err != nil {

		return err
	}
	response, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := errors.New("登录失败当前状态吗" + strconv.Itoa(resp.StatusCode))

		return err
	}

	if err != nil {

		return err
	}
	if err = json.Unmarshal(response, form); err != nil {

		return err
	}
	//	t.Decrypt(form)

	user := &model.User{}
	if _, err := cs.Sql.Where("open_id = ?", form.OpenId).Get(user); err != nil {
		return err
	}

	if user.Id == "" {
		// 用户不存在  新建用户再返回
		// 新建时，给usrname 一个初始值
		form.Base.BeforeInsert()

		if err := t.Save(form); err != nil {

			return err
		}
	} else {
		*form = *user
	}
	return nil
}

//func (t user) Decrypt(form *model.User) error {
//	aesk, err := base64.StdEncoding.DecodeString(form.SessionKey)
//	if err != nil {
//		return err
//	}
//	cipherText, err := base64.StdEncoding.DecodeString(form.EncryptedData)
//	if err != nil {
//		return err
//	}
//	ivBytes, err := base64.StdEncoding.DecodeString(form.Iv)
//	if err != nil {
//		return err
//	}
//	block, err := aes.NewCipher(aesk)
//	if err != nil {
//		return err
//	}
//	mode := cipher.NewCBCDecrypter(block, ivBytes)
//	mode.CryptBlocks(cipherText, cipherText)
//	cipherText = pkcs7Unpad(cipherText, block.BlockSize())
//	if err != nil {
//		return err
//	}
//	err = json.Unmarshal(cipherText, &form)
//	if err != nil {
//		return err
//	}
//	return nil
//}
func pkcs7Unpad(data []byte, blockSize int) []byte {

	if blockSize <= 0 {
		return nil
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil
		}
	}
	return data[:len(data)-n]
}
