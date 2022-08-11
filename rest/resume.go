package rest

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path/filepath"
)

type resumeApi int

var ResumeApi resumeApi

const (
	readModeView = iota
	readModeDownload
)

func (resumeApi) Preview(c *gin.Context) {
	readResume(c, readModeView)
}
func (resumeApi) Download(c *gin.Context) {
	readResume(c, readModeDownload)
}
func readResume(c *gin.Context, mode int) {
	root, _ := filepath.Abs(".")
	fileName := "my-resume.pdf"
	dest := filepath.Join(root, fileName)
	fileByte, err := ioutil.ReadFile(dest)
	if err != nil {
		c.String(500, "读取简历文件失败")
		c.Abort()
		return
	}
	if mode == readModeView {
		c.Header("Content-Type", "application/pdf")
	} else {
		fileName = "李皓.pdf"
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Transfer-Encoding", "binary")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	}
	c.Writer.Write(fileByte)
}

func (resumeApi) Register(c *gin.RouterGroup) {
	c.GET("/v1/resumes/view", ResumeApi.Preview)
	c.GET("/v1/resumes/download", ResumeApi.Download)
}
