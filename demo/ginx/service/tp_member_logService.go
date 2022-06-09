

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpMemberLogService 服务
type TpMemberLogService struct {
}

// TpMemberLogCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpMemberLogCreateParam struct {
	model.TpMemberLog
}

// TpMemberLogCreateBack  返回参数
type TpMemberLogCreateBack struct {
	model.TpMemberLog
}

// Create 创建
func ( *TpMemberLogService) Create(p *TpMemberLogCreateParam) (*TpMemberLogCreateBack, error) {
	dao := &dao.TpMemberLogDao{}
	 
	data, err := dao.Create(&p.TpMemberLog)

	if err != nil {
		return nil, err

	}

	var back = TpMemberLogCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpMemberLogService) Delete(id int) error {
	dao := &dao.TpMemberLogDao{}
	return dao.Delete(id)
}

 

// TpMemberLogSelectBack  返回参数
type TpMemberLogSelectBack struct {
	model.TpMemberLog
}

// Select ...
func (*TpMemberLogService) Select(id int) (*TpMemberLogSelectBack, error) {
	dao := &dao.TpMemberLogDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpMemberLogSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpMemberLogUpdateParam   参数 
type TpMemberLogUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpMemberLogUpdateBack  返回参数
type TpMemberLogUpdateBack struct {
	model.TpMemberLog
}

// Update ...
func (*TpMemberLogService) Update(p *TpMemberLogUpdateParam) (*TpMemberLogUpdateBack, error) {
	dao := &dao.TpMemberLogDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpMemberLogUpdateBack {
		*data,
	}
	 
	return &back, nil
}

