
package api
import (
	"yuhua/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func zoomdeployapplyInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/zoomdeployapply/create
	groupgo.POST("/zoomdeployapply/create", ZoomDeployApplyCreate)
	groupgo.POST("/zoomdeployapply/delete", ZoomDeployApplyDelete)
	groupgo.POST("/zoomdeployapply/update", ZoomDeployApplyUpdate)
	groupgo.POST("/zoomdeployapply/select", ZoomDeployApplySelect)
}

// ZoomDeployApplyCreate ...
func ZoomDeployApplyCreate(c *gin.Context) {

	var param service.ZoomDeployApplyCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomDeployApplyService{}
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

// ZoomDeployApplyDelete ...
func ZoomDeployApplyDelete(c *gin.Context) {

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
	ser := &service.ZoomDeployApplyService{}
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

// ZoomDeployApplyUpdate ...
func ZoomDeployApplyUpdate(c *gin.Context) {

	var param service.ZoomDeployApplyUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomDeployApplyService{}
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

// ZoomDeployApplySelect ...
func ZoomDeployApplySelect(c *gin.Context) {

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
	ser := &service.ZoomDeployApplyService{}
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

