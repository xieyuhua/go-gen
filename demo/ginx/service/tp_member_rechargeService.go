

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpMemberRechargeService 服务
type TpMemberRechargeService struct {
}

// TpMemberRechargeCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpMemberRechargeCreateParam struct {
	model.TpMemberRecharge
}

// TpMemberRechargeCreateBack  返回参数
type TpMemberRechargeCreateBack struct {
	model.TpMemberRecharge
}

// Create 创建
func ( *TpMemberRechargeService) Create(p *TpMemberRechargeCreateParam) (*TpMemberRechargeCreateBack, error) {
	dao := &dao.TpMemberRechargeDao{}
	 
	data, err := dao.Create(&p.TpMemberRecharge)

	if err != nil {
		return nil, err

	}

	var back = TpMemberRechargeCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpMemberRechargeService) Delete(id int) error {
	dao := &dao.TpMemberRechargeDao{}
	return dao.Delete(id)
}

 

// TpMemberRechargeSelectBack  返回参数
type TpMemberRechargeSelectBack struct {
	model.TpMemberRecharge
}

// Select ...
func (*TpMemberRechargeService) Select(id int) (*TpMemberRechargeSelectBack, error) {
	dao := &dao.TpMemberRechargeDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpMemberRechargeSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpMemberRechargeUpdateParam   参数 
type TpMemberRechargeUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpMemberRechargeUpdateBack  返回参数
type TpMemberRechargeUpdateBack struct {
	model.TpMemberRecharge
}

// Update ...
func (*TpMemberRechargeService) Update(p *TpMemberRechargeUpdateParam) (*TpMemberRechargeUpdateBack, error) {
	dao := &dao.TpMemberRechargeDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpMemberRechargeUpdateBack {
		*data,
	}
	 
	return &back, nil
}

