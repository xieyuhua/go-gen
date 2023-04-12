
package api
import (
	"ddada/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func zoomdeploybuildInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/zoomdeploybuild/create
	groupgo.POST("/zoomdeploybuild/create", ZoomDeployBuildCreate)
	groupgo.POST("/zoomdeploybuild/delete", ZoomDeployBuildDelete)
	groupgo.POST("/zoomdeploybuild/update", ZoomDeployBuildUpdate)
	groupgo.POST("/zoomdeploybuild/select", ZoomDeployBuildSelect)
	groupgo.GET("/zoomdeploybuild/all", ZoomDeployBuildAll)
}

// ZoomDeployBuildCreate ...
func ZoomDeployBuildCreate(c *gin.Context) {

	var param service.ZoomDeployBuildCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomDeployBuildService{}
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

// ZoomDeployBuildDelete ...
func ZoomDeployBuildDelete(c *gin.Context) {

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
	ser := &service.ZoomDeployBuildService{}
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

// ZoomDeployBuildUpdate ...
func ZoomDeployBuildUpdate(c *gin.Context) {

	var param service.ZoomDeployBuildUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomDeployBuildService{}
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

// ZoomDeployBuildSelect ...
func ZoomDeployBuildSelect(c *gin.Context) {

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
	ser := &service.ZoomDeployBuildService{}
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

