
package api
import (
	"ginx/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tpshopordergoodInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/tpshopordergood/create
	groupgo.POST("/tpshopordergood/create", TpShopOrderGoodCreate)
	groupgo.POST("/tpshopordergood/delete", TpShopOrderGoodDelete)
	groupgo.POST("/tpshopordergood/update", TpShopOrderGoodUpdate)
	groupgo.POST("/tpshopordergood/select", TpShopOrderGoodSelect)
}

// TpShopOrderGoodCreate ...
func TpShopOrderGoodCreate(c *gin.Context) {

	var param service.TpShopOrderGoodCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpShopOrderGoodService{}
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

// TpShopOrderGoodDelete ...
func TpShopOrderGoodDelete(c *gin.Context) {

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
	ser := &service.TpShopOrderGoodService{}
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

// TpShopOrderGoodUpdate ...
func TpShopOrderGoodUpdate(c *gin.Context) {

	var param service.TpShopOrderGoodUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpShopOrderGoodService{}
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

// TpShopOrderGoodSelect ...
func TpShopOrderGoodSelect(c *gin.Context) {

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
	ser := &service.TpShopOrderGoodService{}
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

