package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"jinbao-cs/cs"
	"jinbao-cs/errors"
	"jinbao-cs/log"
	"jinbao-cs/middleware"
	"jinbao-cs/model"
	"jinbao-cs/service"
	"jinbao-cs/web"
	"os"
)

type userAPI struct {
	*web.RestHandler
}

func NewUserAPI() *userAPI {
	return &userAPI{
		web.DefaultRestHandler,
	}
}

func (api *userAPI) login(c *gin.Context) {
	form := &model.User{}
	err := api.Bind(c, form)
	if err != nil {
		log.Logger.Error("用户登录绑定参数出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	result := web.NewFilterResult(&model.User{})
	err = service.User.Login(form, result)
	if err != nil {
		log.Logger.Error("用户登录失败", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	var sessionID = cs.SessionMgr.StartSession(c.Writer, c.Request)

	//设置变量值
	cs.SessionMgr.SetSessionVal(sessionID, "UserInfo", result.Data)
	api.ResultWithError(c, result, err)
}
func (api *userAPI) create(c *gin.Context) {
	user := &model.User{}
	err := api.Bind(c, user)
	if err != nil {
		log.Logger.Error("创建用户绑定参数出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	result := web.NewFilterResult(&model.User{})
	err = service.User.Create(user, result)
	if err != nil {
		log.Logger.Error("新增用户出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	api.ResultWithError(c, result, err)
}
func (api *userAPI) update(c *gin.Context) {
	user := &model.User{}
	err := api.Bind(c, user)
	if err != nil {
		log.Logger.Error("修改用户绑定参数出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	result := web.NewFilterResult(&model.User{})
	err = service.User.Update(user, result)
	if err != nil {
		log.Logger.Error("修改用户出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	api.ResultWithError(c, result, err)
}
func (api *userAPI) list(c *gin.Context) {
	page := &model.Page{}
	err := api.Bind(c, page)
	if err != nil {
		log.Logger.Error("绑定用户列表分页参数出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	result := web.NewFilterResult(&[]model.User{})
	users := &[]model.User{}
	err = service.User.List(page, users, result)
	if err != nil {
		log.Logger.Error("获取用户列表数据出错", zap.Error(err))
		return
	}
	api.ResultWithError(c, result, err)
}

func (api *userAPI) delete(c *gin.Context) {
	user := &model.User{}
	err := api.Bind(c, user)
	if err != nil {
		log.Logger.Error("删除用户绑定参数出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	result := web.NewFilterResult(&model.User{})
	err = service.User.Delete(user, result)
	if err != nil {
		log.Logger.Error("删除用户出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	api.ResultWithError(c, result, err)
}
func (api *userAPI) BatchDelete(c *gin.Context) {
	user := &model.UserDTO{}
	err := api.Bind(c, user)
	if err != nil {
		log.Logger.Error("删除用户绑定参数出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	result := web.NewFilterResult(&model.User{})
	err = service.User.BatchDelete(user, result)
	if err != nil {
		log.Logger.Error("批量删除用户出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	api.ResultWithError(c, result, err)
}
func (api *userAPI) get(c *gin.Context) {
	strid := c.Param("id")
	user := &model.User{}
	user.Id = strid
	if err := service.User.Get(user); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, user)
	}
}
func (api *userAPI) Me(c *gin.Context) {
	form := &model.UserInfoDTO{}
	err := api.Bind(c, form)
	if err != nil {
		log.Logger.Error("用户登录绑定参数出错", zap.Error(err))
		api.FailWithError(c, errors.InvalidParams())
		return
	}
	middleware.JWY(*form, c)
	claims, _ := middleware.ParseToken(form.Token)
	//if err != nil {
	//	log.Logger.Error("token不合法或token已过期",zap.Error(err))
	//	api.FailWithError(c, errors.InvalidParams())
	//	return
	//}
	//获取用户信息
	result := web.NewFilterResult(&model.User{})
	if claims == nil {
		result.Ok = true
		result.Msg = "Token无效或Token已过期"
		api.ResultWithError(c, result, nil)
	}
	err = service.User.GetUserInfoById(claims.Id, result)
	if err != nil {
		log.Logger.Error("根据ID获取用户信息出错", zap.Error(err))
		return
	}
	api.ResultWithError(c, result, nil)
}

func (api *userAPI) download(c *gin.Context) {
	resume := &model.Resume{}
	if err := c.Bind(resume); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
	if resume.Path != "" {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename=text")
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(".." + resume.Path)
	} else {
		c.Abort()
		c.JSON(500, "内部服务器错误")
		return
	}
}
func (api *userAPI) uploadFile(c *gin.Context) {
	resume := model.Resume{}

	form, _ := c.MultipartForm()
	version := c.PostForm("version")
	files := form.File["file[]"]
	resume.Version = version

	// contract.Files=Files
	pathstr := "/resume/" + version
	path := ".." + pathstr
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			c.String(400, "创建文件失败")
			c.Abort()
			return
		} //0777也可以os.ModePerm
		if err := os.Chmod(path, os.ModePerm); err != nil {
			c.String(400, "chmod失败")
			c.Abort()
			return
		}
	}
	fileName := make([]string, 0)
	filePath := make([]string, 0)
	for _, file := range files {
		//文件存储
		str := path + "/" + file.Filename
		strpath := pathstr + "/" + file.Filename
		fWrite, err := os.Create(str)
		if err != nil {
			c.String(500, "服务器错误")
			c.Abort()
			return
		}
		resume.Size = file.Size
		f, _ := file.Open()
		if _, err := io.Copy(fWrite, f); err != nil {

		}
		defer fWrite.Close()
		filePath = append(filePath, strpath)
		fileName = append(fileName, file.Filename)
		resume.Name = file.Filename
		resume.Path = strpath
	}
	c.JSON(200, resume)
}

func (api *userAPI) getRoles(c *gin.Context) {
	result := web.NewFilterResult(&[]model.Role{})
	err := service.User.GetRole(result)
	if err != nil {
		log.Logger.Error("获取角色列表数据出错", zap.Error(err))
		return
	}
	api.ResultWithError(c, result, err)
}

func (api *userAPI) Register(router gin.IRouter) {
	v1 := router.Group("/user")
	v1.POST("/info/get", api.Me)
	v1.POST("/login", api.login)
	v1.POST("/create", api.create)
	v1.PUT("/update", api.update)
	v1.POST("/list/get", api.list)
	v1.GET("/v1/web/:id", api.get)
	v1.GET("/role/list", api.getRoles)
	v1.DELETE("/delete", api.delete)
	v1.DELETE("/batchDelete", api.BatchDelete)
	v1.POST("/v1/resume/upload", api.uploadFile)
	v1.POST("/v1/resume/download", api.download)
}
