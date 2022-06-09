

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopGoodsAttrService 服务
type TpShopGoodsAttrService struct {
}

// TpShopGoodsAttrCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopGoodsAttrCreateParam struct {
	model.TpShopGoodsAttr
}

// TpShopGoodsAttrCreateBack  返回参数
type TpShopGoodsAttrCreateBack struct {
	model.TpShopGoodsAttr
}

// Create 创建
func ( *TpShopGoodsAttrService) Create(p *TpShopGoodsAttrCreateParam) (*TpShopGoodsAttrCreateBack, error) {
	dao := &dao.TpShopGoodsAttrDao{}
	 
	data, err := dao.Create(&p.TpShopGoodsAttr)

	if err != nil {
		return nil, err

	}

	var back = TpShopGoodsAttrCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopGoodsAttrService) Delete(id int) error {
	dao := &dao.TpShopGoodsAttrDao{}
	return dao.Delete(id)
}

 

// TpShopGoodsAttrSelectBack  返回参数
type TpShopGoodsAttrSelectBack struct {
	model.TpShopGoodsAttr
}

// Select ...
func (*TpShopGoodsAttrService) Select(id int) (*TpShopGoodsAttrSelectBack, error) {
	dao := &dao.TpShopGoodsAttrDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodsAttrSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopGoodsAttrUpdateParam   参数 
type TpShopGoodsAttrUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopGoodsAttrUpdateBack  返回参数
type TpShopGoodsAttrUpdateBack struct {
	model.TpShopGoodsAttr
}

// Update ...
func (*TpShopGoodsAttrService) Update(p *TpShopGoodsAttrUpdateParam) (*TpShopGoodsAttrUpdateBack, error) {
	dao := &dao.TpShopGoodsAttrDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodsAttrUpdateBack {
		*data,
	}
	 
	return &back, nil
}

