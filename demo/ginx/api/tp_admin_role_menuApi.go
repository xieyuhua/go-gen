
package api
import (
	"ginx/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tpadminrolemenuInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/tpadminrolemenu/create
	groupgo.POST("/tpadminrolemenu/create", TpAdminRoleMenuCreate)
	groupgo.POST("/tpadminrolemenu/delete", TpAdminRoleMenuDelete)
	groupgo.POST("/tpadminrolemenu/update", TpAdminRoleMenuUpdate)
	groupgo.POST("/tpadminrolemenu/select", TpAdminRoleMenuSelect)
}

// TpAdminRoleMenuCreate ...
func TpAdminRoleMenuCreate(c *gin.Context) {

	var param service.TpAdminRoleMenuCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpAdminRoleMenuService{}
	back, err := ser.Create(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	reply.Data = back
	c.JSON(http.StatusOK, reply)
	return
}

// TpAdminRoleMenuDelete ...
func TpAdminRoleMenuDelete(c *gin.Context) {

	var param struct {
		ID int "json:\"id\" form:\"id\""
	}
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpAdminRoleMenuService{}
	err = ser.Delete(param.ID)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	c.JSON(http.StatusOK, reply)
	return
}

// TpAdminRoleMenuUpdate ...
func TpAdminRoleMenuUpdate(c *gin.Context) {

	var param service.TpAdminRoleMenuUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpAdminRoleMenuService{}
	back, err := ser.Update(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	reply.Data = back
	c.JSON(http.StatusOK, reply)
	return
}

// TpAdminRoleMenuSelect ...
func TpAdminRoleMenuSelect(c *gin.Context) {

	var param struct {
		ID int "json:\"id\" form:\"id\""
	}
	//解析参数
	err := c.ShouldBindJSON(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpAdminRoleMenuService{}
	back, err := ser.Select(param.ID)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	reply := NewReplyOk()
	reply.Data = back
	c.JSON(http.StatusOK, reply)
	return
}

