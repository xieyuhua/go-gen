
package api
import (
	"ginx/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tpadminoperationlogInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/tpadminoperationlog/create
	groupgo.POST("/tpadminoperationlog/create", TpAdminOperationLogCreate)
	groupgo.POST("/tpadminoperationlog/delete", TpAdminOperationLogDelete)
	groupgo.POST("/tpadminoperationlog/update", TpAdminOperationLogUpdate)
	groupgo.POST("/tpadminoperationlog/select", TpAdminOperationLogSelect)
}

// TpAdminOperationLogCreate ...
func TpAdminOperationLogCreate(c *gin.Context) {

	var param service.TpAdminOperationLogCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpAdminOperationLogService{}
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

// TpAdminOperationLogDelete ...
func TpAdminOperationLogDelete(c *gin.Context) {

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
	ser := &service.TpAdminOperationLogService{}
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

// TpAdminOperationLogUpdate ...
func TpAdminOperationLogUpdate(c *gin.Context) {

	var param service.TpAdminOperationLogUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpAdminOperationLogService{}
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

// TpAdminOperationLogSelect ...
func TpAdminOperationLogSelect(c *gin.Context) {

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
	ser := &service.TpAdminOperationLogService{}
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

