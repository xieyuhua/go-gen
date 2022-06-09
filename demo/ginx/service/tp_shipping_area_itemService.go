

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShippingAreaItemService 服务
type TpShippingAreaItemService struct {
}

// TpShippingAreaItemCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShippingAreaItemCreateParam struct {
	model.TpShippingAreaItem
}

// TpShippingAreaItemCreateBack  返回参数
type TpShippingAreaItemCreateBack struct {
	model.TpShippingAreaItem
}

// Create 创建
func ( *TpShippingAreaItemService) Create(p *TpShippingAreaItemCreateParam) (*TpShippingAreaItemCreateBack, error) {
	dao := &dao.TpShippingAreaItemDao{}
	 
	data, err := dao.Create(&p.TpShippingAreaItem)

	if err != nil {
		return nil, err

	}

	var back = TpShippingAreaItemCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShippingAreaItemService) Delete(id int) error {
	dao := &dao.TpShippingAreaItemDao{}
	return dao.Delete(id)
}

 

// TpShippingAreaItemSelectBack  返回参数
type TpShippingAreaItemSelectBack struct {
	model.TpShippingAreaItem
}

// Select ...
func (*TpShippingAreaItemService) Select(id int) (*TpShippingAreaItemSelectBack, error) {
	dao := &dao.TpShippingAreaItemDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShippingAreaItemSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShippingAreaItemUpdateParam   参数 
type TpShippingAreaItemUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShippingAreaItemUpdateBack  返回参数
type TpShippingAreaItemUpdateBack struct {
	model.TpShippingAreaItem
}

// Update ...
func (*TpShippingAreaItemService) Update(p *TpShippingAreaItemUpdateParam) (*TpShippingAreaItemUpdateBack, error) {
	dao := &dao.TpShippingAreaItemDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShippingAreaItemUpdateBack {
		*data,
	}
	 
	return &back, nil
}

