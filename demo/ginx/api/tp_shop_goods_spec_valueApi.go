
package api
import (
	"ginx/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tpshopgoodsspecvalueInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/tpshopgoodsspecvalue/create
	groupgo.POST("/tpshopgoodsspecvalue/create", TpShopGoodsSpecValueCreate)
	groupgo.POST("/tpshopgoodsspecvalue/delete", TpShopGoodsSpecValueDelete)
	groupgo.POST("/tpshopgoodsspecvalue/update", TpShopGoodsSpecValueUpdate)
	groupgo.POST("/tpshopgoodsspecvalue/select", TpShopGoodsSpecValueSelect)
}

// TpShopGoodsSpecValueCreate ...
func TpShopGoodsSpecValueCreate(c *gin.Context) {

	var param service.TpShopGoodsSpecValueCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpShopGoodsSpecValueService{}
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

// TpShopGoodsSpecValueDelete ...
func TpShopGoodsSpecValueDelete(c *gin.Context) {

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
	ser := &service.TpShopGoodsSpecValueService{}
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

// TpShopGoodsSpecValueUpdate ...
func TpShopGoodsSpecValueUpdate(c *gin.Context) {

	var param service.TpShopGoodsSpecValueUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpShopGoodsSpecValueService{}
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

// TpShopGoodsSpecValueSelect ...
func TpShopGoodsSpecValueSelect(c *gin.Context) {

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
	ser := &service.TpShopGoodsSpecValueService{}
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

