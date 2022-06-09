

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopGoodsSpecPriceService 服务
type TpShopGoodsSpecPriceService struct {
}

// TpShopGoodsSpecPriceCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopGoodsSpecPriceCreateParam struct {
	model.TpShopGoodsSpecPrice
}

// TpShopGoodsSpecPriceCreateBack  返回参数
type TpShopGoodsSpecPriceCreateBack struct {
	model.TpShopGoodsSpecPrice
}

// Create 创建
func ( *TpShopGoodsSpecPriceService) Create(p *TpShopGoodsSpecPriceCreateParam) (*TpShopGoodsSpecPriceCreateBack, error) {
	dao := &dao.TpShopGoodsSpecPriceDao{}
	 
	data, err := dao.Create(&p.TpShopGoodsSpecPrice)

	if err != nil {
		return nil, err

	}

	var back = TpShopGoodsSpecPriceCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopGoodsSpecPriceService) Delete(id int) error {
	dao := &dao.TpShopGoodsSpecPriceDao{}
	return dao.Delete(id)
}

 

// TpShopGoodsSpecPriceSelectBack  返回参数
type TpShopGoodsSpecPriceSelectBack struct {
	model.TpShopGoodsSpecPrice
}

// Select ...
func (*TpShopGoodsSpecPriceService) Select(id int) (*TpShopGoodsSpecPriceSelectBack, error) {
	dao := &dao.TpShopGoodsSpecPriceDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodsSpecPriceSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopGoodsSpecPriceUpdateParam   参数 
type TpShopGoodsSpecPriceUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopGoodsSpecPriceUpdateBack  返回参数
type TpShopGoodsSpecPriceUpdateBack struct {
	model.TpShopGoodsSpecPrice
}

// Update ...
func (*TpShopGoodsSpecPriceService) Update(p *TpShopGoodsSpecPriceUpdateParam) (*TpShopGoodsSpecPriceUpdateBack, error) {
	dao := &dao.TpShopGoodsSpecPriceDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodsSpecPriceUpdateBack {
		*data,
	}
	 
	return &back, nil
}

