
package api
import (
	"ddada/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func zoomprojectmemberInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/zoomprojectmember/create
	groupgo.POST("/zoomprojectmember/create", ZoomProjectMemberCreate)
	groupgo.POST("/zoomprojectmember/delete", ZoomProjectMemberDelete)
	groupgo.POST("/zoomprojectmember/update", ZoomProjectMemberUpdate)
	groupgo.POST("/zoomprojectmember/select", ZoomProjectMemberSelect)
	groupgo.GET("/zoomprojectmember/all", ZoomProjectMemberAll)
}

// ZoomProjectMemberCreate ...
func ZoomProjectMemberCreate(c *gin.Context) {

	var param service.ZoomProjectMemberCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomProjectMemberService{}
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

// ZoomProjectMemberDelete ...
func ZoomProjectMemberDelete(c *gin.Context) {

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
	ser := &service.ZoomProjectMemberService{}
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

// ZoomProjectMemberUpdate ...
func ZoomProjectMemberUpdate(c *gin.Context) {

	var param service.ZoomProjectMemberUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.ZoomProjectMemberService{}
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

// ZoomProjectMemberSelect ...
func ZoomProjectMemberSelect(c *gin.Context) {

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
	ser := &service.ZoomProjectMemberService{}
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

