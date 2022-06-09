

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpAdminOperationLogService 服务
type TpAdminOperationLogService struct {
}

// TpAdminOperationLogCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpAdminOperationLogCreateParam struct {
	model.TpAdminOperationLog
}

// TpAdminOperationLogCreateBack  返回参数
type TpAdminOperationLogCreateBack struct {
	model.TpAdminOperationLog
}

// Create 创建
func ( *TpAdminOperationLogService) Create(p *TpAdminOperationLogCreateParam) (*TpAdminOperationLogCreateBack, error) {
	dao := &dao.TpAdminOperationLogDao{}
	 
	data, err := dao.Create(&p.TpAdminOperationLog)

	if err != nil {
		return nil, err

	}

	var back = TpAdminOperationLogCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpAdminOperationLogService) Delete(id int) error {
	dao := &dao.TpAdminOperationLogDao{}
	return dao.Delete(id)
}

 

// TpAdminOperationLogSelectBack  返回参数
type TpAdminOperationLogSelectBack struct {
	model.TpAdminOperationLog
}

// Select ...
func (*TpAdminOperationLogService) Select(id int) (*TpAdminOperationLogSelectBack, error) {
	dao := &dao.TpAdminOperationLogDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpAdminOperationLogSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpAdminOperationLogUpdateParam   参数 
type TpAdminOperationLogUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpAdminOperationLogUpdateBack  返回参数
type TpAdminOperationLogUpdateBack struct {
	model.TpAdminOperationLog
}

// Update ...
func (*TpAdminOperationLogService) Update(p *TpAdminOperationLogUpdateParam) (*TpAdminOperationLogUpdateBack, error) {
	dao := &dao.TpAdminOperationLogDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpAdminOperationLogUpdateBack {
		*data,
	}
	 
	return &back, nil
}

