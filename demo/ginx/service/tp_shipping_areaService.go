

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShippingAreaService 服务
type TpShippingAreaService struct {
}

// TpShippingAreaCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShippingAreaCreateParam struct {
	model.TpShippingArea
}

// TpShippingAreaCreateBack  返回参数
type TpShippingAreaCreateBack struct {
	model.TpShippingArea
}

// Create 创建
func ( *TpShippingAreaService) Create(p *TpShippingAreaCreateParam) (*TpShippingAreaCreateBack, error) {
	dao := &dao.TpShippingAreaDao{}
	 
	data, err := dao.Create(&p.TpShippingArea)

	if err != nil {
		return nil, err

	}

	var back = TpShippingAreaCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShippingAreaService) Delete(id int) error {
	dao := &dao.TpShippingAreaDao{}
	return dao.Delete(id)
}

 

// TpShippingAreaSelectBack  返回参数
type TpShippingAreaSelectBack struct {
	model.TpShippingArea
}

// Select ...
func (*TpShippingAreaService) Select(id int) (*TpShippingAreaSelectBack, error) {
	dao := &dao.TpShippingAreaDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShippingAreaSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShippingAreaUpdateParam   参数 
type TpShippingAreaUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShippingAreaUpdateBack  返回参数
type TpShippingAreaUpdateBack struct {
	model.TpShippingArea
}

// Update ...
func (*TpShippingAreaService) Update(p *TpShippingAreaUpdateParam) (*TpShippingAreaUpdateBack, error) {
	dao := &dao.TpShippingAreaDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShippingAreaUpdateBack {
		*data,
	}
	 
	return &back, nil
}

