

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpDeliveryLogService 服务
type TpDeliveryLogService struct {
}

// TpDeliveryLogCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpDeliveryLogCreateParam struct {
	model.TpDeliveryLog
}

// TpDeliveryLogCreateBack  返回参数
type TpDeliveryLogCreateBack struct {
	model.TpDeliveryLog
}

// Create 创建
func ( *TpDeliveryLogService) Create(p *TpDeliveryLogCreateParam) (*TpDeliveryLogCreateBack, error) {
	dao := &dao.TpDeliveryLogDao{}
	 
	data, err := dao.Create(&p.TpDeliveryLog)

	if err != nil {
		return nil, err

	}

	var back = TpDeliveryLogCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpDeliveryLogService) Delete(id int) error {
	dao := &dao.TpDeliveryLogDao{}
	return dao.Delete(id)
}

 

// TpDeliveryLogSelectBack  返回参数
type TpDeliveryLogSelectBack struct {
	model.TpDeliveryLog
}

// Select ...
func (*TpDeliveryLogService) Select(id int) (*TpDeliveryLogSelectBack, error) {
	dao := &dao.TpDeliveryLogDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpDeliveryLogSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpDeliveryLogUpdateParam   参数 
type TpDeliveryLogUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpDeliveryLogUpdateBack  返回参数
type TpDeliveryLogUpdateBack struct {
	model.TpDeliveryLog
}

// Update ...
func (*TpDeliveryLogService) Update(p *TpDeliveryLogUpdateParam) (*TpDeliveryLogUpdateBack, error) {
	dao := &dao.TpDeliveryLogDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpDeliveryLogUpdateBack {
		*data,
	}
	 
	return &back, nil
}

