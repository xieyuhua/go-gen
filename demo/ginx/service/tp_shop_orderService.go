

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopOrderService 服务
type TpShopOrderService struct {
}

// TpShopOrderCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopOrderCreateParam struct {
	model.TpShopOrder
}

// TpShopOrderCreateBack  返回参数
type TpShopOrderCreateBack struct {
	model.TpShopOrder
}

// Create 创建
func ( *TpShopOrderService) Create(p *TpShopOrderCreateParam) (*TpShopOrderCreateBack, error) {
	dao := &dao.TpShopOrderDao{}
	 
	data, err := dao.Create(&p.TpShopOrder)

	if err != nil {
		return nil, err

	}

	var back = TpShopOrderCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopOrderService) Delete(id int) error {
	dao := &dao.TpShopOrderDao{}
	return dao.Delete(id)
}

 

// TpShopOrderSelectBack  返回参数
type TpShopOrderSelectBack struct {
	model.TpShopOrder
}

// Select ...
func (*TpShopOrderService) Select(id int) (*TpShopOrderSelectBack, error) {
	dao := &dao.TpShopOrderDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopOrderSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopOrderUpdateParam   参数 
type TpShopOrderUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopOrderUpdateBack  返回参数
type TpShopOrderUpdateBack struct {
	model.TpShopOrder
}

// Update ...
func (*TpShopOrderService) Update(p *TpShopOrderUpdateParam) (*TpShopOrderUpdateBack, error) {
	dao := &dao.TpShopOrderDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopOrderUpdateBack {
		*data,
	}
	 
	return &back, nil
}

