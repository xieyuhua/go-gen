

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpShippingComService 服务
type TpShippingComService struct {
}

// TpShippingComCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpShippingComCreateParam struct {
	model.TpShippingCom
}

// TpShippingComCreateBack  返回参数
type TpShippingComCreateBack struct {
	model.TpShippingCom
}

// Create 创建
func ( *TpShippingComService) Create(p *TpShippingComCreateParam) (*TpShippingComCreateBack, error) {
	dao := &dao.TpShippingComDao{}
	 
	data, err := dao.Create(&p.TpShippingCom)

	if err != nil {
		return nil, err

	}

	var back = TpShippingComCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpShippingComService) Delete(id int) error {
	dao := &dao.TpShippingComDao{}
	return dao.Delete(id)
}

 

// TpShippingComSelectBack  返回参数
type TpShippingComSelectBack struct {
	model.TpShippingCom
}

// Select ...
func (*TpShippingComService) Select(id int) (*TpShippingComSelectBack, error) {
	dao := &dao.TpShippingComDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpShippingComSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpShippingComUpdateParam   参数 
type TpShippingComUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpShippingComUpdateBack  返回参数
type TpShippingComUpdateBack struct {
	model.TpShippingCom
}

// Update ...
func (*TpShippingComService) Update(p *TpShippingComUpdateParam) (*TpShippingComUpdateBack, error) {
	dao := &dao.TpShippingComDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpShippingComUpdateBack {
		*data,
	}
	 
	return &back, nil
}

