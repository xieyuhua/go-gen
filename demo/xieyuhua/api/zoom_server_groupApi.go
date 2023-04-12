
package api
import (
	"xieyuhua/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func zoomservergroupInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/zoomservergroup/create
	groupgo.POST("/zoomservergroup/create", ZoomServerGroupCreate)
	groupgo.POST("/zoomservergroup/delete", ZoomServerGroupDelete)
	groupgo.POST("/zoomservergroup/update", ZoomServerGroupUpdate)
	groupgo.POST("/zoomservergroup/select", ZoomServerGroupSelect)
	groupgo.GET("/zoomservergroup/all", ZoomServerGroupAll)
}

// ZoomServerGroupCreate ...
func ZoomServerGroupCreate(c *gin.Context) {

	var param service.ZoomServerGroupCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomServerGroupService{}
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

// ZoomServerGroupDelete ...
func ZoomServerGroupDelete(c *gin.Context) {

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
	ser := &service.ZoomServerGroupService{}
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

// ZoomServerGroupUpdate ...
func ZoomServerGroupUpdate(c *gin.Context) {

	var param service.ZoomServerGroupUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomServerGroupService{}
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

// ZoomServerGroupAll ...
func ZoomServerGroupAll(c *gin.Context) {

	ser := &service.ZoomServerGroupService{}
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

// ZoomServerGroupSelect ...
func ZoomServerGroupSelect(c *gin.Context) {

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
	ser := &service.ZoomServerGroupService{}
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

