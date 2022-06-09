

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopCartService 服务
type TpShopCartService struct {
}

// TpShopCartCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopCartCreateParam struct {
	model.TpShopCart
}

// TpShopCartCreateBack  返回参数
type TpShopCartCreateBack struct {
	model.TpShopCart
}

// Create 创建
func ( *TpShopCartService) Create(p *TpShopCartCreateParam) (*TpShopCartCreateBack, error) {
	dao := &dao.TpShopCartDao{}
	 
	data, err := dao.Create(&p.TpShopCart)

	if err != nil {
		return nil, err

	}

	var back = TpShopCartCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopCartService) Delete(id int) error {
	dao := &dao.TpShopCartDao{}
	return dao.Delete(id)
}

 

// TpShopCartSelectBack  返回参数
type TpShopCartSelectBack struct {
	model.TpShopCart
}

// Select ...
func (*TpShopCartService) Select(id int) (*TpShopCartSelectBack, error) {
	dao := &dao.TpShopCartDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopCartSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopCartUpdateParam   参数 
type TpShopCartUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopCartUpdateBack  返回参数
type TpShopCartUpdateBack struct {
	model.TpShopCart
}

// Update ...
func (*TpShopCartService) Update(p *TpShopCartUpdateParam) (*TpShopCartUpdateBack, error) {
	dao := &dao.TpShopCartDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopCartUpdateBack {
		*data,
	}
	 
	return &back, nil
}

