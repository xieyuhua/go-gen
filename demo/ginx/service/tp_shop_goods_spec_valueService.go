

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopGoodsSpecValueService 服务
type TpShopGoodsSpecValueService struct {
}

// TpShopGoodsSpecValueCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopGoodsSpecValueCreateParam struct {
	model.TpShopGoodsSpecValue
}

// TpShopGoodsSpecValueCreateBack  返回参数
type TpShopGoodsSpecValueCreateBack struct {
	model.TpShopGoodsSpecValue
}

// Create 创建
func ( *TpShopGoodsSpecValueService) Create(p *TpShopGoodsSpecValueCreateParam) (*TpShopGoodsSpecValueCreateBack, error) {
	dao := &dao.TpShopGoodsSpecValueDao{}
	 
	data, err := dao.Create(&p.TpShopGoodsSpecValue)

	if err != nil {
		return nil, err

	}

	var back = TpShopGoodsSpecValueCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopGoodsSpecValueService) Delete(id int) error {
	dao := &dao.TpShopGoodsSpecValueDao{}
	return dao.Delete(id)
}

 

// TpShopGoodsSpecValueSelectBack  返回参数
type TpShopGoodsSpecValueSelectBack struct {
	model.TpShopGoodsSpecValue
}

// Select ...
func (*TpShopGoodsSpecValueService) Select(id int) (*TpShopGoodsSpecValueSelectBack, error) {
	dao := &dao.TpShopGoodsSpecValueDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodsSpecValueSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopGoodsSpecValueUpdateParam   参数 
type TpShopGoodsSpecValueUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopGoodsSpecValueUpdateBack  返回参数
type TpShopGoodsSpecValueUpdateBack struct {
	model.TpShopGoodsSpecValue
}

// Update ...
func (*TpShopGoodsSpecValueService) Update(p *TpShopGoodsSpecValueUpdateParam) (*TpShopGoodsSpecValueUpdateBack, error) {
	dao := &dao.TpShopGoodsSpecValueDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodsSpecValueUpdateBack {
		*data,
	}
	 
	return &back, nil
}

