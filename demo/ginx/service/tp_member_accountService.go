

package service
import (
	"ginx/dao"
	"ginx/model"
)

// TpMemberAccountService 服务
type TpMemberAccountService struct {
}

// TpMemberAccountCreateParam  参数tag上可以增加 binding:"required" 指定为必传(gin解析层会判断)
type TpMemberAccountCreateParam struct {
	model.TpMemberAccount
}

// TpMemberAccountCreateBack  返回参数
type TpMemberAccountCreateBack struct {
	model.TpMemberAccount
}

// Create 创建
func ( *TpMemberAccountService) Create(p *TpMemberAccountCreateParam) (*TpMemberAccountCreateBack, error) {
	dao := &dao.TpMemberAccountDao{}
	 
	data, err := dao.Create(&p.TpMemberAccount)

	if err != nil {
		return nil, err

	}

	var back = TpMemberAccountCreateBack {
		*data,
	}
	  
	return &back, nil
}

 
// Delete  ...
func ( *TpMemberAccountService) Delete(id int) error {
	dao := &dao.TpMemberAccountDao{}
	return dao.Delete(id)
}

 

// TpMemberAccountSelectBack  返回参数
type TpMemberAccountSelectBack struct {
	model.TpMemberAccount
}

// Select ...
func (*TpMemberAccountService) Select(id int) (*TpMemberAccountSelectBack, error) {
	dao := &dao.TpMemberAccountDao{}
	data, err := dao.SelectByID(id)
	if err != nil {
		return nil, err
	}
	var back = TpMemberAccountSelectBack{
		*data,
	}
	  
	return &back, nil
}

// TpMemberAccountUpdateParam   参数 
type TpMemberAccountUpdateParam struct {
	ID    int                    "json:\"id\" binding:\"required\""
	Param map[string]interface{} "json:\"param\" binding:\"required\""
}

// TpMemberAccountUpdateBack  返回参数
type TpMemberAccountUpdateBack struct {
	model.TpMemberAccount
}

// Update ...
func (*TpMemberAccountService) Update(p *TpMemberAccountUpdateParam) (*TpMemberAccountUpdateBack, error) {
	dao := &dao.TpMemberAccountDao{}
	data, err := dao.Update(p.ID, p.Param)
	if err != nil {
		return nil, err
	}
	var back = TpMemberAccountUpdateBack {
		*data,
	}
	 
	return &back, nil
}

