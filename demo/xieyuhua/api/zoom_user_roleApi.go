
package api
import (
	"xieyuhua/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func zoomuserroleInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/zoomuserrole/create
	groupgo.POST("/zoomuserrole/create", ZoomUserRoleCreate)
	groupgo.POST("/zoomuserrole/delete", ZoomUserRoleDelete)
	groupgo.POST("/zoomuserrole/update", ZoomUserRoleUpdate)
	groupgo.POST("/zoomuserrole/select", ZoomUserRoleSelect)
	groupgo.GET("/zoomuserrole/all", ZoomUserRoleAll)
}

// ZoomUserRoleCreate ...
func ZoomUserRoleCreate(c *gin.Context) {

	var param service.ZoomUserRoleCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomUserRoleService{}
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

// ZoomUserRoleDelete ...
func ZoomUserRoleDelete(c *gin.Context) {

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
	ser := &service.ZoomUserRoleService{}
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

// ZoomUserRoleUpdate ...
func ZoomUserRoleUpdate(c *gin.Context) {

	var param service.ZoomUserRoleUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomUserRoleService{}
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

// ZoomUserRoleAll ...
func ZoomUserRoleAll(c *gin.Context) {

	ser := &service.ZoomUserRoleService{}
	back, err := ser.All()
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

// ZoomUserRoleSelect ...
func ZoomUserRoleSelect(c *gin.Context) {

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
	ser := &service.ZoomUserRoleService{}
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

