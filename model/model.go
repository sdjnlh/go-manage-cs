package model

import (
	"jinbao-cs/util"
	"jinbao-cs/web"
	"time"
)

type Base struct {
	Id  string    `xorm:"pk VARCHAR(100)" json:"id" form:"id"`
	Crt time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP" json:"crt"`
	Lut time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP" json:"lut"`
	Dtd bool      `json:"-"`
}

type User struct {
	Base       `xorm:"extends"`
	LoginName  string          `json:"username" form:"username"`
	Password   string          `json:"password" form:"password"`
	Name       string          `json:"name" form:"name"`
	Email      string          `json:"email" form:"email"`
	Mobile     string          `json:"mobile" form:"mobile"`
	WeChat     string          `json:"wechat" form:"wechat"`
	Qq         string          `json:"qq" form:"qq"`
	AvatarId   string          `json:"-"`
	Language   int             `json:"-"`
	OpenId     string          `json:"openId"`
	Roles      string          `json:"roles"`
	RoleNames  web.StringArray `xorm:"-" json:"roleNames"`
	Permission string          `xorm:"-" json:"permission" form:"pobile"`
	Status     int             `json:"status"`
	IdNo       string          `json:"idNo" form:"idNo"` // 身份证号码
	Code       string          `xorm:"-" json:"code"`
	Token      string          `xorm:"-" json:"token" form:"token"`
}
type Role struct {
	Base        `xorm:"extends"`
	RoleName    string   `json:"roleName" form:"roleName"`
	Code        string   `form:"code"`
	Description string   `json:"description" form:"description"`
}
type Page struct {
	P   int    `json:"p" form:"page"`
	Ps  int    `json:"ps" form:"limit"`
	Cnt int64  `json:"cnt"`
	K   string `json:"k" form:"k"`
	Pc  int    `json:"pc"`
	Od  string `json:"od,omitempty"`
}
type Resume struct {
	Base      `xorm:"extends"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Path      string `json:"path"`
	Size      int64  `json:"size"`
	Desc      string `json:"desc"`
	IsPublish bool   `json:"isPublish"`
	Type      string `json:"type"`
}

func (page *Page) GetPage() *Page {
	return page
}

func (page *Page) GetPager(count int64) *Page {
	page.Cnt = count
	if page.P < 1 {
		page.P = 1
	}
	if page.Ps < 1 {
		page.Ps = 10
	}
	page.Pc = int(page.Cnt)/page.Ps + 1
	return page
}

func (page *Page) Skip() int {
	if page.Ps > 0 {
		return (page.P - 1) * page.Ps
	}

	return (page.P - 1) * 10
}

func (page *Page) Limit() int {
	if page.Ps > 0 {
		return page.Ps
	}

	return 10
}

type NameAndDesc struct {
	Name        string `xorm:"name" json:"name" form:"name"`                      // 名称
	Description string `xorm:"description" json:"description" form:"description"` // 详细描述
}
type File struct {
	Base         `xorm:"extends"`
	NameAndDesc  `xorm:"extends"`
	PrefixUri    string `json:"prefixUri"`    // 网络地址
	RelativePath string `json:"relativePath"` // 绝对路径
	Kind         string `json:"kind"`         // 类型
	OriginName   string `json:"originName"`   // 原始文件名
	Suffix       string `json:"suffix"`       // 后缀
	UniqueName   string `json:"uniqueName"`
}

func (b *Base) BeforeInsert() {
	b.Id = util.GetRandomString(32)
	now := time.Now()
	b.Crt = now
	b.Lut = now
}

type Captcha struct {
	CaptchaId  string `json:"captchaId"`
	CaptchaImg string `json:"captchaImg"`
}
