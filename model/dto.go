package model

import "jinbao-cs/web"

type bookVo struct {
	Type   string `json:"type,string form:type,string"`
	Author string `json:"author,string from:author,string"`
}

type Updatebook struct {
	BookId    int64  `json:"bookId,string" form:"bookId,string"`
	UserId    int64  `json:"userId,string" form:"userId,string"`
	LibraryId int64  `json:"libraryId,string" form:"libraryId,string"`
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
}
type ApplysVo struct {
	Base      `xorm:"extends"`
	Num       string `json:"num"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	School    string `json:"school"`
	Phone     string `json:"phone"`
	LibraryId int64  `json:"libraryId,string"`
	UserId    int64  `json:"userId,string"`
	Site      string `json:"site"`
	Time      string `json:"time"`
	//Library   Library `xorm:"-" json:"library"`
}
type UserDTO struct {
	Ids web.StringArray `form:"ids"`
}

type UserInfoDTO struct {
	Token string `json:"token" form:"token"`
}
