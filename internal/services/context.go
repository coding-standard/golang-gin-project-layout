package services

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	errKey     = "err"
	msgKey     = "msg"
	dataKey    = "data"
	flashKey   = "flash"
	userKey    = "user"
	unreadKey  = "unread"
	versionKey = "version"
)

func Context(ctx *gin.Context) *BaseContext {
	stx := &BaseContext{
		Ctx:     ctx,
		session: sessions.Default(ctx),
		path:    "/",
	}
	return stx
}

type BaseContext struct {
	Ctx     *gin.Context
	session sessions.Session
	path    string
}

// Redirect 处理跳转
func (c *BaseContext) Redirect() {
	c.Ctx.Redirect(http.StatusFound, c.path)
}

func (c *BaseContext) clear() {
	c.session.Delete(errKey)
	c.session.Delete(msgKey)
	c.session.Delete(flashKey)
	_ = c.session.Save()
}

// Back 返回上一页
func (c *BaseContext) Back() *BaseContext {
	c.path = c.Ctx.Request.RequestURI
	return c
}

// To 设置跳转路径
func (c *BaseContext) To(to string) *BaseContext {
	c.path = to
	return c
}

// WithData 闪存消息
func (c *BaseContext) WithData(data interface{}) *BaseContext {
	r, _ := json.Marshal(data)
	c.session.Set(flashKey, string(r))
	_ = c.session.Save()
	return c
}

// WithError 错误信息跳转
func (c *BaseContext) WithError(err interface{}) *BaseContext {
	errStr := ""
	switch v := err.(type) {
	case error:
		errStr = v.Error()
	case string:
		errStr = v
	}
	c.session.Set(errKey, errStr)
	_ = c.session.Save()
	return c
}

// View 模版返回
func (c *BaseContext) View(tpl string, data interface{}) {
	obj := gin.H{
		dataKey: data,
	}
	c.clear()
	c.Ctx.HTML(http.StatusOK, tpl, obj)
}

// Json 通用 JSON 响应
func (c *BaseContext) Json(data interface{}) {
	c.Ctx.JSON(http.StatusOK, data)
}

// MDFileJson markdown 上传图片响应
func (c *BaseContext) MDFileJson(ok int, msg, url string) {
	c.Ctx.JSON(http.StatusOK, gin.H{"success": ok, "message": msg, "url": url})
}
