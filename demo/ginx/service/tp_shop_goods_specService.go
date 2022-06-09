

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopGoodsSpecService 服务
type TpShopGoodsSpecService struct {
}

// TpShopGoodsSpecCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopGoodsSpecCreateParam struct {
	model.TpShopGoodsSpec
}

// TpShopGoodsSpecCreateBack  返回参数
type TpShopGoodsSpecCreateBack struct {
	model.TpShopGoodsSpec
}

// Create 创建
func ( *TpShopGoodsSpecService) Create(p *TpShopGoodsSpecCreateParam) (*TpShopGoodsSpecCreateBack, error) {
	dao := &dao.TpShopGoodsSpecDao{}
	 
	data, err := dao.Create(&p.TpShopGoodsSpec)

	if err != nil {
		return nil, err

	}

	var back = TpShopGoodsSpecCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopGoodsSpecService) Delete(id int) error {
	dao := &dao.TpShopGoodsSpecDao{}
	return dao.Delete(id)
}

 

// TpShopGoodsSpecSelectBack  返回参数
type TpShopGoodsSpecSelectBack struct {
	model.TpShopGoodsSpec
}

// Select ...
func (*TpShopGoodsSpecService) Select(id int) (*TpShopGoodsSpecSelectBack, error) {
	dao := &dao.TpShopGoodsSpecDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodsSpecSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopGoodsSpecUpdateParam   参数 
type TpShopGoodsSpecUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopGoodsSpecUpdateBack  返回参数
type TpShopGoodsSpecUpdateBack struct {
	model.TpShopGoodsSpec
}

// Update ...
func (*TpShopGoodsSpecService) Update(p *TpShopGoodsSpecUpdateParam) (*TpShopGoodsSpecUpdateBack, error) {
	dao := &dao.TpShopGoodsSpecDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopGoodsSpecUpdateBack {
		*data,
	}
	 
	return &back, nil
}

