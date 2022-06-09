
package api
import (
	"ginx/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tpshopgoodsspecpriceInit(groupgo *gin.RouterGroup) {
	 
	// 所有api的开头 都加上/go  例如 /go/tpshopgoodsspecprice/create
	groupgo.POST("/tpshopgoodsspecprice/create", TpShopGoodsSpecPriceCreate)
	groupgo.POST("/tpshopgoodsspecprice/delete", TpShopGoodsSpecPriceDelete)
	groupgo.POST("/tpshopgoodsspecprice/update", TpShopGoodsSpecPriceUpdate)
	groupgo.POST("/tpshopgoodsspecprice/select", TpShopGoodsSpecPriceSelect)
}

// TpShopGoodsSpecPriceCreate ...
func TpShopGoodsSpecPriceCreate(c *gin.Context) {

	var param service.TpShopGoodsSpecPriceCreateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpShopGoodsSpecPriceService{}
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

// TpShopGoodsSpecPriceDelete ...
func TpShopGoodsSpecPriceDelete(c *gin.Context) {

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
	ser := &service.TpShopGoodsSpecPriceService{}
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

// TpShopGoodsSpecPriceUpdate ...
func TpShopGoodsSpecPriceUpdate(c *gin.Context) {

	var param service.TpShopGoodsSpecPriceUpdateParam
	//解析参数
	err := c.ShouldBind(&param)
	if err != nil {
		reply := NewReplyError(err.Error())
		c.JSON(http.StatusOK, reply)
		return
	}
	ser := &service.TpShopGoodsSpecPriceService{}
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

// TpShopGoodsSpecPriceSelect ...
func TpShopGoodsSpecPriceSelect(c *gin.Context) {

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
	ser := &service.TpShopGoodsSpecPriceService{}
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

