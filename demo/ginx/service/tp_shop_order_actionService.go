

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShopOrderActionService 服务
type TpShopOrderActionService struct {
}

// TpShopOrderActionCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShopOrderActionCreateParam struct {
	model.TpShopOrderAction
}

// TpShopOrderActionCreateBack  返回参数
type TpShopOrderActionCreateBack struct {
	model.TpShopOrderAction
}

// Create 创建
func ( *TpShopOrderActionService) Create(p *TpShopOrderActionCreateParam) (*TpShopOrderActionCreateBack, error) {
	dao := &dao.TpShopOrderActionDao{}
	 
	data, err := dao.Create(&p.TpShopOrderAction)

	if err != nil {
		return nil, err

	}

	var back = TpShopOrderActionCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShopOrderActionService) Delete(id int) error {
	dao := &dao.TpShopOrderActionDao{}
	return dao.Delete(id)
}

 

// TpShopOrderActionSelectBack  返回参数
type TpShopOrderActionSelectBack struct {
	model.TpShopOrderAction
}

// Select ...
func (*TpShopOrderActionService) Select(id int) (*TpShopOrderActionSelectBack, error) {
	dao := &dao.TpShopOrderActionDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShopOrderActionSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShopOrderActionUpdateParam   参数 
type TpShopOrderActionUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShopOrderActionUpdateBack  返回参数
type TpShopOrderActionUpdateBack struct {
	model.TpShopOrderAction
}

// Update ...
func (*TpShopOrderActionService) Update(p *TpShopOrderActionUpdateParam) (*TpShopOrderActionUpdateBack, error) {
	dao := &dao.TpShopOrderActionDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShopOrderActionUpdateBack {
		*data,
	}
	 
	return &back, nil
}

