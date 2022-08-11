package service

import (
	"go.uber.org/zap"
	"jinbao-cs/cs"
	"jinbao-cs/dict"
	"jinbao-cs/log"
	"jinbao-cs/middleware"
	"jinbao-cs/model"
	"jinbao-cs/password"
	"jinbao-cs/web"
	"strings"
	"time"
)

type userService struct {
}

var User userService

func (userService) Login(form *model.User, result *web.FilterResult) (err error) {
	var selectUser model.User
	_, err = cs.Sql.Table("op_user").Where(" dtd = false and name = ?", form.LoginName).
		Cols("id", "login_name", "password", "name").Get(&selectUser)
	if err != nil {
		log.Logger.Error("用户登录接口根据用户名获取信息出错", zap.Error(err))
		return err
	}
	if selectUser.Id != "" {
		if password.Validate(form.Password, selectUser.Password) {
			selectUser.Token, err = middleware.GenerateToken(selectUser.Id, form.LoginName, form.Password)
			if err != nil {
				log.Logger.Error("获取用户登录token出错", zap.Error(err))
				return err
			}
			result.Ok = true
		} else {
			result.Ok = false
		}
	}
	result.Ok = true
	result.Data = selectUser
	return err
}
func (userService) Create(form *model.User, result *web.FilterResult) (err error) {
	form.BeforeInsert()
	form.Password, _ = password.Encrypt(dict.DefaultPwd)
	_, err = cs.Sql.Table("op_user").Insert(form)
	if err != nil {
		log.Logger.Error("插入用户数据出错", zap.Error(err))
		return err
	}
	result.Ok = true
	result.Data = form
	return err
}
func (userService) Update(form *model.User, result *web.FilterResult) error {
	form.Lut = time.Now()
	if _, err := cs.Sql.Table("op_user").ID(form.Id).Update(form); err != nil {
		return err
	}
	result.Ok = true
	result.Data = form
	return nil
}
func (userService) List(page *model.Page, users *[]model.User, result *web.FilterResult) error {
	condition := " dtd = false "
	if page.K != "" {
		condition += " and (login_name ilike '%" + page.K + "%' or name ilike '%" + page.K + "%') "
	}
	cnt, err := cs.Sql.Table("op_user").Where(condition).Limit(page.Limit(), page.Skip()).FindAndCount(users)
	if err != nil {
		log.Logger.Error("获取用户列表信息出错", zap.Error(err))
		return err
	}
	var data []model.User
	for _, user := range *users {
		roleArr := strings.Split(user.Roles, ",")
		if len(roleArr) > 0 {
			var roleIds web.StringArray
			for _, roleId := range roleArr {
				roleIds = append(roleIds, roleId)
			}
			_, err = cs.Sql.SQL("select string_agg(code,' | ' order by code) as permission from op_role where id in " + roleIds.ToIn()).Get(&user.Permission)
			if err != nil {
				log.Logger.Error("获取用户角色名称出错", zap.Error(err))
				return err
			}
			data = append(data, user)
		}
	}
	page.Cnt = cnt
	result.Ok = true
	result.Data = data
	return err
}
func (userService) Delete(user *model.User, result *web.FilterResult) error {
	user.Dtd = true
	user.Lut = time.Now()
	if _, err := cs.Sql.Table("op_user").ID(user.Id).Cols("dtd", "lut").Update(user); err != nil {
		return err
	}
	result.Ok = true
	result.Data = user
	return nil
}
func (userService) BatchDelete(user *model.UserDTO, result *web.FilterResult) error {
	sql := "update op_user set dtd = true,lut=now() where dtd = false and id in " + user.Ids.ToIn()
	if _, err := cs.Sql.Exec(sql); err != nil {
		return err
	}
	result.Ok = true
	result.Data = user
	return nil
}
func (userService) Get(user *model.User) error {
	if _, err := cs.Sql.Table("op_user").ID(user.Id).Get(user); err != nil {
		return err
	}
	return nil
}

func (userService) GetUserInfoById(id string, result *web.FilterResult) (err error) {
	var user model.User
	_, err = cs.Sql.Table("op_user").Where("id = ?", id).Get(&user)
	if err != nil {
		log.Logger.Error("根据ID获取用户信息出错", zap.Error(err))
		return err
	}
	roleArr := strings.Split(user.Roles, ",")
	var roles web.StringArray
	if len(roleArr) > 0 {
		var roleIds web.StringArray
		for _, roleId := range roleArr {
			roleIds = append(roleIds, roleId)
		}
		err = cs.Sql.Table("role").Where("id in " + roleIds.ToIn()).Cols("code").Find(&roles)
		if err != nil {
			log.Logger.Error("获取用户角色权限出错", zap.Error(err))
			return err
		}
		user.RoleNames = roles
	}
	result.Ok = true
	result.Data = user
	return err
}

func (userService) GetRole(result *web.FilterResult) (err error) {
	roles := &[]model.Role{}
	err = cs.Sql.Table("op_role").Where(" dtd = false").Find(roles)
	if err != nil {
		log.Logger.Error("获取角色数据出错", zap.Error(err))
		return err
	}
	result.Ok = true
	result.Data = roles
	return err
}
